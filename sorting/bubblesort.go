package main

import "fmt"


var arr = [10]int { 12, 3, 7, 0, 4, 1, 8, 0, 2, 9 }


func main(){
	fmt.Println("Hello World")
	arrLength:= len(arr)
	for i:= 0; i< arrLength -1; i++ {
		for j:= 0; j< arrLength -1; j++ {
			if(arr[j] > arr[j+1]){
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	fmt.Println(arr)
}


func swap(a *int, b *int){
	x := *a;
	*a = *b;
	*b = x;
}