package telegram

import (
	"encoding/json"
	"errors"
)

func (t *TeleBot) Fetch(content interface{}) *TeleBot {
	var data = struct {
		OK     bool        `json:"ok"`
		Result interface{} `json:"result"`
	}{}
	err := json.Unmarshal(t.Body, &data)
	if err != nil {
		t.Errors = append(t.Errors, err)
	}
	if !data.OK {
		t.Errors = append(t.Errors, errors.New("data not ok fetch error"))
		return t
	}

	bytes, err := json.Marshal(data.Result)
	if err != nil {
		t.Errors = append(t.Errors, err)
		return t
	}
	err = json.Unmarshal(bytes, &content)
	if err != nil {
		t.Errors = append(t.Errors, err)
		return t
	}

	return t
}

func (t *TeleBot) FetchUser() (User, []error) {
	var user User
	t.Fetch(&user)
	return user, t.Errors
}

func (t *TeleBot) FetchMessage() (Message, []error) {
	var message Message
	t.Fetch(&message)
	return message, t.Errors
}

func (t *TeleBot) FetchUpdates() ([]Update, []error) {
	var updates []Update
	t.Fetch(&updates)
	return updates, t.Errors
}

func (t *TeleBot) FetchChat() (Chat, []error) {
	var chat Chat
	t.Fetch(&chat)
	return chat, t.Errors
}
