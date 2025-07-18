package main

import (
	"log"
	"nestmate-backend/internal/interfaces/http"
)

func main() {
	server := http.NewServer()
	
	log.Println("Starting NestMate backend server on :8080")
	if err := server.Start(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}