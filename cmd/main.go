package main

import (
	"cmd/app/config"
	"cmd/di"
	"context"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if os.Getenv("IN_CONTAINER") == "True" {
		return
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found: ", err)
	}
}

func main() {
	container, err := di.NewContainer(
		config.Params{
			DatabaseURL:   os.Getenv("DATABASE_URL"),
			ServerAddress: os.Getenv("SERVER_ADDRESS"),
			Secret:        os.Getenv("SECRET"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uuid.Must(uuid.NewV4()))

	server, err := container.Server(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	container.Close()
}
