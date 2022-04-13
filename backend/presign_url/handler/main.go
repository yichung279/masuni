package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	log "github.com/sirupsen/logrus"

	presignUrlHandler "github.com/yichung279/masuni/backend/presign_url"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load default config: %v", err)
	}

	log.Infof("cfg: %+v", cfg)

	handler, err := presignUrlHandler.New(
		presignUrlHandler.Config{
			AwsConfig: cfg,
			Bucket:    "beauty",
		},
	)
	if err != nil {
		log.Fatalf("failed to create presign url handler: %v", err)
	}

	lambda.StartHandlerWithContext(ctx, handler)
}
