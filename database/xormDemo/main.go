package main

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
	engine, err := xorm.NewEngine("mysql", "root:lkg123456@127.0.0.1:3306/study?charset=utf8")
	if err != nil {
		panic(err)
	}
	err = engine.Ping()
	if err != nil {
		panic(err)
	}

}
