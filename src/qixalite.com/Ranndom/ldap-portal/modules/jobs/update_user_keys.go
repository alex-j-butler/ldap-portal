package jobs

import (
    ldap_lib "gopkg.in/ldap.v2"
    
    "qixalite.com/Ranndom/ldap-portal/models"
    "qixalite.com/Ranndom/ldap-portal/modules/ldap"
    "qixalite.com/Ranndom/ldap-portal/modules/database"

    "log"
)

type UpdateUserKeys struct {
    User    models.User
}

func UpdateUserKeysRun(data *UpdateUserKeys) error {
    source := ldap.LDAPSourceFromConfig()
    l, err := source.DialLDAP()
    if err != nil {
        log.Printf("%s", err)
        return err
    }
    source.BindLDAP(l)

    // Retrieve keys.
    var sshKeys []models.SSHKey
    database.DB.Model(&data.User).Related(&sshKeys)
    stringSSHKeys := models.KeysToString(sshKeys)

    modifyRequest := ldap_lib.NewModifyRequest(data.User.DN)
    modifyRequest.Replace("sshPublicKey", stringSSHKeys)
    
    err = l.Modify(modifyRequest)
    if err != nil {
        log.Printf("%s", err)
        return err
    }

    return nil
}

