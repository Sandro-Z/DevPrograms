package api

import "git.ana/xjtuana/qqguildgo/qqguild"

var (
	EndpointUsers = Endpoint + "users/"
	EndpointMe    = EndpointUsers + "@me"
)

func (c *Client) User(userID qqguild.UserID) (*qqguild.User, error) {
	var u *qqguild.User
	return u, c.GetJSON(EndpointUsers+userID.String(), &u)
}

func (c *Client) Me() (*qqguild.User, error) {
	var me *qqguild.User
	return me, c.GetJSON(EndpointMe, &me)
}
