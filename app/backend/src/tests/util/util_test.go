package util

import (
	"log"
	"os"
	"testing"

	"go-rpg/util"
)

func TestCheckEnvValue(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		/* {
			name: "test name",
			want: "expect value",
		}, */
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.CheckEnvValue(); got != tt.want {
				t.Errorf("CheckEnvValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// 開始
	log.Print("util_test: start")
	// パッケージ内のテストの実行
	code := m.Run()
	// 終了
	log.Print("util_test: end")
	// テストの終了コードで exit
	os.Exit(code)
}
