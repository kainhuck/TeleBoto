package telegram

// TODO 目前只支持一个可选参数ParseMode: see https://core.telegram.org/bots/api#sendmessage
type sendMessage struct {
	ChatID              string `json:"chat_id"`
	Text                string `json:"text"`
	ParseMode           string `json:"parse_mode"`
	DisableNotification bool   `json:"disable_notification"`
}

type forwardMessage struct {
	ChatID              string `json:"chat_id"`
	FromChatID          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_notification"`
	MessageID           int    `json:"message_id"`
}

type getChat struct {
	ChatID string `json:"chat_id"`
}

type sendPhoto struct {
	ChatID              string `json:"chat_id"`
	Caption             string `json:"caption"` // 图片标题
	ParseMode           string `json:"parse_mode"`
	DisableNotification bool   `json:"disable_notification"`
}
