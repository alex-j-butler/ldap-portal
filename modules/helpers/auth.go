package helpers

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "qixalite.com/Ranndom/ldap-portal/modules/database"
    "qixalite.com/Ranndom/ldap-portal/models"
)

func LoggedIn(ctx *macaron.Context, sess session.Store) bool {
    userSess := sess.Get("LoggedUser")
    if userSess != nil {
        // User is set, confirm user exists.
        var user models.User
        query := database.DB.Where(&models.User{UID: userSess.(string)}).First(&user)

        if query.Error == nil {
            return true
        }
    }

    return false
}

