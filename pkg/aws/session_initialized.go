package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAwsSessionService(awsKey string, awsSecret string, awsRegion string) *session.Session {
	cred := credentials.NewStaticCredentials(awsKey, awsSecret, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: cred,
	}))
	return sess
}
