package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/josuesantos1/sangue/internal"
)

func main() {
	fmt.Println("Hello")

	app, err := internal.NewApp("8080")
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Start(); err != nil  && err != http.ErrServerClosed {
			panic(err)
		}
	} ()

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

	 if err := app.Shutdown(ctx); err != nil {
        log.Fatalf("Failled when stoping server: %v", err)
    }

    log.Println("Bye Bye")
}
