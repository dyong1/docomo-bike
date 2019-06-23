package main

import (
	"docomo-bike/internal/app"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	s, err := app.NewServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	log.Fatal(s.ServeHTTP(addr))
}
