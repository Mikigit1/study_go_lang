package main

import ("fmt"
		"reflect"
)

//constとデータ型を使った宣言
const n int = 100

//データ型を省略した宣言
const lang = 1+4i

//まとめた宣言
const(
	x = 5
	y = 10
)

func main()  {
	fmt.Println(n)
	fmt.Println(lang)
	fmt.Println(x + y)

	v := reflect.ValueOf(lang)
	fmt.Println(v.Type())
}