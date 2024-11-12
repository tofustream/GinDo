# GinDo

GinDoは、GoとGinフレームワークを用いて開発されたシンプルなToDo管理APIです。

このAPIを利用することで、ToDo項目の作成、取得、更新、削除といった基本的な操作を簡単に行うことができます。

## セットアップ手順

1. リポジトリをクローンします。

```bash
git clone https://github.com/tofustream/gindo.git
cd gindo
```
2. 依存関係をインストールします。

```bash
go mod download
```

3. サーバーを起動します。

```bash
go run main.go
```

これでAPIサーバーが起動し、デフォルトでは http://localhost:8080 でアクセスできます。

## エンドポイントの使い方

- ToDoリストの取得: `GET /todos`
→ 新しいToDo項目を作成します。リクエストボディ例：
```json
{
  "title": "New ToDo",
  "status": false
}
```

- ToDoの更新: `PUT /todos/:id`
→ 指定IDのToDo項目を更新します。リクエストボディ例：
```json
{
  "title": "Updated ToDo",
  "status": true
}
```

- ToDoの削除: `DELETE /todos/:id`
→ 指定IDのToDo項目を削除します。