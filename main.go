package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vyas-git/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "Product api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	server.ListenAndServe()
}
