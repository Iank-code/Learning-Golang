/*
There are different types of variables, eg
	- int - Stores integers (whole numbers) such as 123 or -123
	- float32 - stores floating numbers, with decimals such as 19.99 or -19.99
	- string - stores text, such as "Hello, world". String values are surrounded by double quotes
	- bool - stores boolean values
*/

/*
Declaring variables
	1. with "var" keyword
		var variablename type = value
	2. with the ":=" sign
		variablename := value
*/

// Variable Declaration With Initial Value
/* package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
} */

// Variable Declaration Without Initial Value
// In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:

/* package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
} */

// Value Assignment After Declaration
// It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.

/* package main
import ("fmt")
func main(){
	var student1 string
	student1 = "John DOe"

	fmt.Println(student1)
} */


// Go Multiple Variable Declaration
// In Go, it is possible to declare multiple variables in the same line.


package main
import("fmt")

func main() {
	// var a,b,c int = 1,2,3

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// var a, b = 6, "Hello"
  	// c, d := 7, "World!"

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	// Multiple variable declarations can also be grouped together into a block for greater readability:

	var (
		a int
		b int = 1
		c string = "hello"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}