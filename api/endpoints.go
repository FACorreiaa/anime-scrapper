package api

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"strings"
)

type AnimeListResponse struct {
	Data []struct {
		Node struct {
			Title string `json:"title"`
		} `json:"node"`
	} `json:"data"`
}

func GetUserListSyncedWithNyaa(token string) {
	response, err := GetUserAnimeListAPI(token)
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
	//
}

func GetSeasonalAnime(token string, season string, year int, limit int) {
	seasonalAnime, err := GetSeasonalAnimeAPI(token, season, year, limit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var aniListResponse AnimeListResponse
	err = json.Unmarshal([]byte(seasonalAnime), &aniListResponse)

	for _, entry := range aniListResponse.Data {
		println(entry.Node.Title)
	}
}
