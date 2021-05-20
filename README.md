# Memo

Goの入門
[https://www.tohoho-web.com/ex/golang.html#install]

ここまで実施  
for文  
[https://www.tohoho-web.com/ex/golang.html#for]

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

スライスはJavaでいうリストと理解
```go
aa3 := []int{}
```

## switch文

break文はなくfallthroughがある。  
fallthroughがあると次のcase文が実行される。  
意図があって処理を継続するので、明示的に定義されるために、fallthroughがあると理解。  

##  不明点

- main.go  の ```package main``` を別名にすると実行ファイルが作成されない
- mapの要素が存在するかどうかの調べ方が不明。
- mapの全キーを取得する処理が不明
- rangeって何。
