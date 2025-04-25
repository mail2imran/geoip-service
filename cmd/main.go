package main

import (
	"geoip-service/config"
	"geoip-service/internal/geoip"
	"geoip-service/internal/handler"
	"geoip-service/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadEnv()
	dbPath := config.GetEnv("MAXMIND_DB_PATH", "config/GeoLite2-Country.mmdb")

	resolver, err := geoip.NewResolver(dbPath)
	if err != nil {
		log.Fatalf("Failed to load GeoIP DB: %v", err)
	}

	r := gin.Default()
	r.POST("/auth/token", middleware.LoginHandler)
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	auth.POST("/ip/verify", handler.CheckIPHandler(resolver))

	log.Println("Starting server on :8080")
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
