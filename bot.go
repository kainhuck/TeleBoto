package telegram

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type TeleBot struct {
	baseUrl             string
	sender              *gorequest.SuperAgent
	chatID              string
	Errors              []error
	Body                []byte
	Response            gorequest.Response
	disableNotification bool // 发送消息是否禁用通知
}

func New() *TeleBot {
	return &TeleBot{
		sender: gorequest.New(),
	}
}

func Create(token, chatID string) *TeleBot {
	t := &TeleBot{
		sender: gorequest.New(),
	}
	return t.SetToken(token).SetChatID(chatID)
}

// ============== API ================
// GetUpdates 获取更新
func (t *TeleBot) GetUpdates() *TeleBot {
	return t.simpleSend("/getUpdates")
}

// GetMe 返回机器人自己的信息
func (t *TeleBot) GetMe() *TeleBot {
	return t.simpleSend("/getMe")
}

func (t *TeleBot) GetChat() *TeleBot {
	return t.receive(t.sender.Post(t.baseUrl + "/getChat").SendStruct(getChat{
		ChatID: t.chatID,
	}).EndBytes())
}

// Send 这是一个发送文字的原始方法. chatID 为 int 类型或者string
func (t *TeleBot) SendText(chatID, text, parseMode string) *TeleBot {
	if len(chatID) == 0 {
		t.Errors = append(t.Errors, errors.New("missing chat_id"))
		return t
	}
	return t.receive(t.sender.Post(t.baseUrl + "/sendMessage").SendStruct(sendMessage{
		ChatID:              chatID,
		Text:                text,
		ParseMode:           parseMode,
		DisableNotification: t.disableNotification,
	}).EndBytes())
}

// SendPlain 发送无格式的文本，需要先通过 SetChatID 设置chat_id
func (t *TeleBot) SendPlain(text string) *TeleBot {
	return t.SendText(t.chatID, text, "")
}

// SendMarkdown 发送markdown格式文本
func (t *TeleBot) SendMarkdown(text string) *TeleBot {
	return t.SendText(t.chatID, text, "Markdown")
}

// SendHTML 发送HTML格式文本
func (t *TeleBot) SendHTML(text string) *TeleBot {
	return t.SendText(t.chatID, text, "HTML")
}

// ForwardMessage 转发信息
func (t *TeleBot) ForwardMessage(chatID, fromChatID string, messageID int) *TeleBot {
	if len(chatID)*len(fromChatID)*messageID <= 0 {
		t.Errors = append(t.Errors, errors.New("missing chat_id , from_chat_id or message_id"))
		return t
	}
	return t.receive(t.sender.Post(t.baseUrl + "/forwardMessage").SendStruct(forwardMessage{
		ChatID:              chatID,
		FromChatID:          fromChatID,
		DisableNotification: t.disableNotification,
		MessageID:           messageID,
	}).EndBytes())
}

func (t *TeleBot) ForwardMessageFrom(fromChatID string, messageID int) *TeleBot {
	return t.ForwardMessage(t.chatID, fromChatID, messageID)
}

func (t *TeleBot) SendPhoto(photo string, caption string) *TeleBot {

	// 判断photo是否是本地路径
	if !Exists(photo) {
		t.sender.Post(t.baseUrl + "/sendPhoto").Send(map[string]interface{}{
			"photo": photo,
		})
	} else {
		// FIXME 还不能发送本地照片
		t.sender.Post(t.baseUrl+"/sendPhoto").
			Type("multipart").
			SendFile(photo, "photo")
	}

	return t.receive(t.sender.Send(sendPhoto{
		ChatID:              t.chatID,
		Caption:             caption,
		ParseMode:           "",
		DisableNotification: t.disableNotification,
	}).EndBytes())
}

// =============== config ===============
func (t *TeleBot) SetProxy(proxyUrl string) *TeleBot {
	t.sender.Proxy(proxyUrl)
	return t
}

func (t *TeleBot) UseDefaultProxy() *TeleBot {
	return t.SetProxy("socks5://127.0.0.1:1086")
}

func (t *TeleBot) SetChatID(chatID string) *TeleBot {
	t.chatID = chatID
	return t
}

func (t *TeleBot) SetToken(token string) *TeleBot {
	t.baseUrl = fmt.Sprintf("https://api.telegram.org/bot%s", token)
	return t
}

func (t *TeleBot) SetDisableNotification(b bool) *TeleBot {
	t.disableNotification = b
	return t
}

func (t *TeleBot) DisableNotification() *TeleBot {
	return t.SetDisableNotification(true)
}

func (t *TeleBot) EnableNotification() *TeleBot {
	return t.SetDisableNotification(false)
}

func (t *TeleBot) Clear() *TeleBot {
	t.sender.ClearSuperAgent()
	return t
}

// ============== inner funcs =============
func (t *TeleBot) receive(response gorequest.Response, body []byte, errs []error) *TeleBot {
	t.Response, t.Body, t.Errors = response, body, errs
	return t
}

func (t *TeleBot) simpleSend(pattern string) *TeleBot {
	return t.receive(t.sender.Post(t.baseUrl + pattern).EndBytes())
}
