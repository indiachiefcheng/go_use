package main

/*
函数作为类型传递案例
*/



import (
	"fmt"
)

func isOdd(value int) bool{
	if value%2 == 0{
		return false
	}
	return true
}

func isEven(value int) bool{
	if value%2 == 0{
		return true
	}
	return false
}

type boolFunc func(int) bool

func filter(slice []int,ff boolFunc) []int{
	var result []int
	for _,value := range slice{
		if ff(value){
			result = append(result,value)
		}
	}
	return result
}

func main(){
	slice := []int{3,2,1,5,6,8,4}
	fmt.Pirntln("slice=",slice)
	odd := filter(slice,isOdd)
	fmt.Pirntln("odd=",odd)
	even := filter(slice,isEven)
	fmt.Pirntln("even=",even)
}




