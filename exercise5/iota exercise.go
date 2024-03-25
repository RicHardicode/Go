//iota
/*Within a constant declaration,
the predeclared identifier iota represents successive untyped integer constants.
Its value is the index of the respective ConstSpec in that constant declaration,
starting at zero. It can be used to construct a set of related constants
*/

package main

import "fmt"

const (
	c0 = iota // c0 == 0
	a         // a == 1
	b         // b == 2
	c         // c == 3
	d         // d == 4
	e         // e == 5
	f         // f == 6
)
const (
	kb = iota // c0 == 0
	mb        // a == 1
	gb        // b == 2
	tb        // c == 3
	pb        // d == 4
	eb        // e == 5
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f, "\n")
	fmt.Printf("%d \t %b\n", 1<<c0, 1<<c0)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}
