package rin

import (
	"net/http"
	"testing"
)

func TestProcessRequest(t *testing.T) {
	client := NewClient()
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, content interface{}) error {
		return nil;
	})
	res := NewRessource("/get", "GET", router)
	if err := client.ProcessRequest("http://httpbin.org", res, nil, nil); err != nil{ t.Fail()}
}