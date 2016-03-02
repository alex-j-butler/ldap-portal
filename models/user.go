package models

import (
    ldap_lib "gopkg.in/ldap.v2"
    "fmt"
    "qixalite.com/Ranndom/ldap-portal/modules/database"
    "qixalite.com/Ranndom/ldap-portal/modules/ldap"
)

type User struct {
    ID          int `sql:"AUTO_INCREMENT"`
    DN          string `gorm:"column:dn"`
    UID         string `gorm:"column:uid"`
    GivenName   string
    Surname     string
    Email       string
}

func (u User) VerifyPassword(password string) (bool) {
    source := ldap.LDAPSourceFromConfig()
    l, err := source.DialLDAP()
    
    if err != nil {
        return false
    }
    source.BindLDAP(l)

    err = l.Bind(u.DN, password)
    if err != nil {
        return false
    }

    return true
}

func (u User) ResetPassword(password string) (bool) {
    source := ldap.LDAPSourceFromConfig()
    l, err := source.DialLDAP()

    if err != nil {
        return false
    }
    source.BindLDAP(l)

    passwordModifyRequest := ldap_lib.NewPasswordModifyRequest(u.DN, "", password)
    _, err = l.PasswordModify(passwordModifyRequest)

    if err != nil {
        return false
    }

    return true
}

func GetLDAPUser(username string) (User, bool) {
    var user User
    query := database.DB.First(&user, &User{UID: username})

    if query.Error == nil {
        return user, true
    } else {
        // User currently does not exist in database,
        // attempt to load from LDAP.
        source := ldap.LDAPSourceFromConfig()
        l, err := source.DialLDAP()
        if err != nil {
            return User{}, false
        }

        source.BindLDAP(l)

        // Create a search request to retrieve
        // DN, UID, GivenName, Surname and Email
        // from LDAP.
        searchRequest := ldap_lib.NewSearchRequest(
            source.UserSearch,
            ldap_lib.ScopeWholeSubtree, ldap_lib.NeverDerefAliases, 0, 0, false,
            fmt.Sprintf(source.UserFilter, username),
            []string{"dn", "uid", "givenName", "surname", "mail"},
            nil,
        )

        searchResponse, err := l.Search(searchRequest)
        if err != nil {
            return User{}, false
        }

        if len(searchResponse.Entries) != 1 {
            return User{}, false
        }

        dn := searchResponse.Entries[0].DN
        uid := searchResponse.Entries[0].GetAttributeValue("uid")
        givenName := searchResponse.Entries[0].GetAttributeValue("givenName")
        surname := searchResponse.Entries[0].GetAttributeValue("surname")
        mail := searchResponse.Entries[0].GetAttributeValue("mail")

        user = User{DN: dn, UID: uid, GivenName: givenName, Surname: surname, Email: mail}
        database.DB.Create(&user)

        return user, true
    }
}

