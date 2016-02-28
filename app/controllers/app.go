package controllers

import "github.com/revel/revel"
import "github.com/revel/modules/jobs/app/jobs"
import "qixalite.com/Ranndom/ldap-portal/app"

type App struct {
    *revel.Controller
}

func LoggedIn(c App) bool {
    if c.Session["user"] != "" {
        // User is set, confirm user exists.
        var user app.User
        query := app.DB.Where(&app.User{UID: c.Session["user"]}).First(&user)

        if query.Error == nil {
            return true
        }
    }

    return false
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) Account() revel.Result {
    loggedIn := LoggedIn(c)
    if loggedIn != true {
        c.Flash.Error("You must be logged in to access that!")
        return c.Redirect(App.Index)
    }

    var user app.User
    app.DB.Where(&app.User{UID: c.Session["user"]}).First(&user)

    return c.Render(user)
}

func (c App) ChangePassword() revel.Result {
    loggedIn := LoggedIn(c)
    if loggedIn != true {
        c.Flash.Error("You must be logged in to access that!")
        return c.Redirect(App.Index)
    }

    var user app.User
    app.DB.Where(&app.User{UID: c.Session["user"]}).First(&user)

    return c.Render(user)
}

func (c App) POST_Account(givenName string, surname string) revel.Result {
    loggedIn := LoggedIn(c)
    if loggedIn != true {
        c.Flash.Error("You must be logged in to access that!")
        return c.Redirect(App.Index)
    }

    var user app.User
    app.DB.First(&user, &app.User{UID: c.Session["name"]})

    c.Validation.Required(givenName).Message("Please enter a first name")
    c.Validation.Required(surname).Message("Please enter a surname")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Account)
    }

    user.GivenName = givenName
    user.Surname = surname

    app.DB.Save(&user)
    jobs.Now(app.UpdateUser{user})

    c.Flash.Success("Updated account!")

    return c.Redirect(App.Account)
}

func (c App) POST_ChangePassword(currentPassword string, newPassword string, confirmPassword string) revel.Result {
    loggedIn := LoggedIn(c)
    if loggedIn != true {
        c.Flash.Error("You must be logged in to access that!")
        return c.Redirect(App.Index)
    }

    var user app.User
    app.DB.First(&user, &app.User{UID: c.Session["name"]})

    c.Validation.Required(currentPassword).Message("Please enter your current password")
    c.Validation.Required(newPassword).Message("Please enter your new password")
    c.Validation.Required(confirmPassword).Message("Please confirm your new password")
    c.Validation.Required(newPassword == confirmPassword).Message("Password confirmation does not match")

    c.Validation.MinSize(newPassword, 8).Message("Password must be at least 8 characters")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.ChangePassword)
    }

    // Check if the current password is valid.
    passwordStatus := user.VerifyPassword(currentPassword)
    if passwordStatus != true {
        c.Flash.Error("Current password is incorrect")
        return c.Redirect(App.ChangePassword)
    }

    // Update the password
    updateStatus := user.ResetPassword(newPassword)
    if updateStatus != true {
        c.Flash.Error("Failed to reset password")
        return c.Redirect(App.ChangePassword)
    }

    c.Flash.Success("Updated password!")

    return c.Redirect(App.Account)
}

