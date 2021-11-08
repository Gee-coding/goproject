package main

import (
	"fmt"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "FName",
		LastName:  "LName",
	}
	myMap["me"] = me

	fmt.Println(myMap["me"].FirstName)

	//slice = arrayที่ไม่ได้กำหนดขนาด
	var mySlice []string

	mySlice = append(mySlice, "C")
	mySlice = append(mySlice, "B")
	mySlice = append(mySlice, "A")

	
	sort.Strings(mySlice)/*sort*/

	fmt.Println(mySlice)
}
