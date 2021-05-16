# Memo

Goの入門
[https://www.tohoho-web.com/ex/golang.html#install]

## プロジェクトの作成

プロジェクトの作成のために、```go mod init```が必要。
よく分からない。。。

## ビルド

```go
go build ソースファイル名
```

実行ファイル名を指定したい場合は、-oオプションをつける

```go
go build -o 実行ファイル名 ソースファイル名
```

##  不明点

- main.go  の ```package main``` を別名にすると実行ファイルが作成されない
