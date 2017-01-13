package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

// Shakespeare is a structure used for serializing/deserializing data in Elasticsearch.
type Shakespeare struct {
	LineID       int    `json:"line_id"`
	PlayName     string `json:"play_name"`
	SpeechNumber int    `json:"speech_number"`
	LineNumbar   string `json:"line_number"`
	Speaker      string `json:"speaker"`
	TextEntry    string `json:"text_entry"`
}

func main() {
	client, _ := NewClient()

	// q := elastic.NewTermQuery("_index", "shakespeare")
	q := elastic.NewSimpleQueryStringQuery("Alls")
	result, err := client.Search().Index("shakespeare").Query(q).Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("TotalHits:", strconv.FormatInt(result.TotalHits(), 10))

	var ttype Shakespeare
	for _, v := range result.Each(reflect.TypeOf(ttype)) {
		if s, ok := v.(Shakespeare); ok {
			fmt.Println("play_name:", s.PlayName, "speaker:", s.Speaker, "text_entry:", s.TextEntry)
		}
	}
}

// NewClient is elastic search client of done setting
func NewClient() (*elastic.Client, error) {
	url := elastic.SetURL(os.Getenv("URL"))
	auth := elastic.SetBasicAuth(os.Getenv("USER"), os.Getenv("PASSWORD"))
	sniff := elastic.SetSniff(false)
	return elastic.NewClient(url, auth, sniff)
}
