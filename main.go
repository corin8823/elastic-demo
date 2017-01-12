package main

import (
	"fmt"

	"strconv"

	"os"

	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

// Shakespeare is a structure used for serializing/deserializing data in Elasticsearch.
type Shakespeare struct {
	PlayName     string `json:"play_name,omitempty"`
	Speaker      string `json:"speaker,omitempty"`
	SpeechNumber string `json:"speech_number,omitempty"`
	TextEntry    string `json:"text_entry,omitempty"`
}

func main() {
	url := elastic.SetURL(os.Getenv("URL"))
	ba := elastic.SetBasicAuth(os.Getenv("USER"), os.Getenv("PASSWORD"))
	sniff := elastic.SetSniff(false)
	client, err := elastic.NewClient(url, ba, sniff)
	if err != nil {
		fmt.Println(err.Error())
	}

	q := elastic.NewTermQuery("id", 1)
	searchResult, err2 := client.Search().Index("shakespeare").Query(q).Do(context.Background())
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(strconv.FormatInt(searchResult.TookInMillis, 10))
	fmt.Println(strconv.FormatInt(searchResult.TotalHits(), 10))
}
