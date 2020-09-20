# コード

```go
package main

import "fmt"

func main(){
	fmt.Println("Hello,World!!")
}
```

# Goのコード構成  
##　起点
Goのエントリーポイントとなるのが**main関数**。  
mainパッケージに属していて、Javaで言う、mainメソッドのようなもの。  
  
Goを初期化した後は、ここから実行されることが基本。  
ただ厳密に言うと、init関数等が先に実行される。  
<https://qiita.com/tenntenn/items/7c70e3451ac783999b4f>
  
## パッケージ  
Goのプログラムは、複数のパッケージから構成されている。  
パッケージとは、関数や変数・定数などを意味のある形で一つにまとめたもの。  
  
## 関数
一連の処理を意味のある形で一つにまとめたもの。  
関数には属さない、組み込み関数（println）なども存在する。  

# Goの動かし方  
※goを自分のローカルPCに入れておく必要がある。  
  
任意のディレクトリに移動したら、下記の呪文を唱える。  
```go run main.go```！！！  
これでmain.goのソースが起動する。  
  
出力イメージはこんな感じ。  
![出力イメージ](Chapter2-Hello,Wolrd/img/Hello,World.png)
