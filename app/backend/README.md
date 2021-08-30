# go-basic

my go basic practice.

---

# 構成

| 名前 | バージョン |
| :--- | :---: |
| goenv | goenv 2.0.0beta11 |
| go | 1.16.6 |
| react | 17.0.0 |
| TypeScript | 4.1.3 |

---

# Check Env Data

```Shell-session
$ go env
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/userName/Library/Caches/go-build"
GOENV="/Users/userName/Library/Application Support/go/env"
GOMODCACHE="/Users/userName/go/1.16.4/pkg/mod"
GOPATH="/Users/userName/go/1.16.4"
GOROOT="/Users/userName/.goenv/versions/1.16.4"
...
```

---

# Install Develop Tool

```Shell-session
$ cmd + Shift + P(keyborad)
> Go: Install/Update Tools

Tools environment: GOPATH=/Users/userName/go/1.16.4
Installing 10 tools at /Users/userName/go/1.16.4/bin in module mode.
  gopkgs
  go-outline
  gotests
  gomodifytags
  impl
  goplay
  dlv
  dlv-dap
  staticcheck
  gopls

Installing github.com/uudashr/gopkgs/v2/cmd/gopkgs (/path/bin/gopkgs) SUCCEEDED
Installing github.com/ramya-rao-a/go-outline (/path/bin/go-outline) SUCCEEDED
Installing github.com/cweill/gotests/gotests (/path/bin/gotests) SUCCEEDED
Installing github.com/fatih/gomodifytags (/path/bin/gomodifytags) SUCCEEDED
Installing github.com/josharian/impl (/path/bin/impl) SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay (/path/bin/goplay) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv (/path/bin/dlv) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv@master (/path/bin/dlv-dap) SUCCEEDED
Installing honnef.co/go/tools/cmd/staticcheck (/path/bin/staticcheck) SUCCEEDED
Installing golang.org/x/tools/gopls (/path/bin/gopls) SUCCEEDED
```

インストールしたい拡張機能を選択後にインストールが開始される。

`GOPATH`内の`bin`ディレクトリにインストールされる。

---

# VScode Command


### goplsの再起動

```Shell-session
> Go: Restart Language Server
```

### extentionで利用しているコマンドのインストール/更新

```Shell-session
> Go:  Install/Update Tools
```

### go.modの作成

```Shell-session
> Go: initialize go.mod
```

### テストコードの作成

カーソル直下の関数のテストコードの自動生成

```Shell-session
> Go: Generate Unit Tests For Function
```

### interface stubの作成

特定のinterfaceを実装するためのメソッドの自動生成

```Shell-session
> Go: Generate Interface Stubs
```

### struct dataの作成

カーソル下のstructをfieldを初期値で穴埋めする形で自動生成

```Shell-session
> Go: Fill Struct
```

### struct tagの作成

カーソル下のstructにタグを付与

```Shell-session
> Go: Add Tags To Struct Fields
```

### struct tagの作成

カーソル下のテスト関数のみのテストの実行

```Shell-session
> Go: Test Function At Cursor
```

---


# execution go file

```Shell-session
$ go run main.go
```


---

# Install godotenv

golangで環境変数を使う為のパッケージをインストールする

```Shell-session
$ go install github.com/joho/godotenv@latest
```

　`main.go`にimportの設定を追記する

```golang
pckage main

import "github.com/joho/godotenv"
```

---

# jwt request

```Shell-session
$ curl localhost:8080/private -H "Authorization:Bearer token"
```

---
# Install Tool in Global

```Shell-session
$ go install <package>@<version>

ex:
$ go install golang.org/x/tools/gopls@latest
```

---

# Add Module
開発中にモジュールを追加する場合

```Shell-session
$ mkdir add
$ cd add
$ go mod init add
go: creating new go.mod: module add
$ ls
go.mod
$ touch add.go

# 1行で実行する場合
$ mkdir module && cd module/ && go mod init module && touch module.go

# main.goに実行後に下記のコマンドで追加する
$ go get package go-basic/module
# 上記で追加出来ない場合は下記のコマンド`go.mod`を更新する
$ go mod tidy
```

---

# Make Module
→不要かもしれない

```Shell-session
$ mkdir add
$ cd add
$ go mod init add
go: creating new go.mod: module add
$ ls
go.mod
$ touch add.go
```

---

# Setting Env for Example

```Shell-session
$ mkdir src
$ cd src
$ go mod init go-basic
go: creating new go.mod: module go-basic
go: to add module requirements and sums:
        go mod tidy

$ mkdir basic
$ cd basic
$ go mod init basic
go: creating new go.mod: module basic
$ ls
go.mod
```

## make & egit add.go

```Shell-session
$ touch basic.go
```

```golangs
package add

// use uppercase letter
func Add(a, b int) int {
    return a + b
}
```

## edit main.go

```shell-sesshion
$ cd ..
$ vim main.go
```

## edit gotest/go.mod

```shell-sesshion
$ vim go.mod
```

```golang
module gotest

go 1.16
replace go-test/add => ./add

require go-test/add v0.0.0-00010101000000-000000000000
```

## build main.go

```shell-sesshion
$ go build -o build/test-app main.go
```

## excute application

```shell-sesshion
$ ./build/test-app 14 4
18
```



---
