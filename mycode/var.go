
package main

import (
	"log"
	"os"
	"fmt"	
)

func test_array(x [4]float32) float32 {
	var sum float32
	for i:=0;i<len(x);i++ {
		sum += x[i]
	}
	fmt.Println("sum = ", sum)
	return sum

}

func test_slice(x []float32) float32 {
	var sum float32
	for i:=0; i<len(x);i++ {
		sum += x[i]
	}
	fmt.Println("sum = ", sum)
	return sum
}

func main() {

	dir,_ := os.Getwd()
	log.Print("current dir: ", dir)
	
	// variable declare
	// 1.	
	var a int
	var b string
	var c []float32
	var d func() bool
	var e struct {
		x int
	}
	log.Print("variable declare 1")
	log.Print("a int value: 		", a)
	log.Print("b string value: 		", b)
	log.Print("c []float32 value: 	", c)
	log.Print("d func value: 	", d())
	log.Print("e struct value: 	", e)

	/*	
	// 2.
	var (
		a2 int
		b2 string
		c2 []float32
		d2 func() bool
		e2 struct {
			x int
		}
	)
	log.Print("variable declare 2")
	*/

	// init
	var in int = 100
	var fl float32 = 100.20
	log.Print("in value: ", in)
	log.Print("fl value: ", fl)

	// array
	var arrar = [...]float32{0.1, 0.2, 0.3, 0.4}
	test_array(arrar)

	//slice
	var slice = []float32{0.5, 0.6, 0.7, 0.8}
	test_slice(slice)


}

