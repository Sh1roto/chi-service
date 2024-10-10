package main

import (
	"chi-service/application"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
