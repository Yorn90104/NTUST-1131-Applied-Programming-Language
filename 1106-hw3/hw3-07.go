/*
 * (10 points)
 * Declare a type student as a struct with the following fields:
 * |------------|------------|
 * | Field Name | Field Type |
 * |------------|------------|
 * | SID        | string     |
 * | Name       | string     |
 * | BirthYear  | int        |
 * |------------|------------|
 * Declare three variables of type student with non-empty field
 * values. Print the three variables. Make your program complete
 * and runnable.
 */

// -- <code begin> --
package main

import "fmt"

type student struct {
	SID       string
	Name      string
	BirthYear int
}

func main() {
	student1 := student{
		SID:       "B11309044",
		Name:      "范余振富",
		BirthYear: 2006,
	}

	student2 := student{
		SID:       "B11309021",
		Name:      "林柏諺",
		BirthYear: 2006,
	}

	student3 := student{
		SID:       "B11309042",
		Name:      "吳定衡",
		BirthYear: 2006,
	}

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)

}

// -- <code end> --
