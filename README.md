# Sample Go Echo API

GoとEcho frameworkを使った学習用のサンプルWebAPIです。基本的なREST APIの実装とOpenAPI（Swagger）ドキュメントの自動生成を学ぶことができます。

## 機能

- Echo v4を使用したWebサーバー
- `/hello` エンドポイントでヘルスチェック
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

```bash
go run main.go
```

サーバーが起動すると、以下のURLでアクセスできます：

- **API**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger/index.html

## API エンドポイント

### GET /hello

ヘルスチェックエンドポイント

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

## プロジェクト構成

```
sample-go-echo/
├── main.go          # メインアプリケーション
├── docs/            # Swaggerドキュメント（自動生成）
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod           # Go modules設定
├── go.sum           # 依存関係のハッシュ
└── README.md        # このファイル
```

## 使用しているライブラリ

- [Echo v4](https://echo.labstack.com/): 高性能なWebフレームワーク
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