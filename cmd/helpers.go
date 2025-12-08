package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Takes in the aws.Config 
func(app *application) awsAuthentication() {
	// Create an sts client to log which account the user is authenticating to
	stsClient := sts.NewFromConfig(app.cfg)

	identity, err := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("failed to get caller identity, %v", err)
	}

	app.logger.Info("Caller information", "AccountID", *identity.Account, "UserID", *identity.UserId, "ARN", *identity.Arn)
}
