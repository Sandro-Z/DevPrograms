package qqguild

type Channel struct {
	ID      ChannelID `json:"id"`
	GuildID GuildID   `json:"guild_id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Icon    Hash      `json:"icon,omitempty"`

	Type          ChannelType        `json:"type,omitempty"`
	SubType       ChannelSubType     `json:"sub_type,omitempty"`
	ApplicationID ApplicationID      `json:"application_id,omitempty"`
	PreviewType   ChannelPreviewType `json:"private_type,omitempty"`
	MessageType   ChannelMessageType `json:"speak_permission,omitempty"`

	Position int64     `json:"position,omitempty"`
	ParentID ChannelID `json:"parent_id,omitempty"`
	OwnerID  UserID    `json:"owner_id,omitempty"`

	Permissions string `json:"permissions,omitempty"`
}

type ChannelType uint

const (
	ChannelTypeText ChannelType = iota
	_
	ChannelTypeVoice
	_
	ChannelTypeCategory
	ChannelTypeLive ChannelType = 10000 + iota
	ChannelTypeApp
	ChannelTypeThread
)

type ChannelSubType uint

const (
	ChannelSubTypeDefault ChannelSubType = iota
	ChannelSubTypeNotice
	ChannelSubTypeGuide
	ChannelSubTypeGaming
)

type ChannelPreviewType uint

const (
	ChannelPreviewTypePublic ChannelPreviewType = iota
	ChannelPreviewTypeAdmin
	ChannelPreviewTypeMember
)

type ChannelMessageType uint

const (
	_ ChannelMessageType = iota
	ChannelMessageTypeAdmin
	ChannelMessageTypeMember
)
