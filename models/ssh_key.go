package models

type SSHKey struct {
    ID          int `sql:"AUTO_INCREMENT"`
    Key         string `sql:"type:text"`
    KeyName     string
    User        User
    UserID      int
}

