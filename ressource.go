package rin

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

type RestRessources struct{
	Endpoint string
	Method string
	Router *CBRouter
}

func NewRessource(endpoint, method string, router *CBRouter) *RestRessources{
	return &RestRessources{
		Endpoint: endpoint,
		Method: method,
		Router: router,
	}
}

func (r* RestRessources) RenderEndpoint(params map[string]string) string{
	if params == nil {
		return r.Endpoint
	}
	t, err := template.New("ressource").Parse(r.Endpoint)
	if err!=nil {
		log.Fatalln("Unable to parse Endpoint")
	}
	buffer := &bytes.Buffer{}
	t.Execute(buffer, params)
	endpoint, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalln("Unable to read Endpoint")
	}
	return string(endpoint)
}