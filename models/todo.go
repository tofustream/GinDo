package models

import (
	"gorm.io/gorm"
)

// ToDoモデルの定義
type Todo struct {
	gorm.Model        // ID、作成日、更新日などの標準フィールド
	Title      string `json:"title"`  // ToDoのタイトル
	Status     bool   `json:"status"` // 完了ステータス
}
