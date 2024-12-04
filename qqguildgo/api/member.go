package api

import (
	"net/url"
	"strconv"

	"git.ana/xjtuana/qqguildgo/qqguild"
)

const MaxMemberFetchLimit = 1000

func (c *Client) Member(guildID qqguild.GuildID, userID qqguild.UserID) (*qqguild.Member, error) {
	var m *qqguild.Member
	return m, c.GetJSON(EndpointGuilds+guildID.String()+"/members/"+userID.String(), &m)
}

func (c *Client) Members(guildID qqguild.GuildID, limit uint) ([]qqguild.Member, error) {
	return c.MembersAfter(guildID, 0, limit)
}

func (c *Client) MembersAfter(guildID qqguild.GuildID, after qqguild.UserID, limit uint) ([]qqguild.Member, error) {
	ms := make([]qqguild.Member, 0, limit)
	fetch := uint(MaxMemberFetchLimit)
	unlimited := limit == 0
	for limit > 0 || unlimited {
		// Only fetch as much as we need. Since limit gradually decreases,
		// we only need to fetch intmath.Min(fetch, limit).
		if limit > 0 {
			fetch = limit
			if fetch > MaxMemberFetchLimit {
				fetch = MaxMemberFetchLimit
			}
			limit -= fetch
		}
		m, err := c.membersAfter(guildID, after, fetch)
		if err != nil {
			return ms, err
		}
		ms = append(ms, m...)
		// There aren't any to fetch, even if this is less than limit.
		if len(m) < MaxMemberFetchLimit {
			break
		}
		after = ms[len(ms)-1].User.ID
	}
	if len(ms) == 0 {
		return nil, nil
	}
	return ms, nil
}

func (c *Client) membersAfter(guildID qqguild.GuildID, after qqguild.UserID, limit uint) ([]qqguild.Member, error) {
	params := url.Values{
		"limit": []string{strconv.Itoa(int(limit))},
	}
	if after != 0 {
		params.Set("after", after.String())
	}
	var ms []qqguild.Member
	return ms, c.GetJSON(EndpointGuilds+guildID.String()+"/members?"+params.Encode(), &ms)
}
