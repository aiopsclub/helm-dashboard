package model

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbPath := path.Join(cwd, "db/system.db")
	Engine, err = xorm.NewEngine("sqlite3", dbPath)
	err = Engine.Sync2(new(User))
	if err != nil {
		fmt.Println(err)
		os.Exit(10005)
	}
	fmt.Println("sync models")

}
