package ch_es

import (
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Log(client.Info())
}

func TestCreateIndex(t *testing.T) {
	response, err := client.Indices.Create("book-0.1.0", client.Indices.Create.WithBody(strings.NewReader(`
	{
		"aliases": {
			"book":{}
		},
		"settings": {
			"analysis": {
				"normalizer": {
					"lowercase": {
						"type": "custom",
						"char_filter": [],
						"filter": ["lowercase"]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"name": {
					"type": "keyword",
					"normalizer": "lowercase"
				},
				"price": {
					"type": "double"
				},
				"summary": {
					"type": "text",
					"analyzer": "ik_max_word"
				},
				"author": {
					"type": "keyword"
				},
				"pubDate": {
					"type": "date"
				},
				"pages": {
					"type": "integer"
				}
			}
		}
	}
	`)))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
