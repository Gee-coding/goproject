package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("start")
	// panicker()
	// fmt.Println("end")

	deferTest()
}

func deferTest() {
	a := "start"
	defer fmt.Println(a) //result is "start"เพราะ a ใน defer ยังเก็บค่าstart
	a = "end"
	fmt.Println(a) //result is "end"
}

func panicTest() {
	fmt.Println("start")
	//panic("something bad happened") ไม่มีcodeทำงานหลังpanic
	fmt.Println("end")

	//---------------
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil) //เปิดport 8080 จะได้ Hello Go!
	if err != nil {
		panic(err.Error()) //ให้หยุดหลังจากerror
	}
}
func deferPanic() {
	// fmt.Println("start")
	// defer fmt.Println("this was deferred")
	// panic ("somethimg bad happened")
	/*fmt.Println("end")    ไม่ถูกทำงาน */

	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")

}
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}() //invoke
	panic("something bad happened")

}
