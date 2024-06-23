package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"leart.com/art/cli"
	"leart.com/art/server"
)

// http://localhost:8080/
func main() {
	serverFlag := flag.Bool("server", false, "Start server")
	cliFlag := flag.Bool("o", false, "Run CLI without starting server")
	flag.Parse()

	// Configure logger
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// func logError(err error) {
	// 	_, file, line, _ := runtime.Caller(1)
	// 	log.Printf("%s:%d: %v", file, line, err)
	// }

	if *cliFlag {
		// Run CLI
		cli.RunCLI()
	} else if *serverFlag {
		// Start server
		fmt.Println("Starting server on :8080")
		log.Fatal(http.ListenAndServe(":8080", server.NewHandler()))
	} else {
		// If no flags are specified, print an error message
		fmt.Println("Please specify a mode: -server to start the server or -o to run the CLI")
	}
}
