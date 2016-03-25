package settings

import (
    "log"

    "gopkg.in/ini.v1"
)

var (
    // Web Information
    Web struct {
        Address     string
        Port        int
    }

    // Logging Information
    Logging struct {
        LogFile     bool
        GlobalLevel string
        AppLevel    string
        HTTPLevel   string
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
        Driver          string
        Spec            string
        MaxIdleConns    int
        MaxOpenConns    int
        LogMode         bool
    }

    // Session information
    Session struct {
        Provider        string
        ProviderConfig  string
        CookieName      string
        Secure          bool
        IDLength        int
    }

    // CSRF information
    CSRF struct {
        Secret  string
    }
)

func NewContext() {
    cfg, err := ini.Load("conf/app.ini")
    if err != nil {
        log.Fatal("Failed to parse 'conf/app.ini': %v", err)
    }

    sec := cfg.Section("web")
    Web.Address = sec.Key("ADDRESS").MustString("0.0.0.0")
    Web.Port = sec.Key("PORT").MustInt(4000)

    sec = cfg.Section("logging")
    Logging.LogFile = sec.Key("LOG_TO_FILE").MustBool(false)
    Logging.GlobalLevel = sec.Key("GLOBAL_LEVEL").MustString("INFO")
    Logging.AppLevel = sec.Key("APP_LEVEL").MustString(Logging.GlobalLevel)
    Logging.HTTPLevel = sec.Key("HTTP_LEVEL").MustString(Logging.GlobalLevel)

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

