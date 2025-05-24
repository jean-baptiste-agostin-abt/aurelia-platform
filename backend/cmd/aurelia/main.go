package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/yourorg/aurelia-backend/config"
	"github.com/yourorg/aurelia-backend/internal/auth"
	"github.com/yourorg/aurelia-backend/internal/capsule"
	"github.com/yourorg/aurelia-backend/internal/events"
	"github.com/yourorg/aurelia-backend/internal/family"
	"github.com/yourorg/aurelia-backend/internal/legacyguard"
	"github.com/yourorg/aurelia-backend/pkg/mid"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("db init: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	r.Use(mid.CORSMiddleware())

	api := r.Group("/api")
	auth.RegisterRoutes(api, db)
	family.RegisterRoutes(api, db)
	capsule.RegisterRoutes(api, db)
	legacyguard.RegisterRoutes(api, db)
	events.RegisterRoutes(api, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
