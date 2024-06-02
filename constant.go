// If a variable should have a fixed value that cannot be changed, you can use the const keyword.
// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.


/* package main
import("fmt")

const PI = 3.14

func main() {
	// Multiple constants can be grouped together into a block for readability:

	const (
		A int = 1
		B = "test"
		C  = 4.3
	)

	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
} */

// The Printf() Function
/*
The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:
	%v is used to print the value of the arguments
	%T is used to print the type of the arguments
*/
package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}