package business

import (
	"context"
	"errors"
	"os"

	"example/auth-services/common"
	"example/auth-services/internal/pkg/auth"
	"example/auth-services/internal/pkg/email"
	"example/auth-services/model"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Register interface {
	GetUserByEmail(ctx context.Context, email string) (user *model.User, err error)
	CreateNewUser(ctx context.Context, data *model.UserInput) error
	GetVertifyCodeByEmail(ctx context.Context, email string) (vertifyCode *string, err error)
	UpdateVertifyCode(ctx context.Context, data *model.VertifyInput) error
	CreateNewCode(ctx context.Context, data *model.VertifyInput) error
}

type registerBusiness struct {
	storage Register
}

func RegisterBusiness(storage Register) *registerBusiness {
	return &registerBusiness{storage: storage}
}

func (r *registerBusiness) RegisterBusiness(ctx context.Context, data *model.UserInput) error {
	err := godotenv.Load("app.env")
	if err != nil {
		return err
	}

	var (
		email_sender_name     = os.Getenv("EMAIL_SENDER_NAME")
		sender_email_address  = os.Getenv("SENDER_EMAIL_ADDRESS")
		sender_email_password = os.Getenv("EMAIL_SENDER_PASSWORD")
	)
	
	user1, err := r.storage.GetUserByEmail(ctx, data.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if user1 != nil {
		return errors.New("email already exist")
	}

	password := data.Hashpassword
	var newHashPassword = common.HashPassword(password)
	data.Hashpassword = newHashPassword
	if err := r.storage.CreateNewUser(ctx, data); err != nil {
		return err
	}

	// generate code
	code, err := auth.CreateVerificationCode()
	if err != nil {
		return err
	}

	content := `
		<h1>Verify your email</h1>
		<p>Enter the following code to verify your email address:</p>
		<h2>` + code + `</h2>
	`
	// send email
	if err := email.NewGmailSender(email_sender_name, sender_email_address, sender_email_password).SendEmail("A test email", content, []string{data.Email}, nil, nil); err != nil {
		return err
	}

	vertifyCode, err := r.storage.GetVertifyCodeByEmail(ctx, data.Email)

	verify := &model.VertifyInput{}
	verify.Email = data.Email
	verify.Code = code


	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound); vertifyCode == nil {
			if err := r.storage.CreateNewCode(ctx, verify); err != nil {
				return err
			}
			return nil

		}
	} else {
		if err := r.storage.UpdateVertifyCode(ctx, verify); err != nil {
			return err
		}
	}


	return nil
}
