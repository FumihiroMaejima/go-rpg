package usersService

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rpg/repository/users"
)

/**
 * get users service.
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {[]byte}
 */
func GetUsers(w http.ResponseWriter, r *http.Request) []byte {
	fmt.Println("GetUsers Service: ")
	fmt.Println("Request: ", r)

	users := users.FetchUsers()

	//json形式に変換
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Fatal("Json Encoding Error.")
		log.Fatal(err)
	}
	return bytes
}
