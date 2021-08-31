package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	response := &responseType{
		Id:      1,
		Name:    "test name",
		Detail:  "test detail",
		Address: "test address",
	}
	json.NewEncoder(writer).Encode(response)
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

	//サーバー起動
	if err := http.ListenAndServe(":8080", GetRouter()); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Printf("-------------------------------------\n")
}
