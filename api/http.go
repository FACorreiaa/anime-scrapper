package api

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Season represents a season of the year.
type Season string

// Define constants for valid seasons.
const (
	Spring Season = "spring"
	Summer Season = "summer"
	Fall   Season = "fall"
	Winter Season = "winter"
)

// IsValidSeason checks if the given season is valid.
func IsValidSeason(s Season) bool {
	switch s {
	case Spring, Summer, Fall, Winter:
		return true
	}
	return false
}

// GetSeason returns the Season based on a string input and an error if the input is invalid.
func GetSeason(input string) (Season, error) {
	season := Season(input)

	if !IsValidSeason(season) {
		return "", errors.New("Invalid season")
	}

	return season, nil
}

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

func GetUserAnimeListAPI(token string) (string, error) {
	// Define the URL for the GET request
	url := "https://api.myanimelist.net/v2/users/@me/animelist?fields=list_status&status=watching"

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set the Authorization header with the token
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the response body to a string
	responseBody := string(body)

	return responseBody, nil
}

func GetSeasonalAnimeAPI(token string, season string, year int, limit int) (string, error) {
	// Define the URL for the GET request
	baseURL := fmt.Sprintf("https://api.myanimelist.net/v2/anime/season/%d/%s", year, season)

	// Construct the query parameters
	queryParams := url.Values{}
	queryParams.Set("limit", fmt.Sprintf("%d", limit))

	// Add the query parameters to the base URL
	fullURL := baseURL + "?" + queryParams.Encode()
	fmt.Println("fullURL", fullURL)
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", err
	}

	// Set the Authorization header with the token
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request and get the response
	resp, err := client.Do(req)
	fmt.Println("res", resp)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the response body to a string
	responseBody := string(body)

	return responseBody, nil
}
