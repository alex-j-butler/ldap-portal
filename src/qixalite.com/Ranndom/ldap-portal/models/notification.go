package models

import (
	"time"
	"github.com/go-macaron/session"
	"qixalite.com/Ranndom/ldap-portal/modules/database"
)

type Notification struct {
	ID        int     `sql:"AUTO_INCREMENT"`
	UID       string  `gorm:"column:uid"`
	Title     string
	Message   string  `gorm:"type:TEXT"`
	Read      bool
	Status    int
	CreatedAt time.Time
}

const (
	STATUS_INFO = iota
	STATUS_WARNING
	STATUS_DANGER
)

func (n Notification) StatusIconName() string {
	switch n.Status {
	case STATUS_INFO:
		return "info"
	case STATUS_WARNING:
		return "exclamation-triangle"
	case STATUS_DANGER:
		return "exclamation"
	default:
		return "question"
	}
}

func NewNotification(user *User, title string, message string, status int) Notification {
	notification := Notification{
		UID: user.UID,
		Title: title,
		Message: message,
		Read: false,
		Status: status,
	}

	database.DB.Create(&notification)

	return notification
}

func UnreadNotifications(sess session.Store) []Notification {
	var notifications []Notification

	database.DB.Where("`uid` = ? AND `read` = ?",
		sess.Get("LoggedUser").(string),
		false,
	).Find(&notifications)

	return notifications
}

func ReadNotifications(sess session.Store) []Notification {
	var notifications []Notification

	database.DB.Where("`uid` = ? AND `read` = ?",
		sess.Get("LoggedUser").(string),
		true,
	).Find(&notifications)

	return notifications
}

