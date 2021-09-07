package usersController

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rpg/services/usersService"
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

	// users := users.FetchUsers()
	// サービス側でロジックの実行
	// レスポンスをjsonで取得する
	response := usersService.GetUsers(w, r)
	log.Println("service end.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(response)))
	log.Println("controller end.")
}
