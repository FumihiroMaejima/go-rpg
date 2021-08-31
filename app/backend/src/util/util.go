package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/**
 * check env value.
 * @return {string}
 */
func CheckEnvValue() string {
	msg := ""
	// .eenvの読み込み(pathはgo.modとgosumを基準にしている)s
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		// .env読めなかった場合の処理
		log.Output(0, "no environmental value. check server.\n")
		fmt.Println("Error: ", err)
		msg = "get env error"
		return msg
	}
	// ssssss := os.Getenv("TEST_VALUE")
	return msg
}
