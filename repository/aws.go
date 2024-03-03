package repository

import (
	"gin/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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
