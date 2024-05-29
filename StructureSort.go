package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Printf("How manu students are in the group?")
	var number int
	fmt.Scan(&number)

	group1 := Students{
		ID: make([]int, number),
	}

	inputID(group1.ID, number)
	group1.printSructStudents()
	//fmt.Println(group1)

	sort.Ints(group1.ID)
	group1.printSructStudents()
	//fmt.Println(group1)

	removeDuplicates(group1.ID)
	group1.printSructStudents()
	//fmt.Println(group1)

}

type Students struct {
	ID []int
}

func (s Students) printSructStudents() {
	fmt.Println("Print structure. Students have ID #", s.ID)
}

func inputID(group1 []int, number int) {
	var inputStudentID int
	for i := 0; i < number; i++ {

		fmt.Printf("Input ID of student #%v from group#1  -- ", i)
		fmt.Scan(&inputStudentID)

		group1[i] = inputStudentID

	}
}

func removeDuplicates(group1 []int) {
	for n := 0; n < len(group1)-1; n++ {
		if group1[n] == group1[n+1] {
			group1 = append(group1[:n], group1[n+1:]...)
		}

	}

}
