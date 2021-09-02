package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * connect database.
 * @return {string}
 */
func Connection() string {
	fmt.Println("Connection:...... ")
	msg := ""

	// db, err := sql.Open("mysql", "userName:passWord@/dbName")
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 1
	var name string
	err = db.QueryRow("SELECT first_name FROM users WHERE id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	// ssssss := os.Getenv("TEST_VALUE")
	return msg
}
