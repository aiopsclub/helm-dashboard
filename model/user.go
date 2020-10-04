package model

import (
	"time"
)

type User struct {
	Id      int64
	Name    string    `xorm:"varchar(25) notnull unique"`
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
