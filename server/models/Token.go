package models

import "github.com/pro-assistance-dev/sprob/helpers/token"

type TokensWithUser struct {
	Tokens *token.Details `json:"tokens"`
	User   User           `json:"user"`
}
