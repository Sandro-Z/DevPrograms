package qqguild

import (
	"strconv"
	"strings"
)

// https://bot.q.qq.com/wiki/develop/api/#id-%E6%8F%8F%E8%BF%B0
type ID uint64

const NullID = ^ID(0)

func (id *ID) UnmarshalJSON(v []byte) error {
	sf := strings.Trim(string(v), `"`)
	if sf == "null" {
		*id = NullID
		return nil
	}
	u, err := strconv.ParseUint(sf, 10, 64)
	if err != nil {
		*id = NullID
		return err
	}
	*id = ID(u)
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	if id == NullID {
		return []byte("null"), nil
	}
	return []byte(`"` + strconv.FormatUint(uint64(id), 10) + `"`), nil
}

func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

type ApplicationID ID

const NullApplicationID = ApplicationID(NullID)

func (id ApplicationID) MarshalJSON() ([]byte, error)  { return ID(id).MarshalJSON() }
func (id *ApplicationID) UnmarshalJSON(v []byte) error { return (*ID)(id).UnmarshalJSON(v) }

func (id ApplicationID) String() string { return ID(id).String() }

type ChannelID ID

const NullChannelID = ChannelID(NullID)

func (id ChannelID) MarshalJSON() ([]byte, error)  { return ID(id).MarshalJSON() }
func (id *ChannelID) UnmarshalJSON(v []byte) error { return (*ID)(id).UnmarshalJSON(v) }

func (id ChannelID) String() string { return ID(id).String() }

type GuildID ID

const NullGuildID = GuildID(NullID)

func (id GuildID) MarshalJSON() ([]byte, error)  { return ID(id).MarshalJSON() }
func (id *GuildID) UnmarshalJSON(v []byte) error { return (*ID)(id).UnmarshalJSON(v) }

func (id GuildID) String() string { return ID(id).String() }

type MessageID = string

type PermissionID ID

const NullPermissionID = PermissionID(NullID)

func (id PermissionID) MarshalJSON() ([]byte, error)  { return ID(id).MarshalJSON() }
func (id *PermissionID) UnmarshalJSON(v []byte) error { return (*ID)(id).UnmarshalJSON(v) }

func (id PermissionID) String() string { return ID(id).String() }

type RoleID ID

const NullRoleID = RoleID(NullID)

func (id RoleID) MarshalJSON() ([]byte, error)  { return ID(id).MarshalJSON() }
func (id *RoleID) UnmarshalJSON(v []byte) error { return (*ID)(id).UnmarshalJSON(v) }

func (id RoleID) String() string { return ID(id).String() }

type UserID ID

const NullUserID = UserID(NullID)

func (id UserID) MarshalJSON() ([]byte, error)  { return ID(id).MarshalJSON() }
func (id *UserID) UnmarshalJSON(v []byte) error { return (*ID)(id).UnmarshalJSON(v) }

func (id UserID) String() string  { return ID(id).String() }
func (id UserID) Mention() string { return "<@!" + id.String() + ">" }
