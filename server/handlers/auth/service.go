package auth

import (
	"context"
	"fmt"
	"portal/handlers/users"
	"portal/models"

	"github.com/pro-assistance-dev/sprob/handlers/auth"
	baseModels "github.com/pro-assistance-dev/sprob/models"
)

func (s *Service) Register(c context.Context, email string, password string) (tokenWithUser *models.TokensWithUser, err error) {
	duplicate := false
	item := &models.User{}
	item.UserAccountID, duplicate, err = auth.S.Register(c, email, password)
	if duplicate {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = users.S.Create(c, item)
	if err != nil {
		return nil, err
	}
	ts, err := s.helper.Token.CreateToken(item)
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *item}, nil
}

func (s *Service) Login(c context.Context, item *baseModels.AuthData) (*models.TokensWithUser, error) {
	userAccountID, err := auth.S.Login(c, item)
	fmt.Println(userAccountID)
	if err != nil {
		return nil, err
	}
	user, err := users.S.GetByUserAccountID(c, userAccountID.UUID.String())
	if err != nil {
		return nil, err
	}
	ts, err := s.helper.Token.CreateToken(user)
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *user}, nil
}

func (s *Service) RestorePassword(c context.Context, email string) error {
	// userAccount, err := auth.R.Get(c, email)
	// if err != nil {
	// 	return err
	// }
	// user, err := users.S.GetByUserAccountID(c, userAccount.ID.UUID.String())
	// if err != nil {
	// 	return err
	// }
	//
	// emailStruct := struct {
	// 	RestoreLink string
	// 	Host        string
	// }{
	// 	s.helper.HTTP.GetRestorePasswordURL(user.ID.UUID.String(), userAccount.UUID.String()),
	// 	s.helper.HTTP.Host,
	// }
	//
	// mail, err := s.helper.Templater.ParseTemplate(emailStruct, "email/passwordRestore.gohtml")
	// if err != nil {
	// 	return err
	// }
	// err = s.helper.Email.SendEmail([]string{userAccount.Email}, "Восстановление пароля для сайта Просодействие", mail)
	// if err != nil {
	// 	return err
	// }
	return nil
}
