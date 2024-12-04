package qqguild

type User struct {
	ID       UserID `json:"id"`
	Username string `json:"username"`
	Avatar   URL    `json:"avatar,omitempty"`

	Bot bool `json:"bot,omitempty"`

	UnionOpenID      string `json:"union_openid,omitempty"`
	UnionUserAccount string `json:"union_user_account,omitempty"`
}

func (u User) Mention() string { return u.ID.Mention() }
