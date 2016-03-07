package models

import (
    "regexp"
)

type SSHKey struct {
    ID          int `sql:"AUTO_INCREMENT"`
    Key         string `sql:"type:text"`
    KeyName     string
    User        User
    UserID      int
}

func (key SSHKey) ToString() (string) {
    return key.Key
}

func GenerateKeyName(key string) (string) {
    re, _ := regexp.Compile("ssh-rsa AAAA[0-9A-Za-z+/]+[=]{0,3} ([^@]+@[^@]+)")
    res := re.FindAllStringSubmatch(key, -1)

    return res[0][1]
}

func KeysToString(keys []SSHKey) ([]string) {
    k := make([]string, len(keys))
    for x, key := range keys {
        k[x] = key.ToString()
    }

    return k
}

