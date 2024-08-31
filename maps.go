package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "sohag", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
