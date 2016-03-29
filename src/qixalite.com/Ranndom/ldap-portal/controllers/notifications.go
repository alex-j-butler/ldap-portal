package controllers

import (
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/session"
	"qixalite.com/Ranndom/ldap-portal/models"
	"qixalite.com/Ranndom/ldap-portal/modules/database"
)

func ListNotifications(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	notifications := models.UnreadNotifications(sess)

	ctx.Data["notifications"] = notifications
	ctx.Data["Title"] = "Notifications"
	ctx.Data["Subtitle"] = "Recent activity and changes to your account"
	ctx.HTML(200, TMPL_NOTIFICATIONS_LIST)
}

func ListNotificationHistory(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	notifications := models.ReadNotifications(sess)

	ctx.Data["notifications"] = notifications
	ctx.Data["Title"] = "Notifications"
	ctx.Data["Subtitle"] = "Historic activity and changes to your account"
	ctx.HTML(200, TMPL_NOTIFICATIONS_HISTORY)
}

func ViewNotification(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	id := ctx.ParamsInt(":id")

	var notification models.Notification

	// Get notification.
	database.DB.Where(&models.Notification{
		UID: sess.Get("LoggedUser").(string),
		ID: id,
	}).Find(&notification)

	ctx.Data["notification"] = notification
	ctx.Data["Title"] = "Notifications"
	ctx.Data["Subtitle"] = notification.Title
	ctx.HTML(200, TMPL_NOTIFICATIONS_VIEW)
}

func MarkNotification(ctx *macaron.Context, f *session.Flash, sess session.Store) {
	id := ctx.ParamsInt(":id")

	var notification models.Notification

	// Get notification.
	database.DB.Where(&models.Notification{
		UID: sess.Get("LoggedUser").(string),
		ID: id,
	}).First(&notification)

	notification.Read = true
	database.DB.Save(&notification)

	ctx.Redirect(NOTIFICATIONS_LIST)
}

