package controllers

import (
	"strings"
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/session"
	"github.com/fatih/structs"
	"qixalite.com/Ranndom/ldap-portal/models"
	"qixalite.com/Ranndom/ldap-portal/modules/database"
	"qixalite.com/Ranndom/ldap-portal/modules/jobs"
	"qixalite.com/Ranndom/ldap-portal/modules/validation"
	"qixalite.com/Ranndom/ldap-portal/modules/helpers"
	"qixalite.com/Ranndom/ldap-portal/modules/settings"
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

func UpdateAccountDetails(ctx *macaron.Context, f *session.Flash, sess session.Store, account models.AccountDetailsForm) {
	var user models.User
	database.DB.Where(&models.User{UID: sess.Get("LoggedUser").(string)}).First(&user)

	accountDetails := structs.Map(account)
	helpers.MapWhitelist(&accountDetails, settings.General.UserAllowed)
	helpers.UpdateStruct(user, &accountDetails)

	database.DB.Save(&user)

	go jobs.RunUpdateUser(&jobs.UpdateUser{User: user})

	f.Info("Updating your account, please wait a minute.")
	ctx.Redirect(ACCOUNT_DETAILS)
}

func UpdateAccountPassword(ctx *macaron.Context, f *session.Flash, sess session.Store, password models.AccountChangePasswordForm) {
	valid, errs := validation.Validate(password)

	if !valid {
		// Handle errors by pushing errors to flash.
		f.Error(strings.Join(errs, "//n"))
		ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
		return
	}

	if password.NewPassword != password.ConfirmPassword {
		f.Error("New passwords do not match")
		ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
		return
	}

	var user models.User
	database.DB.Where(&models.User{UID: sess.Get("LoggedUser").(string)}).First(&user)

	if user.VerifyPassword(password.CurrentPassword) == false {
		// Failed to verify their old password.
		f.Error("Invalid current password")
		ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
		return
	}

	if user.ResetPassword(password.NewPassword) == false {
		// Failed to change password for unknown reasons.
		f.Error("Failed to change password, please try again later")
		ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
		return
	}

	f.Info("Updating your password, please wait a minute.")
	ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
}

