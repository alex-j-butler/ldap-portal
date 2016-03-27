package jobs

import (
    ldap_lib "gopkg.in/ldap.v2"

    "qixalite.com/Ranndom/ldap-portal/models"
    "qixalite.com/Ranndom/ldap-portal/modules/ldap"

    "log"
)

type UpdateUser struct {
    User    models.User
}

func RunUpdateUser(data *UpdateUser) error {
    source := ldap.LDAPSourceFromConfig()
    l, err := source.DialLDAP()
    if err != nil {
        SendLDAPFailNotification(data)

        log.Printf("%s", err)
        return err
    }
    source.BindLDAP(l)

    modifyRequest := ldap_lib.NewModifyRequest(data.User.DN)
    modifyRequest.Replace("givenName", []string{data.User.GivenName})
    modifyRequest.Replace("sn", []string{data.User.Surname})

    err = l.Modify(modifyRequest)
    if err != nil {
        // TODO: Create a notification for the user telling them
        //       their details failed to update.

        SendLDAPFailNotification(data)

        log.Printf("%s", err)
        return err
    }

    SendLDAPSuccessNotification(data)

    return nil
}

func SendLDAPFailNotification(data *UpdateUser) {
    models.NewNotification(
        &data.User,
        "Failed to synchronise account details",
        "Recent changes to your account failed to synchronise to the directory server. Changes have not been made to your account.",
        models.STATUS_DANGER,
    )
}

func SendLDAPSuccessNotification(data *UpdateUser) {
    models.NewNotification(
        &data.User,
        "Updated account details",
        "Your account details were successfully updated.",
        models.STATUS_INFO,
    )
}

