package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	BaseEndpoint = "https://api.sgroup.qq.com"
	Version      = ""
	Path         = "" + Version

	Endpoint           = BaseEndpoint + Path + "/"
	EndpointGateway    = Endpoint + "gateway"
	EndpointGatewayBot = EndpointGateway + "/bot"
)

var UserAgent = "QQGuildBot"

type Session struct {
	Token     string
	UserAgent string
}

type Client struct {
	*Session
	*http.Client
}

func NewClient(token string) *Client {
	return NewCustomClient(token, &http.Client{})
}

func NewCustomClient(token string, httpClient *http.Client) *Client {
	return &Client{
		Session: &Session{
			Token:     token,
			UserAgent: UserAgent,
		},
		Client: httpClient,
	}
}

func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.Session.Token)
	req.Header.Set("User-Agent", c.Session.UserAgent)
	return req, nil
}

func (c *Client) GetJSON(url string, v any) error {
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNoContent {
		return nil
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s %s: %s", url, resp.Status, p)
	}
	log.Printf("GET %s %s: %s", url, resp.Status, p)
	if v == nil {
		return nil
	}
	return json.Unmarshal(p, v)
}

func (c *Client) PostJSON(url string, data any, v any) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := c.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if len(b) != 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNoContent {
		return nil
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("POST %s %s: %s", url, resp.Status, p)
	}
	log.Printf("POST %s %s: %s", url, resp.Status, p)
	if v == nil {
		return nil
	}
	return json.Unmarshal(p, v)
}

func (c *Client) DeleteJSON(url string, v any) error {
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNoContent {
		return nil
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("DELETE %s %s: %s", url, resp.Status, p)
	}
	if v == nil {
		return nil
	}
	log.Printf("DELETE %s %s: %s", url, resp.Status, p)
	return json.Unmarshal(p, v)
}
