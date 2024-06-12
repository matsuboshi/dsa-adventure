package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type Data struct {
	Parse struct {
		Text struct{
			Content string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
}

func fetchDataFromWikipedia(topic string) string {
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=%s", topic)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("ERROR:",err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR:",err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR:",err)
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("ERROR:",err)
	}

	return data.Parse.Text.Content
}


func countWordsFromWikipedia(topic string) int {
	text := fetchDataFromWikipedia(topic)
	re := regexp.MustCompile(fmt.Sprintf(`(?i)%s`, topic))
	matches := re.FindAllString(text, -1)
	return len(matches)
}

func main() {	
	count := countWordsFromWikipedia("pizza")
	fmt.Println(count)
}
