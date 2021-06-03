# GoLang Study Memo

## 目標

- GoLangの基本構文を覚える
- 基本的なコーディングができるようになる(for文とかif文をサイトを見ないで書けるようになる)
- GoLangの単体テストのやり方
- GoLangでRESTAPIのプログラムを作成できるようになる
- GoLangからDBの操作する方法を学ぶ

はじめてのGo
[https://gihyo.jp/dev/feature/01/go_4beginners]

ここまで実施(構造体の埋め込み)
[https://gihyo.jp/dev/feature/01/go_4beginners/0003]

Go Web プログラミング
[https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/index.html]

とほほのGo言語入門　[完了]
[https://www.tohoho-web.com/ex/golang.html#install]

次に参考にする
[https://qiita.com/tenntenn/items/0e33a4959250d1a55045]

## メモ

### Goとは

2009年にGoogleから発表されたオープンソースのプログラミング言語  

特徴
- シンプルな言語使用のため、学習が比較的に用意。
- 豊富な標準パッケージ
- コンパイルが高速。大規模開発に適している。
- 環境に合わせたクロスコンパイル

繰り返し構文はfor文のみ。  
ifの波括弧は省略できず、三項演算子もない。  
インポートされているが、使用していない場合はコンパイルエラーになる。  
不要なコードの存在を防ぐ。  
発生した以上を戻り値として呼び出し側に返す方針としている。
-> 契約による設計が標準になっている

### プロジェクトの作成

プロジェクトの作成のために、```go mod init```が必要。
よく分からない。。。

### ビルド

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

### 変数

初期値を指定する場合、:= を用いるとvarも省略できる
(初期値により型名が明白名場合は型名も省略可能)

```go 
// var 変数名 型名　で定義
var a1 int
var a1 int = 123
var a2 = 123
a3 := 123

```

### switch文

break文はなくfallthroughがある。  
fallthroughがあると次のcase文が実行される。  
意図があって処理を継続するので、明示的に定義されるために、fallthroughがあると理解。  

### 構造体(struct)

GoLangdでは、クラスの代わりに構造体を使用。
クラスメソッドに相当する関数は関数名の前に(thisに相当する変数 *構造体名)をつけて定義する。
大文字で始まるメンバ変数は外部からアクセス可能。小文字であれば、パッケージ外からのアクセスは不可。  

構造体の生成時に値を明示的に指定しなかった場合は、ゼロ値で初期化される。  

### ゴルーチン(Goroutine)

Go言語で並行処理を実現するためのもの。スレッドよりも高速に並行処理を実現できる。

### GOPATHとGOROOT

これの影響で実行できない。
go.modの影響だったようだ。

### type

下記のような関数があった場合に、idとpriorityを逆に指定してしまった場合でもコンパイルが通ってしまう。  
このようなミスはテストで発見することができるが、意味に応じた型を持っていれば、コンパイルエラーとして問題を検出できる。  

```go
func ProcessTask(id int, priority int) {
  // 処理
}
```

Goでは、typeを使用すると型の名前、その型の定義が可能。  
関数呼び出し時に、idとpriorityが逆になっていても、コンパイルエラーとして問題を検出することが可能。  

```go
type ID int
type Priority int

func ProcessTask(id ID, priority Priority)
```




##  不明点

- main.go  の ```package main``` を別名にすると実行ファイルが作成されない
- mapの要素が存在するかどうかの調べ方が不明。
- mapの全キーを取得する処理が不明
- rangeって何。
- 複数値を返却する関数の戻り値を設定するときに、 ```go var x4 int, var y5 int = func()``` だとエラーになる
  -> 定義が間違ってた。　```go var x4, y5 int = func() ``` にすべきだった。
- 自作パッケージのところで ```go package local/mypkg is not in GOROOT ``` が出力された
- チャネルを使用するときにアローをつかるのはなぜ？
- GOPATHとGOROOTの影響で実行できない