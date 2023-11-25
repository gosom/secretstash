package awssecret

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	"github.com/gosom/secretstash"
)

var _ secretstash.SecretProvider = (*provider)(nil)

type provider struct {
	svc *secretsmanager.SecretsManager
}

func New(region string) secretstash.SecretProvider {
	sess := session.Must(session.NewSession())
	svc := secretsmanager.New(
		sess,
		aws.NewConfig().WithRegion(region),
	)

	ans := provider{
		svc: svc,
	}

	return &ans
}

func (p *provider) GetSecret(name string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(name),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := p.svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	if result.SecretString == nil {
		return "", secretstash.ErrSecretNotFound
	}

	return *result.SecretString, nil
}
