package app

import "github.com/revel/revel"
import "gopkg.in/ldap.v2"

type UpdateAll struct {}

func (j UpdateAll) Run() {
    var users []User
    DB.Find(&users)

    for _,user := range users {
        source := LDAPSourceFromConfig(revel.Config)
        l, err := source.DialLDAP()
        if err != nil {
            revel.INFO.Printf("%s", err)
        }   
        source.BindLDAP(l)

        searchRequest := ldap.NewSearchRequest(
            user.DN,
            ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false,
            "(&)",
            []string{"dn", "uid", "givenName", "surname", "mail"},
            nil,
        )

        searchResponse, err := l.Search(searchRequest)
        if err != nil {
            revel.INFO.Printf("%s", err)
        }

        if len(searchResponse.Entries) == 1 {
            uid := searchResponse.Entries[0].GetAttributeValue("uid")
            givenName := searchResponse.Entries[0].GetAttributeValue("givenName")
            surname := searchResponse.Entries[0].GetAttributeValue("surname")
            mail := searchResponse.Entries[0].GetAttributeValue("mail")

            user.UID = uid
            user.GivenName = givenName
            user.Surname = surname
            user.Email = mail
            DB.Save(&user)
        }
    }
}

