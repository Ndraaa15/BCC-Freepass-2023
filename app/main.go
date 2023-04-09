package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//laad all env variables
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

}
