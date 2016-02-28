package controllers

import "github.com/revel/revel"
import "qixalite.com/Ranndom/ldap-portal/app"

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) Login() revel.Result {
    return c.Render()
}

func (c App) Logout() revel.Result {
    delete(c.Session, "user")
    c.Flash.Success("Successfully logged out!")

    return c.Redirect(App.Index)
}

func (c App) Account() revel.Result {
    return c.Render()
}

func (c App) POST_Account(givenName string, surname string) revel.Result {
    c.Validation.Required(givenName).Message("Please enter a first name")
    c.Validation.Required(surname).Message("Please enter a surname")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Account)
    }

    c.Flash.Success("Updated account!")

    return c.Redirect(App.Account)
}

func (c App) POST_Login(username string, password string) revel.Result {
    c.Validation.Required(username).Message("Please enter a username")
    c.Validation.Required(password).Message("Please enter a password")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Login)
    }

    u, status := app.GetLDAPUser(username)
    if status == false {
        // User could not be retrieved.
        c.Flash.Error("Invalid username/password")
        return c.Redirect(App.Login)
    }

    c.Session["user"] = u.UID
    c.Flash.Success("Welcome, " + u.UID)

    return c.Redirect(App.Index)
}

