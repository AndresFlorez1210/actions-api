package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type Action struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type ResponseBody struct {
	Actions  []Action `json:"items"`
	NextPage string   `json:"next_page"`
}

func main() {
	connection := createConnection()
	defer connection.Close(context.Background())
	var nextPage = ""
	log.Println("Starting to fetch actions")
	for i := 0; i < 20; i++ {
		actions := getActions(nextPage)
		nextPage = actions.NextPage
		saveActions(actions.Actions, connection)
	}
	log.Println("Finished fetching and saving actions")
}

func createConnection() *pgx.Conn {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	connStr := fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		dbHost,
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"))

	connection, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return connection
}

func getActions(nextPage string) ResponseBody {
	responseBody, _ := getActionsApi(nextPage)
	return responseBody
}

func getActionsApi(nextPage string) (ResponseBody, *http.Response) {
	url := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
	if nextPage != "" {
		url += "?next_page=" + nextPage
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MywiZW1haWwiOiJmZWxpcGUuZmxvcmV6LmFyQGdtYWlsLmNvbSIsImV4cCI6MTc0MTE4NTY0NSwiaWQiOiIwIiwicGFzc3dvcmQiOiJUcnVvcmEnIE9SICcnPScifQ.3nmY4cJJ7ei7XUbvZtIbyPGR6EkvlZ74IEqqoNUFAw4")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var responseBody ResponseBody
	json.NewDecoder(response.Body).Decode(&responseBody)

	return responseBody, response
}

func saveActions(actions []Action, connection *pgx.Conn) {
	for _, action := range actions {
		_, err := connection.Exec(context.Background(),
			"INSERT INTO actions (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			action.Ticker,
			action.TargetFrom,
			action.TargetTo,
			action.Company,
			action.Action,
			action.Brokerage,
			action.RatingFrom,
			action.RatingTo,
			action.Time,
		)
		if err != nil {
			panic(err)
		}
	}
}
