package business

import (
	"context"
	"errors"

	"example/auth-services/internal/pkg/auth"
	"example/auth-services/internal/pkg/email"
	"example/auth-services/model"

	"gorm.io/gorm"
)

type SendCode interface {
	GetVertifyCodeByEmail(ctx context.Context, email string) (vertifyCode *string, err error)
	UpdateVertifyCode(ctx context.Context, data *model.VertifyInput) error
	CreateNewCode(ctx context.Context, data *model.VertifyInput) error
}

type sendCodeBusiness struct {
	storage SendCode
}

func SendCodeBusiness(storage SendCode) *sendCodeBusiness {
	return &sendCodeBusiness{storage: storage}
}

func (s *sendCodeBusiness) SendCode(ctx context.Context, data *model.VertifyInput) error {
	// generate code
	code, err := auth.CreateVerificationCode()
	if err != nil {
		return err
	}

	data.Code = code

	content := `
		<h1>Verify your email</h1>
		<p>Enter the following code to verify your email address:</p>
		<h2>` + code + `</h2>
	`
	// send email
	if err := email.NewGmailSender("AUTH GO TEST", "mailyhai814@gmail.com", "xnog anph lvqy tfga").SendEmail("A test email", content, []string{data.Email}, nil, nil); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	vertifyCode, err := s.storage.GetVertifyCodeByEmail(ctx, data.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound); vertifyCode == nil {
			if err := s.storage.CreateNewCode(ctx, data); err != nil {
				return err
			}
			return nil
			
		}
	}else {
		if err := s.storage.UpdateVertifyCode(ctx, data); err != nil {
			return err
		}
	}

	

	return nil
}
