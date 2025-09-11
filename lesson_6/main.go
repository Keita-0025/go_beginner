package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		var i int = 0
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("%T\n", fl64)
		fmt.Printf("%T\n", i)

		i2 := int(fl64)
		fmt.Printf("%T\n", i2)

		fl := 10.5 //float64
		i3 := int(fl)
		fmt.Printf("%T\n", i3)
		fmt.Println(i3)
	*/

	var s string = "100"
	fmt.Printf("%T\n", s)

	i, _ := strconv.Atoi("A")
	/*
		if err != nil {
			fmt.Println(err)
		}
	*/

	fmt.Println(i)

	/*
		var i2 int = 200
		s2 := strconv.Itoa(i2)
		fmt.Println(s2)
		fmt.Printf("%T\n", s2)
	*/

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	h2 := string(b)
	fmt.Println(h2)
}
