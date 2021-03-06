package settings

import (
	"gopkg.in/ini.v1"
	"qixalite.com/Ranndom/ldap-portal/bindata"
	"fmt"
)

var (
// General Information
	General struct {
		UserAllowed []string
	}

// Web Information
	Web struct {
		Address string
		Port    int
	}

// Logging Information
	Logging struct {
		Address		string
		Port		int
	}

// LDAP Information
	LDAP struct {
		Hostname    string
		Port        int
		UseSSL      bool
		SkipVerify  bool
		BindDN      string
		BindPass    string
		UserSearch  string
		UserFilter  string
		AdminFilter string
	}

// Database information
	Database struct {
		Driver       string
		Spec         string
		MaxIdleConns int
		MaxOpenConns int
		LogMode      bool
	}

// Session information
	Session struct {
		Provider       string
		ProviderConfig string
		CookieName     string
		Secure         bool
		IDLength       int
	}

// CSRF information
	CSRF struct {
		Secret string
	}
)

func NewContext() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		cfg, err = ini.Load(bindata.MustAsset("conf/app-default.ini"))
		if err != nil {
			fmt.Errorf("Failed to parse 'conf/app.ini' and 'conf/app-default.ini': %v", err)
		}
	}

	sec := cfg.Section("general")
	General.UserAllowed = sec.Key("UserAllowed").Strings(",")

	sec = cfg.Section("web")
	Web.Address = sec.Key("ADDRESS").MustString("0.0.0.0")
	Web.Port = sec.Key("PORT").MustInt(4000)

	sec = cfg.Section("logging")
	Logging.Address = sec.Key("Address").MustString("")
	Logging.Port = sec.Key("Port").MustInt(0)

	sec = cfg.Section("ldap")
	LDAP.Hostname = sec.Key("HOSTNAME").MustString("127.0.0.1")
	LDAP.Port = sec.Key("PORT").MustInt(636)
	LDAP.UseSSL = sec.Key("USE_SSL").MustBool(true)
	LDAP.SkipVerify = sec.Key("SKIP_VERIFY").MustBool(true)
	LDAP.BindDN = sec.Key("BIND_DN").MustString("cn=Directory Manager")
	LDAP.BindPass = sec.Key("BIND_PASS").MustString("")
	LDAP.UserSearch = sec.Key("USER_SEARCH").MustString("ou=People,dc=example,dc=com")
	LDAP.UserFilter = sec.Key("USER_FILTER").MustString("(&(objectClass=inetOrgPerson)(uid=%s))")
	LDAP.AdminFilter = sec.Key("ADMIN_FILTER").MustString("(&(objectClass=inetOrgPerson)(isMemberOf=cn=Admin,ou=Groups,dc=example,dc=com))")

	sec = cfg.Section("database")
	Database.Driver = sec.Key("DRIVER").MustString("mysql")
	Database.Spec = sec.Key("SPEC").MustString("root:password1@/ldap_portal?charset=utf8")
	Database.MaxIdleConns = sec.Key("MAX_IDLE_CONNS").MustInt(10)
	Database.MaxOpenConns = sec.Key("MAX_OPEN_CONNS").MustInt(100)
	Database.LogMode = sec.Key("LOG_MODE").MustBool(true)

	sec = cfg.Section("session")
	Session.Provider = sec.Key("PROVIDER").MustString("memory")
	Session.ProviderConfig = sec.Key("PROVIDER_CONFIG").MustString("")
	Session.CookieName = sec.Key("COOKIE_NAME").MustString("QixaliteSession")
	Session.Secure = sec.Key("SECURE").MustBool(false)
	Session.IDLength = sec.Key("ID_LENGTH").MustInt(32)

	sec = cfg.Section("csrf")
	CSRF.Secret = sec.Key("SECRET").MustString("change-me")
}

