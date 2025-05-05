package main

import (
	"github.com/saesetiawan/aws-go/pkg/aws"
	"github.com/saesetiawan/aws-go/pkg/slack"
)

func main() {
	session := aws.NewAwsSessionService("aws key", "aws secret", "aws region")
	cloudWatchLogs := aws.NewCloudWatchLogsService(session)
	slackService := slack.NewSlackService("", "development", false)
	cloudWatchService := aws.NewAwsCloudWatchServiceImpl("group name", "stream name", false, cloudWatchLogs, slackService)
	cloudWatchService.Info("Success send to cloud watch")
}
