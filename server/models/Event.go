package models

import (
	"github.com/google/uuid"
	chatModels "github.com/pro-assistance-dev/sprob/modules/chats/models"
	"github.com/uptrace/bun"
)

type Event struct {
	bun.BaseModel `bun:"events,alias:events"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Description   string        `json:"description"`
	Type          string        `json:"type"`
	Organizer     string        `json:"organizer"`

	Sponsors Sponsors `bun:"m2m:events_to_sponsors,join:Event=Sponsor" json:"sponsors"`
	Perfoms  Perfoms  `bun:"rel:has-many" json:"perfoms"`
	// Perfoms  Perfoms  `bun:"m2m:events_to_perfoms,join:Event=Perfom" json:"perfoms"`

	EventMessages EventMessages `bun:"rel:has-many" json:"eventMessages"`
	EventDays     EventDays     `bun:"rel:has-many" json:"eventDays"`
	Schedules     Schedules     `bun:"rel:has-many" json:"schedules"`
	Archive       bool          `json:"archive"`
	Places        Places        `bun:"m2m:events_to_places,join:Event=Place" json:"places"`

	Chat   *chatModels.Chat[User] `bun:"rel:belongs-to" json:"chat"`
	ChatID uuid.NullUUID          `bun:"type:uuid" json:"chatId"`
}

type Events []*Event

type EventsWithCount struct {
	Events Events `json:"items"`
	Count  int    `json:"count"`
}

//
// type Chat[UserT any] struct {
// 	bun.BaseModel `bun:"chats"`
// 	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
// 	Name          string        `json:"name"`
// 	CreatedOn     time.Time     `bun:",nullzero,notnull" json:"createdOn"`
// 	ChatMessages  ChatMessages  `bun:"rel:has-many,join:id=chat_id" json:"chatMessages"`
// }
//
// type Chats[UserT User] []*Chat[UserT]
//
// type ChatMessage struct {
// 	bun.BaseModel `bun:"chat_messages,alias:chat_messages"`
// 	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
// 	// User          UserT           `bun:"rel:belongs-to" json:"user"`
// 	UserID    uuid.NullUUID   `bun:"type:uuid" json:"userId"`
// 	Chat      *Chat[User]     `bun:"rel:belongs-to,join:chat_id=id" json:"chat"`
// 	ChatID    uuid.NullUUID   `bun:"type:uuid" json:"chatId"`
// 	Message   string          `json:"message"`
// 	Type      ChatMessageType `bun:"-" json:"type"`
// 	CreatedOn time.Time       `bun:",nullzero,notnull" json:"createdOn"`
// }
//
// type ChatMessageType string
//
// const (
// 	ping    ChatMessageType = "ping"
// 	join    ChatMessageType = "join"
// 	exit    ChatMessageType = "exit"
// 	message ChatMessageType = "message"
// 	// write   ChatMessageType = "write"
// )
//
// type ChatMessages []*ChatMessage
