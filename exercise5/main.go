package main

import (
	"fmt"
)

var hi = "hello"
var number4 uint8
var number5 int8

func main() {
	number1, number2, number3 := 747, 911, 90210
	number4 = 255
	number5 = 127
	fmt.Printf("decimal:%d \tbinary:%b \t hexadecimal:0x%x\n", number1, number1, number1)
	fmt.Printf("decimal:%d \tbinary:%b \t hexadecimal:0x%x\n", number2, number2, number2)
	fmt.Printf("decimal:%d   bi2ary:%b hexadecimal:%#x\n", number3, number3, number3)
	fmt.Printf("%d, type is %T\n", number4, number4)
	fmt.Printf("%d, type is %T", number5, number5)
} //# give the 0x to the hexadecimal
