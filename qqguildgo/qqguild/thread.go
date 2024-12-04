package qqguild

type Post struct {
	GuildID   GuildID   `json:"guild_id"`
	ChannelID ChannelID `json:"channel_id"`
	AuthorID  UserID    `json:"author_id"`
	PostInfo  PostInfo  `json:"post_info"`
}

type PostInfo struct {
	ThreadID string    `json:"thread_id"`
	PostID   string    `json:"post_id"`
	Content  string    `json:"content"`
	Created  Timestamp `json:"date_time"`
}

type Reply struct {
	GuildID   GuildID   `json:"guild_id"`
	ChannelID ChannelID `json:"channel_id"`
	AuthorID  UserID    `json:"author_id"`
	ReplyInfo ReplyInfo `json:"reply_info"`
}

type ReplyInfo struct {
	ThreadID string    `json:"thread_id"`
	PostID   string    `json:"post_id"`
	ReplyID  string    `json:"reply_id"`
	Content  string    `json:"content"`
	Created  Timestamp `json:"date_time"`
}

type Thread struct {
	GuildID    GuildID    `json:"guild_id"`
	ChannelID  ChannelID  `json:"channel_id"`
	AuthorID   UserID     `json:"author_id"`
	ThreadInfo ThreadInfo `json:"thread_info"`
}

type ThreadInfo struct {
	ThreadID string    `json:"thread_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Created  Timestamp `json:"date_time"`
}
