package database

import (
	"log"

	"github.com/tofustream/gindo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// データベースの接続設定
func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gindo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	// モデルの自動マイグレーション（テーブル作成）
	DB.AutoMigrate(&models.Todo{})
}
