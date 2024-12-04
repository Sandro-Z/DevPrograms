package api

import (
	"git.ana/xjtuana/qqguildgo/qqguild"
)

type GuildRoles struct {
	GuildID   qqguild.GuildID `json:"guild_id"`
	Roles     []qqguild.Role  `json:"roles"`
	RoleLimit int64           `json:"role_num_limit"`
}

func (c *Client) Roles(guildID qqguild.GuildID) ([]qqguild.Role, error) {
	var grs GuildRoles
	return grs.Roles, c.GetJSON(EndpointGuilds+guildID.String()+"/roles", &grs)
}
