package helpers

import "github.com/gofiber/fiber/v2/log"

func RecoverLoggerError() {
	err := recover()
	if err != nil {
		log.Info(err)
	}
}
