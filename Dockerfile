# ---------------------------------------------------
# 1. フロントエンドのビルド (Node.js環境)
# ---------------------------------------------------
FROM node:22-alpine AS frontend-builder

WORKDIR /app/frontend

# キャッシュ利用のため先にpackage.json等をコピー
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install --frozen-lockfile

# ソースコードをコピーしてビルド
COPY frontend/ ./
RUN yarn build
# -> ここで /app/frontend/dist に静的ファイルが生成される

# ---------------------------------------------------
# 2. バックエンドのビルド (Go環境)
# ---------------------------------------------------
FROM golang:1.23-alpine AS backend-builder

WORKDIR /app

# Goの依存解決
COPY go.mod go.sum ./
RUN go mod download

# Goのソースをコピーしてビルド
COPY main.go ./
# 他にGoのファイルやディレクトリがある場合は適宜コピーしてください
COPY src/ ./src/ 

# バイナリのビルド (-o main で出力)
RUN go build -o main main.go

# ---------------------------------------------------
# 3. 実行用イメージの作成 (軽量なAlpine Linux)
# ---------------------------------------------------
FROM alpine:latest

WORKDIR /app

# Goのビルド済みバイナリをコピー
COPY --from=backend-builder /app/main .

# Reactのビルド済み静的ファイルをコピー
# ディレクトリ構成を保つため frontend/dist に配置
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist

# Cloud Runのお作法：PORT環境変数を受け取る準備（Goコード側で対応が必要）
ENV PORT=8080

# 実行
CMD ["./main"]