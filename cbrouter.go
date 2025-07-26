package rin

import "net/http"

type RouterFunc func (client *http.Client, content interface{})
type CBRouter struct {
	Routers map[string]RouterFunc
	DefaultRouter RouterFunc
}