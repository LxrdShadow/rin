package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Rakotoarilala51/rin"
)
var api = rin.NewApi("https://httpbin.org")
func main(){
	list:=flag.Bool("list", false, "Get list of api-ressource")
	flag.Parse()
	if *list {
		fmt.Println("Available Ressources: ")
		for _, name:= range api.RessourceNames(){
			fmt.Println(name)
		}
		return
	}
	ressource := os.Args[1]
	if err:=api.Call(ressource, nil, nil); err != nil{
		log.Fatalln(err)
	}
}

func init(){
	router := rin.NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			return err
		}
		fmt.Println(string(content))
		return nil
	})
	api.AddRessource("get", rin.NewRessource("/get", "GET", router))

}	