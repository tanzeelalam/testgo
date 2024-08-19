package main

import (
	"fmt"

	"github.com/tanzeelalam/puppy"
)

var x int = 20
var xx = 30

const PI float64 = 3.17234234

func main() {
	y := 66
	fmt.Printf("This is value of x : %v and its type %T\n", x, x)
	fmt.Printf("This is value of xx : %v and its type %T\n", xx, xx)
	fmt.Printf("This is value of PI : %v and its type %T\n", PI, PI)
	fmt.Printf("This is value of y : %v and its type %T\n", y, y)
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.Barks())
}
