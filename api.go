package rin

type API struct{
	BaseUrl string
	Ressources map[string]RestRessources
	DefaultRouter *CBRouter
	client *Client
}