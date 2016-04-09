package ldap

import (
	ldap_lib "gopkg.in/ldap.v2"
	"fmt"
	"crypto/tls"

	"qixalite.com/Ranndom/ldap-portal/modules/settings"
	"qixalite.com/Ranndom/ldap-portal/modules/logging"
)

type LDAPSource struct {
	Name         string
	Host         string
	Port         int
	UseSSL       bool
	SkipVerify   bool
	BindDN       string
	BindPassword string
	UserSearch   string
	UserFilter   string
	AdminFilter  string
}

func (source LDAPSource) DialLDAP() (*ldap_lib.Conn, error) {
	if source.UseSSL {
		if source.SkipVerify {
			// Warn.
			logging.Logger.Warningf("Connecting to LDAP %s (%s:%d) without TLS verification!", source.Name, source.Host, source.Port)
		}

		return ldap_lib.DialTLS("tcp", fmt.Sprintf("%s:%d", source.Host, source.Port), &tls.Config{InsecureSkipVerify: source.SkipVerify, })
	} else {
		return ldap_lib.Dial("tcp", fmt.Sprintf("%s:%d", source.Host, source.Port))
	}
}

func (source LDAPSource) BindLDAP(l *ldap_lib.Conn) {
	// Bind LDAP using the BindDN & BindPassword.
	l.Bind(source.BindDN, source.BindPassword)
}

func LDAPSourceFromConfig() LDAPSource {
	source := LDAPSource{
		"Config source",
		settings.LDAP.Hostname,
		settings.LDAP.Port,
		settings.LDAP.UseSSL,
		settings.LDAP.SkipVerify,
		settings.LDAP.BindDN,
		settings.LDAP.BindPass,
		settings.LDAP.UserSearch,
		settings.LDAP.UserFilter,
		settings.LDAP.AdminFilter,
	}

	return source
}

