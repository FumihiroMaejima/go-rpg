package db

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"

	"go-rpg/db"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	tests := []struct {
		name string
		want *sql.DB
	}{
		// TODO: Add test cases.
		/* {
			name: "test name",
			want: "expect value",
		}, */
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := db.Connection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// 開始
	log.Print("db_test: start")
	// パッケージ内のテストの実行
	code := m.Run()
	// 終了
	log.Print("db_test: end")
	// テストの終了コードで exit
	os.Exit(code)
}
