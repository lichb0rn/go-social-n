package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"social/internal/store"
)

var usernames = []string{
	"CyberKnight42", "QuantumRogue", "PixelPirate", "NeonSamurai", "WarpDrifter",
	"ShadowSynth", "EchoNomad", "LunarVortex", "CodePhantom", "AetherSeeker",
	"ZenGlitch", "DataRonin", "ByteWarlock", "OmegaCipher", "GlitchHunter",
	"PlasmaShade", "VoidSprinter", "HoloRaider", "NovaStrider", "MetaWanderer",
	"CircuitRanger", "BinaryShaman", "SynthMarauder", "HyperZenith", "NeuralVoyager",
	"NanoSpecter", "CelestialWeaver", "PhantomFrame", "QuantumSage", "CyberSentry",
	"WarpCaster", "EchoCipher", "MechaMystic", "LunarRonin", "PixelOracle",
	"HoloWraith", "DataValkyrie", "ShadowChronicle", "AetherPilgrim", "NeonWarlock",
	"OmegaGlider", "CircuitNomad", "GlitchWanderer", "ZenByte", "BinaryKnight",
	"HyperDrifter", "NovaSynth", "VoidRaider", "NeuralHunter", "PlasmaRogue",
}

var titles = []string{
	"Mastering Go: Tips and Tricks for Efficient Coding",
	"The Rise of Microservices: Benefits and Challenges",
	"Building Scalable APIs with gRPC in Go",
	"Concurrency in Go: Goroutines and Channels Explained",
	"Optimizing Docker Containers for Production Workloads",
	"Understanding Clean Architecture in Go Applications",
	"Deploying Go Applications with Kubernetes",
	"A Deep Dive into MongoDB with Go",
	"Building a CI/CD Pipeline for Your Go Projects",
	"Logging and Monitoring in Distributed Systems",
	"How to Design a Domain-Driven Microservice Architecture",
	"Handling Authentication and Authorization in Go",
	"Creating Performant Web Applications with React and Go",
	"Using WebSockets for Real-Time Applications in Go",
	"Automating Infrastructure with Terraform and Ansible",
	"Designing RESTful APIs: Best Practices and Common Pitfalls",
	"Scaling Applications with Message Queues and Event-Driven Architecture",
	"Writing Unit and Integration Tests in Go",
	"Understanding the Role of Feature Flags in Software Development",
	"Breaking Down the Monolith: A Guide to Service Decomposition",
}

var contents = []string{
	"The future belongs to those who prepare for it today.",
	"Life is 10% what happens to us and 90% how we react to it.",
	"Simplicity is the ultimate sophistication.",
	"Go boldly where no one has gone before.",
	"In the middle of difficulty lies opportunity.",
	"The best way to predict the future is to create it.",
	"Code is like humor. When you have to explain it, it’s bad.",
	"Failure is simply the opportunity to begin again, this time more intelligently.",
	"A journey of a thousand miles begins with a single step.",
	"Your limitation—it’s only your imagination.",
	"Do what you can, with what you have, where you are.",
	"Every accomplishment starts with the decision to try.",
	"Dream big and dare to fail.",
	"Happiness depends upon ourselves.",
	"Turn your wounds into wisdom.",
	"Strive not to be a success, but rather to be of value.",
	"The secret to getting ahead is getting started.",
	"A smooth sea never made a skilled sailor.",
	"Difficulties strengthen the mind, as labor does the body.",
	"Quality means doing it right when no one is looking.",
}

var tags = []string{
	"golang", "microservices", "webdev", "cloud", "docker",
	"kubernetes", "backend", "frontend", "javascript", "typescript",
	"react", "devops", "testing", "cicd", "api",
	"database", "security", "opensource", "performance", "scalability",
}

var comments = []string{
	"Great post! Very informative.",
	"I totally agree with this!",
	"This helped me a lot, thanks!",
	"Interesting perspective, I never thought about it that way.",
	"Could you elaborate on this part?",
	"I ran into an issue, any advice?",
	"Awesome work, keep it up!",
	"Not sure I understand this, can you clarify?",
	"Thanks for sharing, really useful!",
	"Well explained, appreciate it!",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)
	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user: ", err)
		}
	}
	tx.Commit()

	post := generatePosts(200, users)
	for _, post := range post {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
		}
	}

	comments := generateComments(500, users, post)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment: ", err)
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(count int) []*store.User {
	log.Println("Generating users")

	users := make([]*store.User, count)
	for i := range count {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}
	return users
}

func generatePosts(count int, users []*store.User) []*store.Post {
	log.Println("Generating post")

	posts := make([]*store.Post, count)
	for i := range count {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags:    []string{tags[rand.Intn(len(tags))], tags[rand.Intn(len(tags))]},
		}
	}
	return posts
}

func generateComments(count int, users []*store.User, posts []*store.Post) []*store.Comment {
	log.Println("Generating comments")
	cms := make([]*store.Comment, count)
	for i := range count {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		cms[i] = &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
