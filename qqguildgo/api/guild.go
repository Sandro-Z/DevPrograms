package api

import (
	"net/url"
	"strconv"

	"git.ana/xjtuana/qqguildgo/qqguild"
)

const MaxGuildFetchLimit = 100

var EndpointGuilds = Endpoint + "guilds/"

func (c *Client) Guild(id qqguild.GuildID) (*qqguild.Guild, error) {
	var g *qqguild.Guild
	return g, c.GetJSON(EndpointGuilds+id.String(), &g)
}

func (c *Client) Guilds(limit uint) ([]qqguild.Guild, error) {
	return c.GuildsAfter(0, limit)
}

func (c *Client) GuildsAfter(after qqguild.GuildID, limit uint) ([]qqguild.Guild, error) {
	guilds := make([]qqguild.Guild, 0, limit)
	fetch := uint(MaxGuildFetchLimit)
	unlimited := limit == 0
	for limit > 0 || unlimited {
		if limit > 0 {
			// Only fetch as much as we need. Since limit gradually decreases,
			// we only need to fetch intmath.Min(fetch, limit).
			fetch = limit
			if fetch > MaxGuildFetchLimit {
				fetch = MaxGuildFetchLimit
			}
			limit -= fetch
		}
		g, err := c.guildsRange(0, after, fetch)
		if err != nil {
			return guilds, err
		}
		guilds = append(guilds, g...)
		// There aren't any to fetch, even if this is less than limit.
		if len(g) < MaxGuildFetchLimit {
			break
		}
		after = g[len(g)-1].ID
	}
	if len(guilds) == 0 {
		return nil, nil
	}
	return guilds, nil
}

func (c *Client) guildsRange(before, after qqguild.GuildID, limit uint) ([]qqguild.Guild, error) {
	params := url.Values{
		"limit": []string{strconv.Itoa(int(limit))},
	}
	if before != 0 {
		params.Set("before", before.String())
	}
	if after != 0 {
		params.Set("after", after.String())
	}
	var gs []qqguild.Guild
	return gs, c.GetJSON(EndpointMe+"/guilds?"+params.Encode(), &gs)
}
