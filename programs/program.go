package programs

import (
	"com.talreg.awstests/utils"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"os"
)

type Program struct {
	cfg                  aws.Config
	iamRoleArn           string
	webIdentityTokenFile string
	roleProvider         *stscreds.WebIdentityRoleProvider
	selectedRegion       string
}

func NewProgram() *Program {
	return &Program{
		selectedRegion: "ca-central-1",
	}
}

func (p *Program) Initialize() error {
	var err error
	if p.cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion(p.selectedRegion)); err != nil {
		return err
	}
	if err = p.getCredentials(); err != nil {
		return err
	}
	return nil
}

func (p *Program) getCredentials() error {
	p.iamRoleArn = os.Getenv("AWS_ROLE_ARN")
	p.webIdentityTokenFile = os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE")
	stsClient := sts.NewFromConfig(p.cfg)
	tokenRetriever := &utils.FileTokenRetriever{
		TokenFilePath: p.webIdentityTokenFile,
	}
	p.roleProvider = stscreds.NewWebIdentityRoleProvider(stsClient, p.iamRoleArn, tokenRetriever)
	p.cfg.Credentials = aws.NewCredentialsCache(p.roleProvider)
	return nil
}

func (p *Program) Execute() (string, error) {
	secretName := os.Getenv("SELECTED_SECRET_NAME")
	secretsClient := secretsmanager.NewFromConfig(p.cfg)
	data, err := secretsClient.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return "", err
	} else {
		return *data.SecretString, nil
	}
}
