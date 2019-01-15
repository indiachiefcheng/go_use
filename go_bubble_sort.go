package main

import "fmt"

func main(){
	//定义一个切片
	var values = []int{0,2,1,4,5,7,1,0}
	bubbleSort(values)
	for _,value := range(values){
		fmt.Println(value)
	}
}

func BubbleSort(values []int){
	flag:=true
	for i:=0;i<len(values)-1;i++{
		flag=true
		for j :=0;j<len(values)-1-i;j++{
			if values[j]>values[j+1]{
				values[j],values[j+1]=values[j+1],values[j]
				flag=false
			}
		}
		if flag==true{
			break
		}
	}
}

func bubbleSort(theArray []int)[]int{
	for i:=0;i<len(theArray)-1;i++{
		for j:=0;j<len(theArray)-1-i;j++{
			if theArray[j] > theArray[j+1]{
				theArray[j], theArray[j+1] = theArray[j+1], theArray[j]
			}
		}
	}
	return theArray
}
