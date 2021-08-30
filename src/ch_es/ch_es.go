package ch_es

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"log"
)

var (
	client *es.Client
)

func init() {
	var err error
	client, err = es.NewClient(es.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "username",
		Password:  "password",
	})
	if err != nil {
		log.Fatal(err)
	}
}
