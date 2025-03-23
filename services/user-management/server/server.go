package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func SetUpUserServer(routes *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           routes,
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
	}

	// Run server in a goroutine
	go func() {
		log.Printf("Server running on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error %v:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("Shutting down server on port %s", port)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

	log.Println("server shutting down gracefully")

}
