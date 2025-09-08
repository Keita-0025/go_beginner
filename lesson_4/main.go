package main

import "fmt"

func main() {
	var arr1 [4]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1) //サイズの変更不可
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	// arr1[4] = 50 //エラー
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]bool{true, false, false}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3} //...で要素数を自動で決定
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr4[0])
	fmt.Println(arr4[3-1])

	arr2[3-1] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int = [4]int
	// arr5 = arr2
	// fmt.Println(arr5) //型が違うので代入できない

	fmt.Println(len(arr1))
}
