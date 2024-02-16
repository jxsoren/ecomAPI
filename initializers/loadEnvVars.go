package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env")
	}

	fmt.Println("Env Loaded Successfully!")
}
