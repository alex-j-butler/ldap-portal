package controllers

import (
	"strings"
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/session"
	"qixalite.com/Ranndom/ldap-portal/models"
	"qixalite.com/Ranndom/ldap-portal/modules/validation"
	_ "qixalite.com/Ranndom/ldap-portal/modules/helpers"
)

const (
	AUTH_LOGIN = "/auth/login"
	AUTH_LOGOUT = "/auth/logout"
)

const (
	TMPL_AUTH_LOGIN = "auth/login"
)

func AuthLogin(ctx *macaron.Context) {
	ctx.Data["Title"] = "Login"
	ctx.HTML(200, TMPL_AUTH_LOGIN)
}

func AuthLogout(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	// Delete the session.
	sess.Delete("LoggedUser")
	f.Success("Successfully logged out!")

	ctx.Redirect(HOME)
}

func POSTAuthLogin(ctx *macaron.Context, f *session.Flash, sess session.Store, login models.LoginForm) {
	valid, errs := validation.Validate(login)

	if !valid {
		// Handle errors by pushing errors to flash.
		f.Error(strings.Join(errs, "//n"))
		ctx.Redirect(AUTH_LOGIN)
		return
	}

	u, status := models.GetLDAPUser(login.Username)
	passwordStatus := u.VerifyPassword(login.Password)
	if status == false || passwordStatus == false {
		// User could not be retrieved.
		f.Error("Invalid username/password")
		ctx.Redirect(AUTH_LOGIN)
		return
	}

	sess.Set("LoggedUser", u.UID)
	f.Success("Welcome, " + u.UID)

	ctx.Redirect(HOME)
}

