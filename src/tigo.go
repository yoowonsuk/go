package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "go-ora"
	//_ "godror-master"
	//      "time"
)

func main() {
	// oracle
	conn, err := sql.Open("oracle", "oracle://system:tibero@192.168.6.85:1521/xe")
	// oracle
	//conn, err := sql.Open("godror", `user="system" password="tibero" connectString="192.168.6.85:1521/xe"`)

	//conn, err := sql.Open("oracle", "oracle://hello:tibero@172.23.4.101:31551/tibero")
	//conn, err := sql.Open("godror", `user="hello" password="tibero" connectString="172.23.4:wq.101:1521/xe" libDir="C:\Users\Tmax\Downloads\instantclient_19_10"`)

	//conn, err := sql.Open("godror", `user="system" password="tibero" connectString="172.23.4.101:31551/xe"`)
	if err != nil {
		fmt.Println("sql.open error")
	}
	defer conn.Close()

	fmt.Println("completed")

	//      rows, err := conn.Query("SELECT NAME, AGE FROM TEST")
	//      fmt.Println(rows)

	//conn.Exec("INSERT INTO test(name, age) VALUES ('Brother', 14)")

	var name string
	var age string
	err = conn.QueryRow("SELECT name, age FROM test").Scan(&name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	/*
	   var id int
	   var name string
	   rows, err := conn.Query("SELECT id, name FROM test1 where id >= ?", 1)
	   if err != nil {
	           log.Fatal(err)
	   }
	   defer rows.Close() //▒~X▒~S~\▒~K~\ ▒~K▒▒~J~T▒~K▒ (▒~@▒~W▒▒~U~X▒~W▒ ▒~K▒기)

	   for rows.Next() {
	           err := rows.Scan(&id, &name)
	           if err != nil {
	               log.Fatal(err)
	           }
	           fmt.Println(id, name)

	   }
	*/
}
