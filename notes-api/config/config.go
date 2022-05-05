package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	
	// load .env file
	err := godotenv.Load(".env")
	
	if err != nil {
		fmt.Printf("Try get key %s but error loading .env file \n", key)
	}
	
	val := os.Getenv(key)
	fmt.Printf("key %s = %s \n", key, val )
	
	return val
}