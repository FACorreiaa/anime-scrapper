package main

import (
	"encoding/json"
	"fmt"
	"github.com/FACorreiaa/anime-scrapper/api"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"strings"
)

type AnimeListResponse struct {
	Data []struct {
		Node struct {
			Title string `json:"title"`
		} `json:"node"`
	} `json:"data"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	token := os.Getenv("ACCESS_TOKEN") // Replace with your actual token

	response, err := api.GetUserAnimeList(token)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse the JSON response
	var data AnimeListResponse
	err = json.Unmarshal([]byte(response), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	parser := gofeed.NewParser()
	feed, err := parser.ParseURL("https://nyaa.si/?page=rss&c=1_2&f=0")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range feed.Items {
		for _, entry := range data.Data {
			println(item.Title)
			println(entry.Node.Title)

			if strings.Contains(item.Title, entry.Node.Title) {
				fmt.Println("Title: ", item.Title)
				fmt.Println("Released in: ", item.Published)
				break
			}
		}
	}
	// Create a map to store titles from the API response
	// If we can check later for the exact ti
	//apiResponseTitles := make(map[string]struct{})
	//
	//for _, entry := range data.Data {
	//	apiResponseTitles[entry.Node.Title] = struct{}{}
	//}
	//
	//for _, item := range feed.Items {
	//	fmt.Println("Title: ", item.Title)
	//
	//	if _, ok := apiResponseTitles[item.Title]; ok {
	//
	//		fmt.Println("Title: ", item.Title)
	//		fmt.Println("Released in: ", item.Published)
	//	}
	//}
}
