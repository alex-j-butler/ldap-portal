package app

import "github.com/revel/revel"
import "gopkg.in/ldap.v2"

type UpdateUser struct {
    User    User
}

func (j UpdateUser) Run() {
    source := LDAPSourceFromConfig(revel.Config)
    l, err := source.DialLDAP()
    if err != nil {
        revel.ERROR.Fatal(err)
    }
    source.BindLDAP(l)

    modifyRequest := ldap.NewModifyRequest(j.User.DN)
    modifyRequest.Replace("givenName", []string{j.User.GivenName})
    modifyRequest.Replace("sn", []string{j.User.Surname})

    err = l.Modify(modifyRequest)
    if err != nil {
        revel.ERROR.Fatal(err)
    }
}

