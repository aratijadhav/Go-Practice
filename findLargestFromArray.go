package main
import (
	"fmt"
)


func checkLargest(arr [5] int) int{
	max := arr[0]

	for i := 0; i < 5 ; i++ {
		if max < arr[i]{
			max = arr[i]
		}
	}

	return max
}

func main(){
	array :=[5]int{2,5,8,6,3}
	fmt.Println(array)
	fmt.Println(checkLargest(array))

}