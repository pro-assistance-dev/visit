package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChatMessage struct {
	bun.BaseModel `bun:"chat_messages,alias:chat_messages"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	User          *User         `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`
	// UserName      string        `bun:"type:uuid" json:"userName"`

	// Chat      *Chat           `bun:"rel:belongs-to" json:"chat"`
	ChatID    uuid.NullUUID   `bun:"type:uuid" json:"chatId"`
	Message   string          `json:"message"`
	Type      ChatMessageType `bun:"-" json:"type"`
	CreatedOn time.Time       `bun:",nullzero,notnull" json:"createdOn"`
}

type ChatMessageType string

const (
	ping  ChatMessageType = "ping"
	enter ChatMessageType = "enter"
	exit  ChatMessageType = "exit"
	// message ChatMessageType = "message"
	// write   ChatMessageType = "write"
)

type ChatMessages []*ChatMessage

func (item *ChatMessage) IsPing() bool {
	return item.Type == ping
}

func (item *ChatMessage) ToBytes() []byte {
	b, _ := json.Marshal(item)
	return b
}
