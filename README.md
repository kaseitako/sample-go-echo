# Sample Go Echo API

GoとEcho frameworkを使った学習用のサンプルWebAPIです。SQLiteデータベースを使用したユーザー管理のCRUD APIとOpenAPI（Swagger）ドキュメントの自動生成を学ぶことができます。

## 機能

- Echo v4を使用したWebサーバー
- SQLiteデータベースとの統合
- Squirrel SQLビルダーを使用したクエリ構築
- ユーザー管理のCRUD API（作成・読み取り・更新・削除）
- `/hello` エンドポイントでヘルスチェック
- Bearer Token認証（`/protected`エンドポイント用）
- OpenAPI/Swagger自動ドキュメント生成
- CORS対応
- リクエストログ出力
- エラーハンドリング

## 必要な環境

- Go 1.19以上
- swagコマンドラインツール

## セットアップ

1. **リポジトリのクローン**
   ```bash
   git clone <repository-url>
   cd sample-go-echo
   ```

2. **依存関係のインストール**
   ```bash
   go mod download
   ```

3. **swagツールのインストール**
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

4. **Swaggerドキュメントの生成**
   ```bash
   swag init
   ```

## 実行方法

### 開発環境での実行

```bash
go run main.go
```

### ビルドして実行

```bash
# ビルド（SQLite使用のためCGOが必要）
CGO_ENABLED=1 go build -o app

# 実行
./app
```

サーバーが起動すると、以下のURLでアクセスできます：

- **API**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger/index.html

### データベース

アプリケーション起動時に自動的に：
- SQLiteデータベースファイル（`sample.db`）が作成されます
- usersテーブルが作成されます（user_id, name, created_at）

## API エンドポイント

### GET /hello

ヘルスチェックエンドポイント（認証不要）

**レスポンス:**
```json
{
  "status": "OK"
}
```

**curlでのテスト:**
```bash
curl http://localhost:8080/hello
```

### ユーザー管理API（認証不要）

#### POST /users - ユーザー作成

**リクエストボディ:**
```json
{
  "name": "John Doe"
}
```

**レスポンス:**
```json
{
  "user_id": 1,
  "name": "John Doe",
  "created_at": "2025-09-15T06:01:36Z"
}
```

**curlでのテスト:**
```bash
curl -X POST "http://localhost:8080/users" \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe"}'
```

#### GET /users - 全ユーザー取得

**レスポンス:**
```json
[
  {
    "user_id": 1,
    "name": "John Doe",
    "created_at": "2025-09-15T06:01:36Z"
  },
  {
    "user_id": 2,
    "name": "Jane Smith",
    "created_at": "2025-09-15T06:02:15Z"
  }
]
```

**curlでのテスト:**
```bash
curl http://localhost:8080/users
```

#### GET /users/{id} - 特定ユーザー取得

**レスポンス:**
```json
{
  "user_id": 1,
  "name": "John Doe",
  "created_at": "2025-09-15T06:01:36Z"
}
```

**curlでのテスト:**
```bash
curl http://localhost:8080/users/1
```

#### PUT /users/{id} - ユーザー更新

**リクエストボディ:**
```json
{
  "name": "John Smith"
}
```

**レスポンス:**
```json
{
  "user_id": 1,
  "name": "John Smith",
  "created_at": "2025-09-15T06:01:36Z"
}
```

**curlでのテスト:**
```bash
curl -X PUT "http://localhost:8080/users/1" \
  -H "Content-Type: application/json" \
  -d '{"name": "John Smith"}'
```

#### DELETE /users/{id} - ユーザー削除

**レスポンス:**
```
204 No Content
```

**curlでのテスト:**
```bash
curl -X DELETE http://localhost:8080/users/1
```

### 認証が必要なAPI

#### GET /protected - 保護されたエンドポイント

Bearer Token認証が必要です。

**ヘッダー:**
```
Authorization: Bearer your-secret-bearer-token
```

**レスポンス:**
```json
{
  "message": "Access granted to protected resource",
  "user_id": "authenticated_user"
}
```

**curlでのテスト:**
```bash
curl -H "Authorization: Bearer your-secret-bearer-token" \
  http://localhost:8080/protected
```

## プロジェクト構成

```
sample-go-echo/
├── main.go          # メインアプリケーション
├── database/        # データベース関連
│   └── database.go  # SQLite接続・マイグレーション
├── models/          # データモデル
│   └── user.go      # Userモデル・CRUD操作
├── handlers/        # APIハンドラー
│   └── user.go      # User CRUD APIハンドラー
├── docs/            # Swaggerドキュメント（自動生成）
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── sample.db        # SQLiteデータベースファイル（自動生成）
├── app              # ビルド済みバイナリ
├── go.mod           # Go modules設定
├── go.sum           # 依存関係のハッシュ
└── README.md        # このファイル
```

## 使用しているライブラリ

- [Echo v4](https://echo.labstack.com/): 高性能なWebフレームワーク
- [SQLite3](https://github.com/mattn/go-sqlite3): SQLiteデータベースドライバー
- [Squirrel](https://github.com/Masterminds/squirrel): SQLクエリビルダー
- [Swagger/OpenAPI](https://swagger.io/): APIドキュメント自動生成
- [echo-swagger](https://github.com/swaggo/echo-swagger): EchoでSwagger UIを提供

## 開発のヒント

### 新しいエンドポイントの追加

1. main.goに新しいハンドラー関数を追加
2. Swaggerアノテーションを記述
3. `swag init` でドキュメントを再生成

### Swaggerアノテーションの例

```go
// getUserByID godoc
// @Summary Get user by ID
// @Description Get user information by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func getUserByID(c echo.Context) error {
    // 実装...
}
```

## デプロイ

### Dockerを使用する場合

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

## トラブルシューティング

### よくある問題

1. **swagコマンドが見つからない**
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. **import pathのエラー**
   - go.modのモジュール名とimportパスが一致しているか確認

3. **ポート8080が使用中**
   ```bash
   # 別のポートを使用
   e.Logger.Fatal(e.Start(":8081"))
   ```

## ライセンス

MIT License