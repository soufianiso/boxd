package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// "io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type movie struct{
	Id int `json:"id"`
	Original_title string `json:"original_title"`
	Original_language string `json:"original_language"`
	Overview string `json:"overview"`
	Adult bool	`json:"adult"`
	Release_date string `json:"release_date"`
	Poster_path string	`json:"poster_path"`
}

// Struct to represent the API response, which includes a list of movies
type movieResponse struct {
	Page    int     `json:"page"`
	Results []movie `json:"results"`
	Total_results int `json:"total_results"`
	Total_pages   int `json:"total_pages"`
}

func main() {
	godotenv.Load()
	conn := os.Getenv("postgres")
	

	url := "https://api.themoviedb.org/3/discover/movie?include_adult=any&include_video=false&language=en-US&page=5&sort_by=popularity.desc"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI3NjJlNWE2Y2VlOTZlZTU3ZDM1ZWI5N2Y1NTY2YmI1YiIsIm5iZiI6MTcyNTE4OTA0Ni4xNTYxMzUsInN1YiI6IjY2ZDE4OTdhNWY1OTQ1ZTA2ODQ1NTZjZCIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.C-ZbvS2VkAT48FH22IXunlmATkvcjHPI-FNGQeK_B6w")

	res, _ := http.DefaultClient.Do(req)

	movieResponse := new(movieResponse)

	if err := json.NewDecoder(res.Body).Decode(&movieResponse) ; err != nil{
		fmt.Println(err)
	}

	db ,err := sql.Open("postgres",conn)
	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	for _, m := range movieResponse.Results {
			fmt.Printf("ID: %d\n", m.Id)
			fmt.Printf("Title: %s\n", m.Original_title)
			fmt.Printf("Language: %s\n", m.Original_language)
			fmt.Printf("Overview: %s\n", m.Overview)
			fmt.Printf("Adult: %t\n", m.Adult)
			fmt.Printf("release_date: %s\n", m.Release_date)
			fmt.Printf("Poster Path: %s\n", m.Poster_path)
			fmt.Println("-------------------------------")
		_ , err := db.Exec("INSERT INTO movies (original_title , original_language , overview , release_date , adult , poster_path) VALUES ($1, $2 ,$3, $4, $5, $6)",
			m.Original_title, m.Original_language, m.Overview, m.Release_date, m.Adult, m.Poster_path)
		if err != nil{
			log.Println("###############################")
			log.Println(err )
			log.Println("###############################")
		}
	}


}
