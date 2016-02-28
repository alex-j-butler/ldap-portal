package app

import "github.com/revel/revel"
import "gopkg.in/ldap.v2"
import "fmt"
import "crypto/tls"

type LDAPSource struct {
    Name                string
    Host                string
    Port                int
    UseSSL              bool
    SkipVerify          bool
    BindDN              string
    BindPassword        string
    UserSearch          string
    UserFilter          string
    AdminFilter         string
}

func LDAPSourceFromConfig(config *revel.MergedConfig) LDAPSource {
    source := LDAPSource{"Config source",
                         config.StringDefault("ldap.hostname", "localhost"),
                         config.IntDefault("ldap.port", 636),
                         config.BoolDefault("ldap.usessl", true),
                         config.BoolDefault("ldap.skipverify", true),
                         config.StringDefault("ldap.binddn", "cn=Directory Manager"),
                         config.StringDefault("ldap.bindpass", ""),
                         config.StringDefault("ldap.usersearch", "ou=People,dc=example,dc=com"),
                         config.StringDefault("ldap.userfilter", "(&(objectClass=inetOrgPerson)(uid=%s))"),
                         config.StringDefault("ldap.adminfilter", "(&(objectClass=inetOrgPerson)(isMemberOf=cn=Admin,ou=Groups,dc=example,dc=com))"),
    }
    return source
}

func (source LDAPSource) DialLDAP() (*ldap.Conn, error) {
    if source.UseSSL {
        if source.SkipVerify {
            revel.WARN.Printf("Connecting to LDAP %s (%s:%d) without TLS verification!", source.Name, source.Host, source.Port)
        }

        return ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", source.Host, source.Port), &tls.Config{InsecureSkipVerify: source.SkipVerify,})
    } else {
        return ldap.Dial("tcp", fmt.Sprintf("%s:%d", source.Host, source.Port))
    }
}

func (source LDAPSource) BindLDAP(l *ldap.Conn) {
    err := l.Bind(source.BindDN, source.BindPassword)
    if err != nil {
        revel.ERROR.Printf("%s", err)
    }
}

