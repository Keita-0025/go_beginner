package main

import "fmt"

func main() {
	// int型
	// var i int = 10
	// var i2 int32 = 20
	// fmt.Println(i + 20)
	// // fmt.Println(i + i2) // 型が違うのでエラー
	// fmt.Printf("%T\n", i2)

	// fmt.Printf("%T\n", int64(i2)) // int32 -> int64 に変換
	// fmt.Println(i + int(i2))

	//float型
	// var fl64 float64 = 3.1
	// fmt.Println(fl64)

	// fl := 2.5
	// fmt.Println(fl64 + fl)
	// fmt.Printf("%T, %T\n", fl64, fl)

	// var fl32 float32 = 1.1 //float32はあまり使わない
	// fmt.Printf("%T\n", fl32)

	// fmt.Printf("%T\n", float64(fl32))

	//論理値型
	// var a, b bool = true, false
	// a2 := true
	// b2 := false

	// a = false
	// b2 = true

	// fmt.Println(a, b)
	// fmt.Println(a2, b2)

	//文字列型
	// var str string = "Hello"
	// str2 := "World"

	// fmt.Println(str + " " + str2)
	// fmt.Printf("%T\n", str)
	// fmt.Printf("%T, %T\n", str, str2)

	// var si string = "30"
	// si2 := "40"
	// si3 := 41
	// fmt.Printf("%T, %T, %T\n", si, si2, si3)
	// // fmt.Println(si + si3) //型が違うのでエラー
	// fmt.Println(si + si2)

	// fmt.Println(`----------------
	// test
	// 	test
	// 	-----------------`)

	// fmt.Println("\"")
	// fmt.Println(`"`)

	// fmt.Println(str[0]) // 72 (Hの文字コード)
	// fmt.Println(string(str[0]))// "H"

	//byte(uint8)型
	byteA := []byte{72, 77}
	fmt.Println(byteA)         // [72 77]
	fmt.Println(string(byteA)) //  "HM"　byte型→string型

	c := []byte("HM")
	fmt.Println(c)         // [72 77]　string型→byte型
	fmt.Println(string(c)) // "HM" byte型→string型
}
