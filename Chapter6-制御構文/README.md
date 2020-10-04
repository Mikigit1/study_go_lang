# For文（ループ）  
Goでは繰り返しを表す文はFor文しか存在しない。  
（while文がない）  
  
For文を構成する条件は3つ。（他言語と同じ）  
```
①初期化：最初の繰り返しを行う際の初期値を設定  
②条件式：各繰り返しで評価される
③処理文：各繰り返し後に実行される処理
```  
各要素は、```;```で区切られる。  

【例①】
```go
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
```
## 省略形  
①初期化と③処理文は、省略が可能である。  
こちらに関してはwhile文に似た記述である。  
【例②】  
```go
	sum := 1 //関数内でしか使えない宣言
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
```

## rangeを使った記述  
拡張for文を使った書き方は、Goではrangeを使って記述することができる。  
【例③】
```go
	stringArray := [3]string{"鈴木", "佐藤", "加藤"}
	for i, s := range stringArray {
		fmt.Printf("index: %d, name:%s\n", i, s) //正規表現の利用（今回は無視でOK）
	}
```
配列型からfor文range句を使って取り出す際、index番号が振られて取り出される。  
このindexを無視する場合は以下のように、_をつける
```go
	stringArray2 := [3]string{"愛川", "一二三", "浮世絵"}
	for _, s := range stringArray2 {
		fmt.Printf("name:%s\n", s) //正規表現の利用（今回は無視でOK）
	}
```