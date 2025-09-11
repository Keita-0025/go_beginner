package main

import "fmt"

func main() {
	var x interface{}
	fmt.Println(x)

	x = 111
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = "H"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// x = 2
	// fmt.Println(x + 2) //エラーになる

}
