package main

import "fmt"


func main() {
	fmt.Println(1+2)
	fmt.Println("Hello" + "World")

	fmt.Println(4%4)

	n := 300 //変数の宣言と代入を同時に行う 型を省略できる
	n += 200 //n = n + 200 と同じ
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)

}
