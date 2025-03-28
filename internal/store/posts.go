package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
	Version   int       `json:"version"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentCount int `json:"comment_count"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (title, content, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, postId int64) (*Post, error) {
	query := `
	SELECT id, title, content, user_id, tags, created_at, updated_at, version
	FROM posts
	WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeoutDuration)
	defer cancel()

	post := &Post{}
	err := s.db.QueryRowContext(ctx, query, postId).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)

	if err != nil {
		switch errors.Is(err, sql.ErrNoRows) {
		case true:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return post, nil
}

func (s *PostStore) Delete(ctx context.Context, postId int64) error {
	query := `
	DELETE FROM posts
	WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeoutDuration)
	defer cancel()

	result, err := s.db.ExecContext(ctx, query, postId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	query := `
	UPDATE posts
	SET title = $1, content = $2, updated_at = NOW(), version = version + 1
	WHERE id = $3 AND version = $4
	RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version)

	if err != nil {
		switch errors.Is(err, sql.ErrNoRows) {
		case true:
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *PostStore) GetUserFeed(ctx context.Context, userId int64, fq PaginatedFeedQuery) ([]*PostWithMetadata, error) {
	query := `
	SELECT p.id,p.user_id, p.title, p.content, p.created_at, p.version, p.tags, u.username,
	COUNT(c.id) AS comment_count
	FROM posts p
	LEFT JOIN comments c ON c.post_id = p.id
	LEFT JOIN users u ON p.user_id = u.id
	JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
	WHERE f.user_id = $1 OR p.user_id = $1
	GROUP BY p.id, u.username
	ORDER BY p.created_at ` + fq.Sort + `
	LIMIT $2 OFFSET $3
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userId, fq.Limit, fq.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feed []*PostWithMetadata
	for rows.Next() {
		var post PostWithMetadata
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.Version, pq.Array(&post.Tags), &post.User.Username, &post.CommentCount)
		if err != nil {
			return nil, err
		}
		feed = append(feed, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return feed, nil
}
