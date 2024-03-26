package infrastructure

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/deltrinos/fizzbuzz_server/infrastructure/rest"
	"github.com/deltrinos/fizzbuzz_server/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(port string) {
	// Initialize dependencies
	fizzBuzzService := service.NewFizzBuzzService()

	// Create HTTP handler
	fizzBuzzHandler := rest.NewFizzBuzzHandler(fizzBuzzService)
	statsHandler := rest.NewStatisticsHandler()

	// Add endpoints
	http.HandleFunc("/fizzbuzz", fizzBuzzHandler.HandleFizzBuzz)
	http.HandleFunc("/stats", statsHandler.HandleStatistics)

	// Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Start HTTP server in a separate goroutine
	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second, // Maximum duration for reading the entire request
		WriteTimeout: 5 * time.Second, // Maximum duration for writing the response
	}

	go func() {
		// Start the HTTP server and listen on the specified port.
		log.Printf("Server is running at :%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start server: %v\n", err)
		}
	}()

	// Wait for interrupt signal (e.g., Ctrl+C) to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server is shutting down...")

	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server with the given context
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	}

	// Bye
	log.Println("Server stopped gracefully.")
}
