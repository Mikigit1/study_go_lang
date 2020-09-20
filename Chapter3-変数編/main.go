package main

import "fmt"

//「var 変数名 データ型」の例
var n int = 150

//「var 変数名」の例
var lang = "文字列"

//「まとめて宣言する」例
var (
	num1 = 10
	lang1 = "国語"
)

func main() {
	fmt.Println("「var 変数名 データ型」の例")
	fmt.Println(n)

	fmt.Println("「var 変数名」の例")
	fmt.Println(lang)

	fmt.Println("「まとめて宣言する」例")
	fmt.Println(num1)
	fmt.Println(lang1)

	//Hello関数呼び出し
	hello()
}

func hello(){
	word := "Hello"
	fmt.Println("「関数内で利用できる変数宣言例」：" + word)
}

