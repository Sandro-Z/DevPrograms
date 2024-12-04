package qqguild

const (
	RoleIDEveryone = RoleID(1)
	RoleIDAdmin    = RoleID(2)
	RoleIDOwner    = RoleID(4)
	RoleIDSubAdmin = RoleID(5)
)

type Role struct {
	ID          RoleID `json:"id"`
	Name        string `json:"name"`
	Color       uint32 `json:"color"`
	Hoist       uint32 `json:"hoist"`
	MemberCount int64  `json:"number,omitempty"`
	MemberLimit int64  `json:"member_limit,omitempty"`
}
