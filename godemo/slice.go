package main

import "fmt"

func main() {
	var s []byte
	s = make([]byte, 6, 9) //([]type, len,cap)
	s = []byte{2, 4, 6, 8, 1, 3, 5, 7}

	fmt.Println(len(s), cap(s), s[2:5])

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}
}
