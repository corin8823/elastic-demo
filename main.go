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
	PlayName     string `json:"play_name,omitempty"`
	Speaker      string `json:"speaker,omitempty"`
	SpeechNumber int    `json:"speech_number,omitempty"`
	TextEntry    string `json:"text_entry,omitempty"`
}

func main() {
	url := elastic.SetURL(os.Getenv("URL"))
	auth := elastic.SetBasicAuth(os.Getenv("USER"), os.Getenv("PASSWORD"))
	sniff := elastic.SetSniff(false)
	client, err := elastic.NewClient(url, auth, sniff)
	if err != nil {
		fmt.Println(err.Error())
	}

	q := elastic.NewTermQuery("_index", "shakespeare")
	result, err2 := client.Search().Index("shakespeare").Query(q).Do(context.Background())
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(strconv.FormatInt(result.TotalHits(), 10))

	var ttype Shakespeare
	for _, v := range result.Each(reflect.TypeOf(ttype)) {
		if s, ok := v.(Shakespeare); ok {
			fmt.Println("play_name:", s.PlayName, "speaker:", s.Speaker)
		}
	}
}
