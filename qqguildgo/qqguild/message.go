package qqguild

type MessageReference struct {
	MessageID MessageID `json:"message_id,omitempty"`
	ChannelID ChannelID `json:"channel_id,omitempty"`
	GuildID   GuildID   `json:"guild_id,omitempty"`
}

type Message struct {
	ID        MessageID `json:"id"`
	ChannelID ChannelID `json:"channel_id"`
	GuildID   GuildID   `json:"guild_id"`
	Seq       int64     `json:"seq"`

	DirectMessage   bool   `json:"direct_message"`
	MentionEveryone bool   `json:"mention_everyone"`
	Mentions        []User `json:"mentions,omitempty"`

	Author  *User  `json:"author"`
	Content string `json:"content"`

	Timestamp       *Timestamp `json:"timestamp,omitempty"`
	EditedTimestamp *Timestamp `json:"edited_timestamp,omitempty"`

	Attachments []Attachment `json:"attachments,omitempty"`
	Embeds      []Embed      `json:"embeds,omitempty"`
	Reactions   []Reaction   `json:"reactions,omitempty"`

	ReferencedMessage *Message `json:"message_reference,omitempty"`

	// Ark Ark `json:"ark"`
}

type Attachment struct {
	URL string `json:"url,omitempty"`
}

type Embed struct {
	Title     string         `json:"title,omitempty"`
	Prompt    string         `json:"prompt,omitempty"`
	Thumbnail EmbedThumbnail `json:"thumbnail,omitempty"`
	Fields    []EmbedField   `json:"fields,omitempty"`
}

type EmbedThumbnail struct {
	URL string `json:"url,omitempty"`
}

type EmbedField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

type Emoji struct {
	ID   string `json:"id"`
	Type uint32 `json:"type"`
}
