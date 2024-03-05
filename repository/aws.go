package repository

import (
	"context"
	"gin/config"
	"gin/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type AWS struct {
	conn *session.Session
}

func NewAwsConnection() (*AWS, error) {

	conf := config.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Aws.Region),
		Credentials: credentials.NewStaticCredentials(conf.Aws.AccessKey, conf.Aws.SecretKey, conf.Aws.Token),
	})

	if err != nil {
		return nil, err
	}

	repo := &AWS{
		conn: sess,
	}

	return repo, nil

}

func (a *AWS) VerifyEmail(ctx context.Context, email *string) error {

	conf := config.Get()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.Aws.Region),
		Credentials: credentials.NewStaticCredentials(
			conf.Aws.AccessKey,
			conf.Aws.SecretKey,
			conf.Aws.Token),
	})

	if err != nil {
		return nil
	}

	svc := ses.New(sess)

	emailVerified := &ses.VerifyEmailIdentityInput{
		EmailAddress: aws.String(*email),
	}

	_, err = svc.VerifyEmailIdentity(emailVerified)

	if err != nil {
		return err
	}

	return nil

}

func (a *AWS) SendEmail(ctx context.Context, email *types.Email) error {

	conf := config.Get()

	svc := ses.New(a.conn)

	emailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(email.Recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(email.Body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(email.Subject),
			},
		},
		Source: aws.String(conf.Aws.Source),
	}

	_, err := svc.SendEmail(emailInput)

	if err != nil {
		return err
	}

	return nil

}
