package main

import (
	"context"
	"fizzbuzz_server/internal/fizzbuzz"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Default port if not specified
const DEFAULT_PORT = "8000"

func main() {
	var port string

	// Define a command-line flag named "port" to specify the port number.
	// The default value is an empty string.
	// Usage: ./fizzbuzz_server -port=8000
	flag.StringVar(&port, "port", "", "Port number for the server (e.g., "+DEFAULT_PORT+")")

	// Parse the command-line flags.
	flag.Parse()

	// If the port number is not provided via the flag, check the PORT environment variable.
	if port == "" {
		port = os.Getenv("PORT")
	}

	// If the port number is still not set, use the default port DEFAULT_PORT.
	if port == "" {
		port = DEFAULT_PORT
	}

	// Create instances of FizzBuzzHandler and StatisticsHandler.
	fbHandler := fizzbuzz.NewFizzBuzzHandler()
	statsHandler := fizzbuzz.NewStatisticsHandler()

	// Create a new HTTP server with a custom handler
	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second, // Maximum duration for reading the entire request
		WriteTimeout: 5 * time.Second, // Maximum duration for writing the response
	}

	// Register the endpoints with their corresponding handlers.
	http.HandleFunc("/fizzbuzz", fbHandler.HandleFizzBuzz)
	http.HandleFunc("/stats", statsHandler.HandleStatistics)

	// Start the server in a separate goroutine
	go func() {
		// Start the HTTP server and listen on the specified port.
		log.Printf("Server is running at :%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
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
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")
}
