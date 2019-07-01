package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env
func LoadEnv() {
	godotenv.Load()
	fmt.Println(os.Getenv("PORT"))
}
