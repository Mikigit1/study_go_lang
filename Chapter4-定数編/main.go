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

//型推論で連続して値を入れてくれる例  
const(
	a = 1 + 1
	b
	c
)

//iotaを使った例
const(
	d = 1 + iota
	e
	f
)

func main()  {
	fmt.Println(n)
	fmt.Println(lang)
	fmt.Println(x + y)

	//定数langに入れた値を型推論でデフォルトの型を調べる
	v := reflect.ValueOf(lang)
	fmt.Println(v.Type())

	//型推論で連続して値を入れてくれる例
	fmt.Println(a,b,c)

	//iotaを使った例
	fmt.Println(d,e,f)
}