package main

import "fmt"

func main() {
	//例①
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("=====")

	//例②
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("=====")

	//例③
	stringArray := [3]string{"鈴木", "佐藤", "加藤"}
	for i, s := range stringArray {
		fmt.Printf("index: %d, name:%s\n", i, s) //正規表現の利用（今回は無視でOK）
	}
	stringArray2 := [3]string{"愛川", "一二三", "浮世絵"}
	for _, s := range stringArray2 {
		fmt.Printf("name:%s\n", s) //正規表現の利用（今回は無視でOK）
	}
}
