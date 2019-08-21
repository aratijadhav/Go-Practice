package main

import (
	"fmt"
)

func findLargest(NO1,NO2 int64) int64{
	// fmt.Println(NO1,NO2)
	if NO1 > NO2 { 
		return NO1
	}
	
	return NO2

}

func main(){
	var NO1,NO2,output int64
	fmt.Scan(&NO1,&NO2)
	output = findLargest(NO1,NO2)
	fmt.Println("maximum number is:=",output)
}