package models

import (
	"strings"

	"github.com/pro-assistance-dev/sprob/middleware"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"users,select:users_view,alias:users_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	CheckUUID     uuid.UUID     `bun:"check_uuid,type:uuid,nullzero,notnull,default:uuid_generate_v4()"  json:"uuid"` // для восстановления пароля - обеспечивает уникальность страницы на фронте
	LastName      string        `json:"lastName"`
	FirstName     string        `json:"firstName"`
	PatronName    string        `json:"patronName"`
	Company       string        `json:"company"`
	Post          string        `json:"post"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone"`
	Password      string        `json:"password"`
	Role          *Role         `bun:"rel:belongs-to" json:"role"`
	RoleID        uuid.NullUUID `bun:"type:uuid" json:"roleId"`
	UserAccountID uuid.NullUUID `bun:"type:uuid" json:"-"`

	FullName string `bun:"-" json:"fullName"`
}

type (
	Users          []*User
	UsersWithCount struct {
		Users Users `json:"items"`
		Count int   `json:"count"`
	}
)

func (u *User) Normalize() {
	u.Email = strings.ToLower(u.Email)
}

// func (i *User) SetForeignKeys() {
// 	if i.Role != nil && i.Role.ID.Valid {
// 		i.RoleID = i.Role.ID
// 	}
// }

func (u *User) GenerateHashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	pass := string(hash)
	u.Password = pass
	return nil
}

func (u *User) CompareWithHashPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) CompareWithUUID(externalUUID string) bool {
	return u.CheckUUID.String() == externalUUID
}

func (u *User) SetJWTClaimsMap(claims map[string]interface{}) {
	claims[middleware.ClaimUserID.String()] = u.ID.UUID
}
