package main

import (
	"context"
	"log"
	"net/http"
	"net/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/mux1"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(CreateGreeting(name)))
}

func CreateGreeting(name string) string {
	if name == "" {
		name = "Guest"
	}
	return "Hello, " + name + "\n"
}

func CreateGreeting1(name string) string {
	if name == "" {
		name = "Guest"
	}
	return "Hello, " + name + "\n"
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)


	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
