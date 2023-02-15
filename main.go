package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/skpr/guardduty-slack-lambda/internal/guardduty"
	"github.com/skpr/guardduty-slack-lambda/internal/slack"
	"github.com/skpr/guardduty-slack-lambda/internal/util"
)

var (
	// GitVersion overridden at build time by:
	//   -ldflags="-X main.GitVersion=${VERSION}"
	GitVersion string
)

func main() {
	lambda.Start(HandleLambdaEvent)
}

// HandleLambdaEvent will respond to a CloudWatch Alarm, check for rate limited IP addresses and send a message to Slack.
func HandleLambdaEvent(event guardduty.Event) error {
	log.Printf("Running Lambda (%s)\n", GitVersion)

	config, err := util.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	errs := config.Validate()
	if len(errs) > 0 {
		return fmt.Errorf("configuration error: %s", strings.Join(errs, "\n"))
	}

	log.Println("Sending Slack message for finding ID:", event.Detail.ID)

	err = slack.PostMessage(config, event.Detail)
	if err != nil {
		return fmt.Errorf("failed to post Slack message: %w", err)
	}

	log.Println("Function complete")

	return nil
}
