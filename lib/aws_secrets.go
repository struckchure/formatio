package lib

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/joho/godotenv"
)

// TODO: refector
func writeTextToFile(filename string, text string) error {
	// Get the directory part of the filename
	dir := filepath.Dir(filename)

	// Create the directory path if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Open the file for writing. Create it if it doesn't exist, or truncate it if it does.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

type AwsSecrets struct {
	ctx *context.Context
}

var (
	AWS_SECRET_NAME       = os.Getenv("AWS_SECRET_NAME")
	AWS_SECRET_VERSION    = os.Getenv("AWS_SECRET_VERSION")
	AWS_REGION            = os.Getenv("AWS_REGION")
	AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
)

// Load implements SecretsInterface.
func (a AwsSecrets) Load() error {
	// Create Secrets Manager client
	svc := secretsmanager.New(secretsmanager.Options{
		Region: AWS_REGION,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, "")),
	})

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(AWS_SECRET_NAME),
		VersionStage: aws.String(AWS_SECRET_VERSION),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())

		return err
	}

	writeTextToFile(".env", *result.SecretString)
	godotenv.Load(".env")

	return nil
}

func (a AwsSecrets) Get(key string) string {
	return os.Getenv(key)
}

// GetOrPanic implements SecretsInterface.
func (a AwsSecrets) GetOrPanic(key string) string {
	value := a.Get(key)
	if value == "" {
		log.Panicf("%s is empty", key)
	}

	return value
}

func (*AwsSecrets) NewAwsSecret() ISecrets {
	ctx := context.Background()

	return AwsSecrets{
		ctx: &ctx,
	}
}
