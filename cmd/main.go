package main

import (
	"context"
	"flag"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
)

type Application struct{
	cfg config.Config
}

func main() {

	// Get command line flags
	awsRegion := flag.String("region", "us-west-1", "AWS region to use e.x. 'us-west-1'")

	// Get AWS configuration by priority order in AWS SDK
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*awsRegion))
	if err != nil {
		log.Fatalf("unable to load config, %v", err)
	}

	// Login using local AWS credentials and log
	awsAuthentication(cfg)
}
