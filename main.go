package main

import (
	"ToDo/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func main() {

	port := os.Getenv("PORT")
	r, cancelFunc := router.Router()
	defer cancelFunc()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-ctx.Done()
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	} else {
		log.Print("Server Exited Properly")
	}
}
