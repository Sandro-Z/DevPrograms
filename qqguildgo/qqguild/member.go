package qqguild

type Member struct {
	User    *User      `json:"user,omitempty"`
	Nick    string     `json:"nick,omitempty"`
	RoleIDs []RoleID   `json:"roles,omitempty"`
	Joined  *Timestamp `json:"joined_at,omitempty"`
}

func (m Member) Mention() string { return m.User.Mention() }
