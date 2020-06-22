package databases

import (
	"database/sql"
	"log"
)
var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:asdfgh@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
