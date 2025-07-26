package rin

import (
	"net/http"
	"testing"
)

func TestAPICall(t *testing.T){
	api := NewApi("http://httpbin.org")
	router:=NewRouter()
	router.RegisterFunc(200, func(resp *http.Response,  _ interface{}) error {
		return nil
	})
	res := NewRessource("/get", "GET", router)
	api.AddRessource("get", res)
	if err := api.Call("get", nil); err!=nil{
		t.Fail()
	}

}