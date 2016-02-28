package app

import "github.com/revel/revel"
import "gopkg.in/ldap.v2"
import "fmt"

type User struct {
    ID              int `sql:"AUTO_INCREMENT"`
    UID             string
    GivenName       string
    Surname         string
    Email           string
}

func GetLDAPUser(username string) (User, bool) {
    var user User
    // query := DB.Where(&User{UID: username}).First(&user)
    query := DB.First(&user, &User{UID: username})

    if query.Error == nil {
        return user, true
    } else {
        // User currently does not exist in database,
        // attempt to load from LDAP.
        source := LDAPSourceFromConfig(revel.Config)
        l, err := source.DialLDAP()
        if err != nil {
            revel.ERROR.Fatal(err)
            return User{}, false
        }
        source.BindLDAP(l)

        // Create a search request to retrieve
        // DN, UID, GivenName, Surname and Email
        // from LDAP.
        searchRequest := ldap.NewSearchRequest(
            source.UserSearch,
            ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
            fmt.Sprintf(source.UserFilter, username),
            []string{"dn", "uid", "givenName", "surname", "mail"},
            nil,
        )

        searchResponse, err := l.Search(searchRequest)
        if err != nil {
            revel.ERROR.Fatal(err)
            return User{}, false
        }

        if len(searchResponse.Entries) != 1 {
            revel.ERROR.Printf("User %s does not exist or too many entries", username)
            return User{}, false
        }

        //dn := searchResponse.Entries[0].DN
        uid := searchResponse.Entries[0].GetAttributeValue("uid")
        givenName := searchResponse.Entries[0].GetAttributeValue("givenName")
        surname := searchResponse.Entries[0].GetAttributeValue("surname")
        mail := searchResponse.Entries[0].GetAttributeValue("mail")

        // Create the new user and save it to the persistent DB.
        user = User{UID: uid, GivenName: givenName, Surname: surname, Email: mail}
        DB.Create(&user)

        return user, true
    }
}

