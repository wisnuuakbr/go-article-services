package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/wisnuuakbr/sagala/config"
	"github.com/wisnuuakbr/sagala/internal/di"
)

func main() {
	cfg := config.New()

	// connection to database
	db, err := sql.Open("postgres", cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// running migration
	if err := goose.Up(db, "../database/migration"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// connection to redis
	redisClient := redis.NewClient(cfg.RedisOptions())
	defer redisClient.Close()

	// route
	r := di.NewRouter(db, redisClient)

	port := fmt.Sprintf(":%d", cfg.App.Port)
	fmt.Printf("Server is running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
