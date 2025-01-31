package models

import (
	"github.com/google/uuid"
	baseModels "github.com/pro-assistance-dev/sprob/models"
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"menus,alias:menus"`
	ID            uuid.UUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string               `json:"name"`
	Link          string               `json:"link"`
	Top           bool                 `json:"top"`
	Side          bool                 `json:"side"`
	Order         uint                 `bun:"menu_order" json:"order"`
	Icon          *baseModels.FileInfo `bun:"rel:belongs-to" json:"icon"`
	IconID        uuid.NullUUID        `bun:"type:uuid"  json:"iconId"`
	Hide          bool                 `json:"hide"`
}

type Menus []*Menu

type MenusWithCount struct {
	Menus Menus `json:"items"`
	Count int   `json:"count"`
}

func (items Menus) GetIcons() baseModels.FileInfos {
	itemsForGet := make(baseModels.FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Icon)
	}
	return itemsForGet
}

func (item *Menu) SetForeignKeys() {
	item.IconID = item.Icon.ID
}

func (items Menus) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}
