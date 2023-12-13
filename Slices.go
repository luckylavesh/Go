//slices are array in which we can make changes
package main
import"fmt"
func main(){
var scores = []int{100,80,90}
scores[2] = 25
//append adds new value to the existing array
scores = append(scores,99)
fmt.Println(scores, len(scores))

postt :=scores[1:3]
fmt.Println(postt)
}