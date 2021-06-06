FROM golang:1.16-buster

# 開発中にライブコンパイルを可能にする gin コマンドのインストール
RUN go install github.com/codegangsta/gin@latest

WORKDIR /app
COPY go.mod go.mod
RUN go mod download

