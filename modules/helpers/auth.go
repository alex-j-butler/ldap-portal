package helpers

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "qixalite.com/Ranndom/ldap-portal/modules/database"
    "qixalite.com/Ranndom/ldap-portal/models"
)

func IsLoggedIn(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    userSess := sess.Get("LoggedUser")
    if userSess != nil {
        // User is set, confirm user exists.
        var user models.User
        query := database.DB.Where(&models.User{UID: userSess.(string)}).First(&user)

        if query.Error == nil {
            return
        }
    }

    f.Error("You must be logged in to access that!")
    ctx.Redirect("/")
}

