package main

import (
	"fmt"

	"github.com/RicHardicode/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.Bigbark()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.Bigbark())
	fmt.Println(puppy.Bigbarks())
}
