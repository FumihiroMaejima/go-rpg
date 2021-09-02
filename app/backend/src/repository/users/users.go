package users

import (
	"fmt"
	"log"

	"go-rpg/db"
)

// jsonにエンコードする為jsonのタグをつける
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

/**
 * fetch users.
 * @return {[]User}
 */
func FetchUsers() []User {
	db := db.Connection()
	defer db.Close()

	//rowを取得
	rows, err := db.Query("SELECT id, name FROM users")
	// aワイルドカードまたはすべてのkey指定だとrows.Scan()でエラー
	// rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("SQL Query Error.")
		panic(err.Error())
	}

	//User型のスライスに格納
	users := make([]User, 0)

	fmt.Println("rows: ", rows)

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal("Rows Scan Error.")
			panic(err.Error())
		}
		users = append(users, user)
	}
	return users
}
