package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vyas-git/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "Product api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	sig := <-signChan
	l.Println("Received terminate, graceful shutdown", sig)
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	s.Shutdown(ctx)
}
