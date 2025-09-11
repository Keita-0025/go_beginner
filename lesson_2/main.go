package main

import "fmt"

// i5 := 600 //エラーになる
func main() {
	var i int = 200
	fmt.Println(i)

	var s string = "Hello Go" + " こんにちは"
	fmt.Println(s)

	var x, y bool = true, false
	fmt.Println(x, y)

	var (
		i2 int    = 400
		s2 string = "hello Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 700
	s3 = "Go言語"
	fmt.Println(i3, s3)

	i2 = i + i3
	fmt.Println(i2)

	//暗黙的変数宣言
	i4 := 500
	fmt.Println(i4)

	i4 = 800
	fmt.Println(i4)

	//エラーになる
	// i4 := 300
	// fmt.Println(i4)

	// i3 = "hello"
	// fmt.Println(i3) //エラーになる

	// fmt.Println(i5) //エラーになる

}
