// package main
// import("fmt")

// func main(){
// 	for i:=1; i <= 5; i++{
// 		fmt.Println(i)
// 	}
// }


// Using range
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}