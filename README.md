# Memo

Goの入門
[https://www.tohoho-web.com/ex/golang.html#install]

ここまで実施  
for文  
[https://www.tohoho-web.com/ex/golang.html#structures]

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

## 変数

初期値を指定する場合、:= を用いるとvarも省略できる
(初期値により型名が明白名場合は型名も省略可能)

```go 
// var 変数名 型名　で定義
var a1 int
var a1 int = 123
var a2 = 123
a3 := 123

```

## switch文

break文はなくfallthroughがある。  
fallthroughがあると次のcase文が実行される。  
意図があって処理を継続するので、明示的に定義されるために、fallthroughがあると理解。  

## 構造体(struct)

GoLangdでは、クラスの代わりに構造体を使用。
クラスメソッドに相当する関数は関数名の前に(thisに相当する変数 *構造体名)をつけて定義する。

##  不明点

- main.go  の ```package main``` を別名にすると実行ファイルが作成されない
- mapの要素が存在するかどうかの調べ方が不明。
- mapの全キーを取得する処理が不明
- rangeって何。
- 複数値を返却する関数の戻り値を設定するときに、 ```go var x4 int, var y5 int = func()``` だとエラーになる
  -> 定義が間違ってた。　```go var x4, y5 int = func() ``` にすべきだった。
