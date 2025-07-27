package rin

import "fmt"

type API struct{
	BaseUrl string
	Ressources map[string]*RestRessources
	DefaultRouter *CBRouter
	Client *Client
}

func (api *API) SetAuth(auth Authentication) {
	api.Client.SetAuth(auth)
}

func NewApi(baseURL string) *API {
	return &API{
		BaseUrl: baseURL,
		Ressources: make(map[string]*RestRessources),
		DefaultRouter: NewRouter(),
		Client: NewClient(),
	}
}

func (api *API) AddRessource(name string, res *RestRessources){
	api.Ressources[name] = res
}

func (api *API) Call(name string, params map[string]string, payload interface{}) error{
	res, ok:=api.Ressources[name]
	if !ok {
		return fmt.Errorf("Ressources does not exist %s", name)
	}
	if err:=api.Client.ProcessRequest(api.BaseUrl, res, params, payload); err != nil{
		return err
	}
	return nil
}

func (api *API) RessourceNames() []string {
	ressources := []string{}
	for k := range api.Ressources{
		ressources = append(ressources, k)
	}
	return ressources
}