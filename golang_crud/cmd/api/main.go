package main

import (
	"fmt"
	"log"

	"golang-crud/internal/config"
	"golang-crud/internal/db"
	"golang-crud/internal/server"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// Connect to MongoDB
	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Fatalf("Failed to disconnect: %v", err)
		}
	}()

	fmt.Println("Database connected:", database.Name())

	// Setup router
	router := server.NewRouter(database)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
