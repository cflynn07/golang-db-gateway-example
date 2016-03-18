package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	connString := os.ExpandEnv("root:${MYSQL_ROOT_PASSWORD}@tcp(mysql:3306)/${MYSQL_DATABASE}")

	db, err := sql.Open(
		"mysql",
		connString,
	)

	fmt.Printf(connString + "\n")
	checkErr(err)

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS `userinfo` (`uid` INT(10))")
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)

	fmt.Printf("Success!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
