package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getNewCodeVerifier() (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("MY_ANIME_LIST_API_KEY")
	apiKeyBytes := []byte(apiKey)

	randomBytes := make([]byte, 32)
	_, err = rand.Read(apiKeyBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 URL-safe format
	codeVerifier := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Trim padding characters (if any) and ensure the length is 43 characters
	codeVerifier = codeVerifier[:43]

	return codeVerifier, nil
}

func main() {
	codeVerifier, err := getNewCodeVerifier()
	if err != nil {
		fmt.Println("Error generating code verifier:", err)
		return
	}

	fmt.Println(len(codeVerifier))
	fmt.Println(codeVerifier)
}
