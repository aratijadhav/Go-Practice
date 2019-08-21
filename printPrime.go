package main
import ("fmt"
		"math"
)


func printPrime(No int){
	count := 0

	for i := 2;i <= int(math.Floor(float64(No)/2));i++{
		if No % i == 0{
			// fmt.Println(No,i)
			count++
		}

	}
	
	if count >= 1{
		fmt.Println("Number is not prime ",No)
	}else{
		fmt.Println("Number is prime ",No)
	}

   
}


func main() {
	for No := 2 ; No <= 10; No++{
		printPrime(No)
	}
}