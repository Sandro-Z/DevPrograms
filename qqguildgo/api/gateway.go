package api

type Gateway struct {
	URL string `json:"url"`
}

func (c *Client) GatewayBotURL(token string) (string, error) {
	var g Gateway
	return g.URL, c.GetJSON(EndpointGatewayBot, &g)
}

func (c *Client) GatewayURL() (string, error) {
	var g Gateway
	return g.URL, c.GetJSON(EndpointGateway, &g)
}
