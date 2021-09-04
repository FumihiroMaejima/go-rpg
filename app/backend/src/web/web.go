package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rpg/controllers/usersController"
	"go-rpg/repository/users"

	"github.com/gorilla/mux"
)

type responseType struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Detail  string `json:"detail"`
	Address string `json:"address"`
}

func Public(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieve Request: ")
	fmt.Println("Request: ", r)

	users := users.FetchUsers()

	/* response := &responseType{
		Id:      1,
		Name:    "test name",
		Detail:  "test detail",
		Address: "test address",
	}
	json.NewEncoder(w).Encode(response) */

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

/**
 * get router data.
 * @return {*mux.Router}
 */
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	// localhost:8080/publicでpublicハンドラーを実行
	// router.Handle("/public", public)
	// router.HandleFunc("/public", Public).Methods("GET")
	router.HandleFunc("/public", usersController.Get).Methods("GET")

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
