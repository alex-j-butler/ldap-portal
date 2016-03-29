package logging

import (
	"github.com/ian-kent/go-log/log"
	"github.com/ian-kent/go-log/appenders"
	"github.com/ian-kent/go-log/logger"
	"github.com/ian-kent/go-log/layout"
	"qixalite.com/Ranndom/ldap-portal/modules/settings"
)

var (
	AppLogger logger.Logger
	HTTPLogger logger.Logger
)

func NewContext() {
	AppLayout := layout.Pattern("%d - [app %p] : %m")
	HTTPLayout := layout.Pattern("%d - [http %p] : %m")

	AppLogger = log.Logger("app")
	AppLogger.SetLevel(log.Stol(settings.Logging.AppLevel))

	HTTPLogger = log.Logger("http")
	HTTPLogger.SetLevel(log.Stol(settings.Logging.HTTPLevel))

	if settings.Logging.LogFile {
		AppAppender := appenders.RollingFile("logs/app.log", true)
		HTTPAppender := appenders.RollingFile("logs/http.log", true)

		AppAppender.SetLayout(AppLayout)
		HTTPAppender.SetLayout(HTTPLayout)

		AppLogger.SetAppender(AppAppender)
		HTTPLogger.SetAppender(HTTPAppender)
	} else {
		AppAppender := appenders.Console()
		HTTPAppender := appenders.Console()

		AppAppender.SetLayout(AppLayout)
		HTTPAppender.SetLayout(HTTPLayout)

		AppLogger.SetAppender(AppAppender)
		HTTPLogger.SetAppender(HTTPAppender)
	}
}
