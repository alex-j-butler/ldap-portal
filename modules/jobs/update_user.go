package jobs

import (
    ldap_lib "gopkg.in/ldap.v2"

    "qixalite.com/Ranndom/ldap-portal/models"
    "qixalite.com/Ranndom/ldap-portal/modules/ldap"
)

type UpdateUser struct {
    User    models.User
}

func UpdateUserRun(data *UpdateUser) error {
    source := ldap.LDAPSourceFromConfig()
    l, err := source.DialLDAP()
    if err != nil {
        return err
    }
    source.BindLDAP(l)

    modifyRequest := ldap_lib.NewModifyRequest(data.User.DN)
    modifyRequest.Replace("givenName", []string{data.User.GivenName})
    modifyRequest.Replace("sn", []string{data.User.Surname})

    err = l.Modify(modifyRequest)
    if err != nil {
        return err
    }

    return nil
}

