package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoInstance struct {
	Client *dynamodb.Client
}

func Connect(ctx context.Context) error {

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))

	if err != nil {
		log.Println("Error loading AWS configuration")
		return err
	}

	DynamoInstance.Client = dynamodb.NewFromConfig(cfg)

	return nil
}
