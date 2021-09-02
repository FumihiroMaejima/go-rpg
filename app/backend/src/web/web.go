package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rpg/repository/users"

	"github.com/gorilla/mux"
)

type responseType struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Detail  string `json:"detail"`
	Address string `json:"address"`
}

var public = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Recieve Request: ")
	fmt.Println("Request: ", request)

	users := users.FetchUsers()

	/* response := &responseType{
		Id:      1,
		Name:    "test name",
		Detail:  "test detail",
		Address: "test address",
	}
	json.NewEncoder(writer).Encode(response) */

	//json形式に変換
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Fatal("Json Encoding Error.")
		log.Fatal(err)
	}
	writer.Write([]byte(string(bytes)))
})

/**
 * get router data.
 * @return {*mux.Router}
 */
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	// localhost:8080/publicでpublicハンドラーを実行
	router.Handle("/public", public)

	return router
}

/**
 * execution module.
 * @return {void}
 */
func App() {
	fmt.Println("Execute Application Server:...... ")

	// db.Connection()

	//サーバー起動
	if err := http.ListenAndServe(":8080", GetRouter()); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Printf("-------------------------------------\n")
}
