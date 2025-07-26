package rin

import (
	"net/http"
	"strings"
)

type Client struct{
	Client *http.Client
	AuthInfo Authentication
}

func NewClient() *Client{
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) SetAuth(auth Authentication){
	c.AuthInfo = auth
}

func (c *Client) ProcessRequest(baseURL string, res *RestRessources, params map[string]string) error{
	endpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")
	trimedBaseURL := strings.TrimRight(baseURL, "/")
	url := trimedBaseURL +"/" + endpoint
	req, err:=http.NewRequest(res.Method, url, nil)
	if err != nil {
		return err
	}
	if c.AuthInfo != nil{
		req.Header.Add("Authorization", c.AuthInfo.AuthorizationHeader())
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if err := res.Router.CallFunc(resp, nil); err != nil{
		return err
	}
	return nil
}