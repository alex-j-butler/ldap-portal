package logging

import (
	"github.com/op/go-logging"
	"os"
	"qixalite.com/Ranndom/ldap-portal/modules/settings"
	"fmt"

	"github.com/Ranndom/PapertrailBackend"
)

var Logger *logging.Logger = logging.MustGetLogger("main")

var format = logging.MustStringFormatter(
	`%{time:15:04:05} %{color} [%{level:.4s}] %{shortfile} %{shortfunc}() : %{message} %{color:reset}`,
)

func NewContext() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)

	if settings.Logging.Address != "" {
		format = logging.MustStringFormatter(`[%{level:.4s}] %{shortfile} %{shortfunc}() : %{message}`)

		papertrailBackend, err := RemoteSyslog.NewPapertrailBackend(&RemoteSyslog.PapertrailBackend{
			Hostname: settings.Logging.Address,
			Port: settings.Logging.Port,
			Network: "udp",
			Tag: "ldap_portal",
		})
		if err != nil {
			fmt.Printf("Papertrail error: %s\n", err)
		}

		syslogBackend := logging.NewBackendFormatter(papertrailBackend, format)
		logging.SetBackend(backendFormatter, syslogBackend)
	}
}
