package main

import "fmt"
//定数
//定数は変数と違って、一度値が設定されたら変更できない
const Pi = 3.14

const (
	URL      = "http://xxx.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var big int = 9223372036854775807 + 1
const big = 9223372036854775807 + 1

// iotaは定数のインクリメント(連番を返す)
const (
	C1 = iota
	C2
	C3
	C4
	C5
	C6
	C7
)

func main() {
	fmt.Println(Pi)

	// Pi = 100
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(big -1 )

	fmt.Println(C1, C2, C3, C4, C5, C6, C7)
}