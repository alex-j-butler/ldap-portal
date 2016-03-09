package controllers

import (
    "time"
    "fmt"
    "strconv"
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

func AccountSSHKeys(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    var user models.User
    var sshKeys []models.SSHKey
    database.DB.First(&user, &models.User{UID: sess.Get("LoggedUser").(string)})
    database.DB.Model(&user).Related(&sshKeys)

    ctx.Data["user"] = user
    ctx.Data["sshKeys"] = sshKeys
    ctx.Data["Title"] = "Account"
    ctx.Data["Subtitle"] = "SSH Keys"
    ctx.HTML(200, TMPL_ACCOUNT_SSH_KEYS)
}

func AccountNewSSHKey(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    ctx.Data["Title"] = "SSH Keys"
    ctx.Data["Subtitle"] = "New key"
    ctx.HTML(200, TMPL_ACCOUNT_NEW_SSH_KEY)
}

func AccountEditSSHKey(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    id64, _ := strconv.ParseInt(ctx.Params(":id"), 0, 64)
    id := int(id64)

    var sshKey models.SSHKey
    database.DB.First(&sshKey, &models.SSHKey{ID: id})

    ctx.Data["ssh_key"] = sshKey
    ctx.Data["Title"] = "SSH Keys"
    ctx.Data["Subtitle"] = fmt.Sprintf("Edit %s", sshKey.KeyName)
    ctx.HTML(200, TMPL_ACCOUNT_EDIT_SSH_KEY)
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

func POSTAccountSSHKeys(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    ctx.Redirect(ACCOUNT_SSH_KEYS)
}

func POSTAccountNewSSHKey(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    ctx.Redirect(ACCOUNT_SSH_KEYS)
}

func POSTAccountEditSSHKey(ctx *macaron.Context, f *session.Flash, sess session.Store, keyForm models.AccountSSHKeyForm) {
    id := ctx.ParamsInt(":id")

    var sshKey models.SSHKey
    database.DB.First(&sshKey, &models.SSHKey{ID: id})

    sshKey.KeyName = keyForm.KeyName
    sshKey.Key = keyForm.Key

    database.DB.Save(&sshKey)

    ctx.Redirect(ACCOUNT_EDIT_SSH_KEY(id))
}

func POSTAccountDeleteSSHKey(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    ctx.Redirect(ACCOUNT_SSH_KEYS)
}

func POSTAccountChangePassword(ctx *macaron.Context, f *session.Flash, sess session.Store, password models.AccountChangePasswordForm) {
    validation.Validate(password)

    ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
}

