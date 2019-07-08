package main

import "fmt"

var arr = [10]int { 12, 3, 7, 0, 4, 1, 8, 0, 2, 9 }

func main(){
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++{
			if(arr[minIndex] > arr[j]){
				minIndex = j
			}
		}
		swap( &arr[minIndex], &arr[i])
	}
	fmt.Println(arr)
}


func swap(a *int, b *int){
	x := *a
	*a = *b
	*b = x
}