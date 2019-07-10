package main

import "fmt"

var arr = []string { "Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipisicing", "elit", "sed", "do" }

func main() {
	fmt.Println( find("sit", arr) )
}

func find( str string, arr []string ) int{
	for i := 0; i < len(arr); i++ {
		if( arr[i] == str ) {
			return i
		}
	}
	return -1
}