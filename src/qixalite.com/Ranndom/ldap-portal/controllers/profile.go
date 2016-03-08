package controllers

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "qixalite.com/Ranndom/ldap-portal/models"
    _ "qixalite.com/Ranndom/ldap-portal/modules/database"
    "fmt"
)

var (
    USER_PROFILE = func (name string) string {
	return fmt.Sprintf("/profile/%s", name)
    }
)

const (
    TMPL_USER_PROFILE = "profile/user_profile"
)

func UserProfile(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    user, _ := models.GetLDAPUser(ctx.Params(":name"))

    ctx.Data["user"] = user
    ctx.Data["title"] = fmt.Sprintf("%s", user.UID)
    ctx.HTML(200, TMPL_USER_PROFILE)
}

