package main

import "fmt"

var arr1 = []int { 12, 3, 7, 0, 4, 1, 8, 0, 2, 9 }

func main() {
	fmt.Println(mergeSort(arr1))
}

func mergeSort(arr []int) []int{
	var out []int
	if( len(arr) > 1) {
		midPoint := len(arr)/2
		out = merge( mergeSort( arr[0:midPoint] ), mergeSort( arr[midPoint:] ) )
		return out;
	} else {
		return arr;
	}
	return out
}

func merge(a []int, b []int) []int{
	i := 0
	j := 0
	newArr := make( []int, 0 )
	for {
		if( i >= len(a) || j >= len(b)) {
			break
		}
		if( a[i] < b[j] ) {
			newArr = append( newArr, a[i] )
			i++
		} else {
			newArr = append( newArr, b[j] )
			j++
		}
	}
	for i < len(a) {
		newArr = append( newArr, a[i] )
		i++
	}
	for j < len(b) {
		newArr = append( newArr, b[j] )
		j++
	}
	return newArr
}