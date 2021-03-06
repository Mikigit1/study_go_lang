# 定数
定数とは…  
「データを入れる箱」のようなもの。  
難しく言うと、「メモリ上で管理するデータを保持する領域」。  
  
変数との違いは、  
「**値が変えられない**」ということ。  
変数は同じデータ型であれば、途中値を変更することは可能だが、  
定数宣言は、コンパイル時に与えられた値を変えることは出来ない。  
  
# 定数式  
定数式とは…  
**コンパイル時に計算され、値が確定される式**のこと。  
  
# 定数宣言  
定数は基本的には変数と同じように宣言する必要がある。  
しかし宣言の方法が変数とは少し違うため、別途まとめた。  
  
## constを使った宣言
変数宣言の場合は、varを利用していたが、  
定数宣言では、**const**を使って宣言することができる。  

```go
//constとデータ型を使った宣言
const n int = 100

//データ型を省略した宣言
const lang = "文字列"

//まとめた宣言
const(
	x = 5
	y = 10
)
```
## 型を明示しない場合  
定数宣言で、型を省略した書き方```const 変数名　= 値```というようなものがあった。  
Goではこの場合、型推論行い、デフォルトの型を設定する。  
  
|Name|Example|Type|
|:---|:---|:---|
|真偽値|true,false|bool|
|整数|100,10|int|
|浮動小数点数|1.5|float64|
|複素数|2+3i|complex128|
|文字列|"hoge"|string|
|ルーン|"あ","A"|rune|  
  
このルーンについては、***unicode***を扱う為に最適な型らしいのですが、  
詳しいことは調べきれていないので、参考URLを記載  
参考→「【Go】rune型という見慣れない型」<https://cres-tech.hatenablog.com/entry/2019/11/10/200000>
  
定数の型を調べるために、reflectを利用して調べることができる。  
langという定数に上記表の例をそれぞれ代入し調べてみてほしい。  
```go
package main

import ("fmt"
		"reflect"
)
//データ型を省略した宣言
const lang = 1+4i //←ここに好きな値を入れる

func main()  {
	v := reflect.ValueOf(lang)
	fmt.Println(v.Type())
} 
```
この場合複素数を入れているため、```complex128```と出力されるはず。  
  
## 右辺省略  
名前付き定数定義をまとめて宣言する場合、  
右辺を省略することができる。  
  
省略された場合、  
2つ目からは、１つ目の右辺と同じになる。  
  
```go
const(
    a = 1 + 1
    b
    c
)

func main(){
    fmt.Println(a,b,c)
}
```
この場合、a,b,cともに２が出力されるはずである。  
  
# iota（イオタ）を使う   
iotaの意味は、**「i」に相当するギリシャ文字**。  
**０から始まり、1ずつ加算される**仕組みをつけることができる。  
定数をまとめて宣言する際に、利用できる。  

```go
const(
    d = iota
    e
    f
)

func main(){
    fmt.Println(d,e,f)
}
```
上記の例であれば、```0,1,2```を出力されるはずである。  
  
また定数式の中でもiotaは利用可能。  
```go
const(
    d = 1 + iota 
    e
    f
)

func main(){
    fmt.Println(d,e,f)
}
```
先程の例に、1を加算した。  
出力イメージは、```1,2,3```である。  
dには１+0（iota）が入れられる。  
  
つづくeには1 + 2(iota)が入れられる。  
  

# 参考サイト  
「Goの定数の型宣言」<https://qiita.com/Hiraku/items/9edcb355b21f760dcee0>