package controllers

import "github.com/revel/revel"
import "qixalite.com/Ranndom/ldap-portal/app"

type Auth struct {
    *revel.Controller
}

func (c Auth) Login() revel.Result {
    return c.Render()
}

func (c Auth) Logout() revel.Result {
    // Delete the session.
    delete(c.Session, "user")
    c.Flash.Success("Successfully logged out!")

    return c.Redirect(App.Index)
}

func (c Auth) POST_Login(username string, password string) revel.Result {
    c.Validation.Required(username).Message("Please enter a username")
    c.Validation.Required(password).Message("Please enter a password")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(Auth.Login)
    }

    u, status := app.GetLDAPUser(username)
    passwordStatus := u.VerifyPassword(password)
    if status == false || passwordStatus == false {
        // User could not be retrieved.
        c.Flash.Error("Invalid username/password")
        return c.Redirect(Auth.Login)
    }

    c.Session["user"] = u.UID
    c.Flash.Success("Welcome, " + u.UID)

    return c.Redirect(App.Index)
}

