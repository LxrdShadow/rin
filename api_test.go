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
		resources := api.RessourceNames()
	if len(resources) != 1 || resources[0] != "get" {
		t.Fail()
	}

}