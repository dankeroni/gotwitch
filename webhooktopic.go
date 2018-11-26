package gotwitch

type WebhookTopic int

const (
	WebhookTopicFollowers WebhookTopic = iota
	WebhookTopicStreams
	WebhookTopicUserChanged
)

func (t WebhookTopic) URL(twitchUserID string) string {
	switch t {
	case WebhookTopicFollowers:
		return "https://api.twitch.tv/helix/users/follows?first=1&to_id=" + twitchUserID
	case WebhookTopicStreams:
		return "https://api.twitch.tv/helix/streams?user_id=" + twitchUserID
	case WebhookTopicUserChanged:
		return "https://api.twitch.tv/helix/users?id=" + twitchUserID
	}

	return ""
}

func (t WebhookTopic) String() string {
	switch t {
	case WebhookTopicFollowers:
		return "followers"
	case WebhookTopicStreams:
		return "streams"
	case WebhookTopicUserChanged:
		return "user_changed"
	}

	return ""
}
