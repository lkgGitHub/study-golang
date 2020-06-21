package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql","root:123456@tcp(localhost:3306)/study")
	if err != nil {
		log.Println(err)
	}
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		log.Println(err)
	}
	for rows.Next(){
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age);err != nil{log.Fatal(err)}
		fmt.Println(name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}


	defer db.Close()
}


