// Single-Case switch Syntax
/*
	switch expression{
	case X:
		code block
	case Y:
		code block
	case Z:
		code block
	default:
		code block
	}
*/
package main
import("fmt")

func main() {
	day := 7

	switch day{
	case 1:
		fmt.Println("Monday" )
	case 2:
		fmt.Println("Tuesday" )
	case 3:
		fmt.Println("Wednesday" )
	case 4:
		fmt.Println("Thursday" )
	case 5:
		fmt.Println("Friday" )
	case 6:
		fmt.Println("Saturday" )
	default:
    	fmt.Println("Not a weekday")
	}
}

