package telegram

// 注意：不支持官方文档中的递归定义，这会导致不能百分百的序列化结果

type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type Chat struct {
	ID                          int             `json:"id"`
	Type                        string          `json:"type"`
	Title                       string          `json:"title"`
	Username                    string          `json:"username"`
	FirstName                   string          `json:"first_name"`
	LastName                    string          `json:"last_name"`
	Photo                       ChatPhoto       `json:"photo"`
	Description                 string          `json:"description"`
	InviteLink                  string          `json:"invite_link"`
	Permissions                 ChatPermissions `json:"permissions"`
	SlowModeDelay               int             `json:"slow_mode_delay"`
	StickerSetName              string          `json:"sticker_set_name"`
	CanSetStickerSet            bool            `json:"can_set_sticker_set"`
	AllMembersAreAdministrators bool            `json:"all_members_are_administrators"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

// TODO: 增加其他类型https://core.telegram.org/bots/api#message
type Message struct {
	MessageID            int             `json:"message_id"`
	From                 User            `json:"from"`
	Date                 int             `json:"date"`
	Chat                 Chat            `json:"chat"`
	ForwardFrom          User            `json:"forward_from"`
	ForwardFromChat      Chat            `json:"forward_from_chat"`
	ForwardFromMessageID int             `json:"forward_from_message_id"`
	ForwardSignature     string          `json:"forward_signature"`
	ForwardSenderName    string          `json:"forward_sender_name"`
	ForwardDate          int             `json:"forward_date"`
	ViaBot               User            `json:"via_bot"`
	EditDate             int             `json:"edit_date"`
	MediaGroupID         string          `json:"media_group_id"`
	AuthorSignature      string          `json:"author_signature"`
	Text                 string          `json:"text"`
	NewChatMembers       []User          `json:"new_chat_members"`
	Entities             []MessageEntity `json:"entities"`
	Photo                []PhotoSize     `json:"photo"`
	Caption              string          `json:"caption"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	Url      string `json:"url"`
	User     User   `json:"user"`
	Language string `json:"language"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// TODO 增加别的字段
type Update struct {
	UpdateID          int     `json:"update_id"`
	Message           Message `json:"message"`
	EditedMessage     Message `json:"edited_message"`
	ChannelPost       Message `json:"channel_post"`
	EditedChannelPost Message `json:"edited_channel_post"`
}

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     string `json:"file_path"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}
