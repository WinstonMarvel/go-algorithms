package main 

import "fmt"

var arr = []int { 7, 8, 9, 20, 30, 70, 100, 299, 907, 3435 }

func main() {
	fmt.Println( find(20, arr) )
}

func find( x int, arr []int ) int{
	arrLen := len(arr)
	if( arrLen > 1 ) {
		midpoint := (arrLen)/2
		if( x == arr[midpoint] ) {
			return x
		} else if( x < arr[midpoint] ) {
			return find( x, arr[0:midpoint] ) 
		} else {
			return find( x, arr[midpoint:] )
		}
	}
	return -1
}