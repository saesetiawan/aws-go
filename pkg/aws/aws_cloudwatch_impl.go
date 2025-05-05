package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/gofiber/fiber/v2/log"
	"github.com/saesetiawan/aws-go/pkg/helpers"
	"github.com/saesetiawan/aws-go/pkg/slack"
	"time"
)

type CloudWatchServiceImpl struct {
	Logging      *cloudwatchlogs.CloudWatchLogs
	SlackService slack.SlackService
	SendToLog    bool
	Group        string
	Stream       string
}

func NewCloudWatchLogsService(sess *session.Session) *cloudwatchlogs.CloudWatchLogs {
	return cloudwatchlogs.New(sess)
}

func NewAwsCloudWatchServiceImpl(groupName string, streamName string, sendToLog bool, logging *cloudwatchlogs.CloudWatchLogs, slackService slack.SlackService) CloudWatchService {
	return &CloudWatchServiceImpl{
		Logging:      logging,
		SendToLog:    sendToLog,
		SlackService: slackService,
		Group:        groupName,
		Stream:       streamName,
	}
}

func (service *CloudWatchServiceImpl) SendLog(flag string, a ...interface{}) bool {
	defer helpers.RecoverLoggerError()
	if !service.SendToLog {
		return true
	}
	message := fmt.Sprintf("[%s] ", flag)
	for _, item := range a {
		message += fmt.Sprint(item) + " "
	}
	_, err := service.Logging.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(service.Group),
		LogStreamName: aws.String(service.Stream),
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(message),
				Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		},
	})
	helpers.IfErrorHandler(err)
	return true
}

func (service *CloudWatchServiceImpl) Info(a ...interface{}) bool {
	log.Info(a)
	return service.SendLog("info", a...)
}
func (service *CloudWatchServiceImpl) Warning(a ...interface{}) bool {
	log.Warn(a)
	return service.SendLog("warning", a...)
}
func (service *CloudWatchServiceImpl) Error(a ...interface{}) bool {
	log.Error(a)
	if service.SendToLog {
		service.SlackService.SendMessage(a)
	}
	return service.SendLog("error", a...)
}
