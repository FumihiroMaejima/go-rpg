package usersController

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rpg/repository/users"
)

/**
 * get request handling.
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {void}
 */
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieve Request: ")
	fmt.Println("Request: ", r)

	users := users.FetchUsers()

	//json形式に変換
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Fatal("Json Encoding Error.")
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(bytes)))
}
