package main

import (
	"github.com/tofustream/gindo/database"
	"github.com/tofustream/gindo/routes"
)

func main() {
	// データベース接続を初期化
	database.Connect()

	// ルーターを設定しサーバーを起動
	r := routes.SetupRouter()
	r.Run(":8080")
}
