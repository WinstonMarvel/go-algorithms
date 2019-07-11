package main 

import "fmt"

var arr = []int { 7, 8, 9, 20, 30, 70, 100, 299, 907, 3435 }

func main() {
	fmt.Println( find(20, arr) )
}

func find( x int, arr []int ) string{
	
	for {
		arrLen := len(arr)
		midpoint := (arrLen)/2
		if( arrLen < 1 ) {
			return fmt.Sprintf("Not found")
		} else if( arr[midpoint] == x ) {
			return fmt.Sprintf("Found at index: %d", x)
		} else if ( x < arr[midpoint] ) {
			arr = arr[0:midpoint]
		} else {
			arr = arr[midpoint:]
		}
	}
	return fmt.Sprintf("Not found")
}