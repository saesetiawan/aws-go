package slack

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/saesetiawan/aws-go/pkg/helpers"
	"time"
)

type SlackService interface {
	SendMessage(a ...interface{})
}

type SlackServiceImpl struct {
	WebhookUrl string
	ENV        string
	SendLog    bool
}

func NewSlackService(webHookUrl string, environment string, sendLog bool) SlackService {
	return &SlackServiceImpl{
		WebhookUrl: webHookUrl,
		ENV:        environment,
		SendLog:    sendLog,
	}
}

func (service *SlackServiceImpl) SendMessage(a ...interface{}) {
	defer helpers.RecoverLoggerError()
	if len(service.WebhookUrl) == 0 || !service.SendLog {
		return
	}
	message := ""
	for _, item := range a {
		message += fmt.Sprint(item) + " "
	}
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	payload := slack.Payload{
		Text: fmt.Sprintf(":sos: [%s][%s] %s", timeNow, service.ENV, message),
	}
	err := slack.Send(service.WebhookUrl, "", payload)
	if len(err) > 0 {
		helpers.IfErrorHandler(err[0])
	}
}
