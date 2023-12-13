package main
import"fmt"

func main(){
	age := 50
	fmt.Println(age<=55)
	fmt.Println(age<=40)

	if age<30{
		fmt.Println("small")
	}else{
		fmt.Println("big")
	}
}