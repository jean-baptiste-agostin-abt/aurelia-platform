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
	"github.com/yourorg/aurelia-backend/internal/user"
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

	userRepo := user.NewRepository(db)
	authHandler := auth.NewHandler(userRepo)

	familyRepo := family.NewRepository(db)
	familyHandler := family.NewHandler(familyRepo)

	capsuleRepo := capsule.NewRepository(db)
	capsuleHandler := capsule.NewHandler(capsuleRepo)

	legacyRepo := legacyguard.NewRepository(db)
	legacyHandler := legacyguard.NewHandler(legacyRepo)

	eventsRepo := events.NewRepository(db)
	eventsHandler := events.NewHandler(eventsRepo)

	api := r.Group("/api")
	auth.RegisterRoutes(api, authHandler)
	family.RegisterRoutes(api, familyHandler)
	capsule.RegisterRoutes(api, capsuleHandler)
	legacyguard.RegisterRoutes(api, legacyHandler)
	events.RegisterRoutes(api, eventsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
