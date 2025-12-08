package main

import (
	"context"
	"flag"
	"log"
	"os"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Struct to hold the various dependency injections required
type application struct{
	cfg aws.Config
	logger *slog.Logger
}

func main() {

	// Get command line flags
	awsRegion := flag.String("region", "us-west-1", "AWS region to use e.x. 'us-west-1'")

	// Create a new json logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Get AWS configuration by priority order in AWS SDK
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*awsRegion))
	if err != nil {
		log.Fatalf("unable to load config, %v", err)
	}

	// Initialize the Application struct
	application := &application{
		cfg: cfg,
		logger: logger,
	}

	// Check what credentials are being utilized and log them
	application.awsAuthentication()
}
