package main

import (
	"fmt"
)

var allc [][]string

func combinate(x []string, c []string, s int, n int){
	l := len(x)
	c = append(c, x[s])
	if n == 1 {
		allc = append(allc, c)
		return
	}
	for i := s; i <= l - n ; i++ {
		combinate(x, c, i+1, n-1)
	}
}

func main() {
	x := []string{"a", "b", "c", "d", "e"}
	var c []string
	b := 3
	for a := 0; a <= len(x) - b; a++ { 
		combinate(x, c, a, b)
	}
	fmt.Println(allc)
}
