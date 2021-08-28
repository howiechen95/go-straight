package netpkg

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func _TestGetUrl(t *testing.T) {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
}
