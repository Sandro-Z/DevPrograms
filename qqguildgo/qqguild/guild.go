package qqguild

type Guild struct {
	ID       GuildID `json:"id"`
	Name     string  `json:"name"`
	Icon     URL     `json:"icon"`
	IconHash Hash    `json:"icon_hash,omitempty"`

	Joined  Timestamp `json:"joined_at"`
	Owner   bool      `json:"owner"`
	OwnerID UserID    `json:"owner_id"`

	Description string `json:"description,omitempty"`
	MemberCount int64  `json:"member_count,omitempty"`
	MemberLimit int64  `json:"max_members,omitempty"`
}
