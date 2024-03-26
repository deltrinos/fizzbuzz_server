package main

import (
	"flag"
	"os"

	"github.com/deltrinos/fizzbuzz_server/infrastructure"
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

	infrastructure.StartServer(port)
}
