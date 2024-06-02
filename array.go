// Declare an Array
// 1. with the "var" keyword:
	/*
		var array_name = [length]datatype{values}

			package main
			import("fmt")
		
			func main(){
				var arr1 = [3]int{1, 2, 3}
				arr2 := [5]int{4, 5, 6, 7, 8}
		
				fmt.Println(arr1)
				fmt.Println(arr2)
			}

		or

		var array_name = [...]datatype{values}
			package main
			import("fmt")
		
			func main(){
				var arr1 = [...]int{1, 2, 3}
				arr2 := [...]int{4, 5, 6, 7, 8}
		
				fmt.Println(arr1)
				fmt.Println(arr2)
			}
	*/


// 2. with the ":=" keyword
	/*
		array_name := [length]datatype{values}
		or
		array_name := [...]datatype{values}
	*/
	

// Initialize Only Specific Elements
package main
import("fmt")
func main(){
	arr1 := [5]int{0:5, 3:13, 4:7}
	fmt.Println(arr1)
	fmt.Println(len(arr1))
}