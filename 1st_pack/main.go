package main

import (
	"1st_pack/nd_student"
	"fmt"
)

func main() {
	name := "irfan"
	engMarks := 22
	mathMarks := 25
	results := add(engMarks, mathMarks)
	printDetails(name, results)
	fmt.Println(nd_student.Multiplications(engMarks, mathMarks))

}
