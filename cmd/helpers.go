package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func awsAuthentication(cfg aws.Config) {
	// Create an sts client to log which account the user is authenticating to
	stsClient := sts.NewFromConfig(cfg)

	identity, err := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("failed to get caller identity, %v", err)
	}

	fmt.Printf("Account: %s\n", *identity.Account)
	fmt.Printf("UserID: %s\n", *identity.UserId)
	fmt.Printf("ARN: %s\n", *identity.Arn)
}
