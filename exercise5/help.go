package main

import (
	"fmt"
)

func main() {
	fmt.Println("😃")
	fmt.Println(`df
	sdf
	dsf
	sdf
	df`)
	var x int     //=0
	var y string  //''
	var z float32 //0.0
	a := 234      // short declare
	fmt.Println(x, y, z)
	x = 12
	y = "richard"
	z = 5.5
	fmt.Println(x, y, z, a)
	fmt.Println("====")
	a = 15
	fmt.Printf("%#x\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%b\n", x)
	//option(alt)+shift+down arrow can copy row
	fmt.Printf("%d type %T", a, a)
	//宣告func時，大小寫可以決定是否為全域可用的func或是只有區域func
}
