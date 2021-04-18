package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	envFile = flag.String("envFile", ".env", "The environment File")
)

func loadEnv(){
	fmt.Println(*envFile)

	err := godotenv.Load(*envFile)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func goDotEnvVariable(key string) string {
	return os.Getenv(key)
}