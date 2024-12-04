package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"git.ana/xjtuana/api-micro-mail/config"
	"git.ana/xjtuana/api-micro-mail/http"
)

func main() {
	cfg := config.New()
	cfg.Load()

	srv := http.NewServer(cfg)
	srv.SetRoutes()

	// service connections
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
