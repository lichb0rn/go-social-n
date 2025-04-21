package main

import (
	"social/internal/auth"
	"social/internal/db"
	"social/internal/env"
	"social/internal/mailer"
	"social/internal/store"
	"time"

	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			Simple Social Network API
//	@description	API for simple social network.

//	@license.name	MIT
//	@license.url	https://mit-license.org

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				API key authentication
func main() {
	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		env:         env.GetString("ENV", "development"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://http://localhost:5173/"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:password@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		mail: mailConfig{
			exp:       time.Hour * 24 * 3,
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicAuthConfig{
				username: env.GetString("AUTH_BASIC_USER", "admin"),
				password: env.GetString("AUTH_BASIC_PASSWORD", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "supersecret"),
				exp:    time.Hour * 24 * 3,
				iss:    "simple-social-network",
			},
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewPostgresStorage(db)
	mailer := mailer.NewSendgrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)

	jwtAuth := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		mailer:        mailer,
		authenticator: jwtAuth,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
