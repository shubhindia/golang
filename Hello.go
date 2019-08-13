// Hello
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//here we are defining the datatype too
	/*var x int

	var y int
	*/

	//we can automate the defination as mentioned below. We use ":=" so that it
	//automatically takes the datatype
	x := 1
	y := 2

	fmt.Println("x=", x)
	fmt.Println("y=", y)

	var mean int
	mean = (x + y) / 2
	fmt.Println("Mean is ", mean)
}
