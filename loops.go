//we always use for keyword for loop in go
package main
import"fmt"
func main(){
	x :=0
	for x<5{
		fmt.Println("value of x is :", x)
		x++
	}

	//another method
	for y:=0; y < 6; y++ {
		fmt.Println("values" , y)
	}
	//loops for slices
	 nnmes := []string{"lavesh", "soni", "hello"}
	   for i:=0; i<len(nnmes); i++ {
		fmt.Println(nnmes[i])
	   }
}
