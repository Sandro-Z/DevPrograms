package gateway

import (
	"github.com/diamondburned/arikawa/v3/utils/ws"
)

type Intents uint32

const (
	IntentGuilds                Intents = 1 << 0  // GUILDS
	IntentGuildMembers          Intents = 1 << 1  // GUILD_MEMBERS
	IntentGuildMessages         Intents = 1 << 9  // GUILD_MESSAGES
	IntentGuildMessageReactions Intents = 1 << 10 // GUILD_MESSAGE_REACTIONS
	IntentDirectMessages        Intents = 1 << 12 // DIRECT_MESSAGE
	IntentInteraction           Intents = 1 << 26 // INTERACTION
	IntentMessageAudit          Intents = 1 << 27 // MESSAGE_AUDIT
	IntentForumsEvent           Intents = 1 << 28 // FORUMS_EVENT
	IntentAudioAction           Intents = 1 << 29 // AUDIO_ACTION
	IntentAtMessage             Intents = 1 << 30 // AT_MESSAGES
)

const IntentsAll = IntentGuilds | IntentGuildMembers | IntentGuildMessages | IntentGuildMessageReactions | IntentDirectMessages | IntentInteraction | IntentMessageAudit | IntentForumsEvent | IntentAudioAction | IntentAtMessage

var EventIntents = map[ws.EventType]Intents{
	"GUILD_CREATE":   IntentGuilds,
	"GUILD_UPDATE":   IntentGuilds,
	"GUILD_DELETE":   IntentGuilds,
	"CHANNEL_CREATE": IntentGuilds,
	"CHANNEL_UPDATE": IntentGuilds,
	"CHANNEL_DELETE": IntentGuilds,

	"GUILD_MEMBER_ADD":    IntentGuildMembers,
	"GUILD_MEMBER_REMOVE": IntentGuildMembers,
	"GUILD_MEMBER_UPDATE": IntentGuildMembers,

	"MESSAGE_CREATE": IntentGuildMessages,

	"MESSAGE_REACTION_ADD":    IntentGuildMessageReactions,
	"MESSAGE_REACTION_REMOVE": IntentGuildMessageReactions,

	"DIRECT_MESSAGE_CREATE": IntentDirectMessages,

	"INTERACTION_CREATE": IntentInteraction,

	"MESSAGE_AUDIT_PASS":   IntentMessageAudit,
	"MESSAGE_AUDIT_REJECT": IntentMessageAudit,

	"FORUM_THREAD_CREATE":        IntentForumsEvent,
	"FORUM_THREAD_UPDATE":        IntentForumsEvent,
	"FORUM_THREAD_DELETE":        IntentForumsEvent,
	"FORUM_POST_CREATE":          IntentForumsEvent,
	"FORUM_POST_DELETE":          IntentForumsEvent,
	"FORUM_REPLY_CREATE":         IntentForumsEvent,
	"FORUM_REPLY_DELETE":         IntentForumsEvent,
	"FORUM_PUBLISH_AUDIT_RESULT": IntentForumsEvent,

	"AUDIO_START":   IntentAudioAction,
	"AUDIO_FINISH":  IntentAudioAction,
	"AUDIO_ON_MIC":  IntentAudioAction,
	"AUDIO_OFF_MIC": IntentAudioAction,

	"AT_MESSAGE_CREATE": IntentAtMessage,
}
