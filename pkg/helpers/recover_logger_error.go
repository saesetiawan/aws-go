package helpers

import "github.com/go-playground/log/v8"

func RecoverLoggerError() {
	err := recover()
	if err != nil {
		log.Info(err)
	}
}
