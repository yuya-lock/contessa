package lib

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Error loading .env file. \n%s", err)
	}
	return nil
}
