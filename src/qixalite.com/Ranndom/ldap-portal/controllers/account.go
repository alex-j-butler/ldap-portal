package controllers

import (
    "time"
    "strings"
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "qixalite.com/Ranndom/ldap-portal/models"
    "qixalite.com/Ranndom/ldap-portal/modules/database"
    "qixalite.com/Ranndom/ldap-portal/modules/jobs"
    "qixalite.com/Ranndom/ldap-portal/modules/validation"
)

func AccountDetails(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    var user models.User
    database.DB.Where(&models.User{UID: sess.Get("LoggedUser").(string)}).First(&user)

    ctx.Data["user"] = user
    ctx.Data["Title"] = "Account"
    ctx.Data["Subtitle"] = "Details"
    ctx.HTML(200, TMPL_ACCOUNT_DETAILS)
}

func AccountChangePassword(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    ctx.Data["Title"] = "Account"
    ctx.Data["Subtitle"] = "Change password"
    ctx.HTML(200, TMPL_ACCOUNT_CHANGE_PASSWORD)
}

func POSTAccountDetails(ctx *macaron.Context, f *session.Flash, sess session.Store, account models.AccountDetailsForm) {
    var user models.User
    database.DB.Where(&models.User{UID: sess.Get("LoggedUser").(string)}).First(&user)

    user.GivenName = account.GivenName
    user.Surname = account.Surname

    database.DB.Save(&user)
    jobs.UpdateUserJob.Schedule(1, time.Now(), &jobs.UpdateUser{User: user})

    f.Success("Updated account!")
    ctx.Redirect(ACCOUNT_DETAILS)
}

func POSTAccountChangePassword(ctx *macaron.Context, f *session.Flash, sess session.Store, password models.AccountChangePasswordForm) {
    valid, errs := validation.Validate(password)

    if !valid {
        // Handle errors by pushing errors to flash.
        f.Error(strings.Join(errs, "//n"))
        ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
        return
    }

    ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
}

