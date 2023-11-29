package infra

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Errr in env file loading")
	}
} 