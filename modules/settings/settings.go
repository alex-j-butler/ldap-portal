package settings

import (
    "gopkg.in/ini.v1"
)

var (
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
)

func NewContext() {
    cfg, err := ini.Load("conf/app.ini")
    if err != nil {
        
    }

    sec := cfg.Section("ldap")
    LDAP.Hostname = sec.Key("HOSTNAME").MustString("127.0.0.1")
    LDAP.Port = sec.Key("PORT").MustInt(636)
    LDAP.UseSSL = sec.Key("USE_SSL").MustBool(true)
    LDAP.SkipVerify = sec.Key("SKIP_VERIFY").MustBool(true)
    LDAP.BindDN = sec.Key("BIND_DN").MustString("cn=Directory Manager")
    LDAP.BindPass = sec.Key("BIND_PASS").MustString("")
    LDAP.UserSearch = sec.Key("USER_SEARCH").MustString("ou=People,dc=example,dc=com")
    LDAP.UserFilter = sec.Key("USER_FILTER").MustString("(&(objectClass=inetOrgPerson)(uid=%s))")
    LDAP.AdminFilter = sec.Key("ADMIN_FILTER").MustString("(&(objectClass=inetOrgPerson)(isMemberOf=cn=Admin,ou=Groups,dc=example,dc=com))")
}

