/* (10 points)
 * Design a book search system.
 * - Declare a type book as a struct with the following fields:
 *   |------------|------------|
 *   | Field Name | Field Type |
 *   |------------|------------|
 *   | id         | int        |
 *   | name       | string     |
 *   | publisher  | string     |
 *   | year       | int        |
 *   |------------|------------|
 * - Declare a variable bookshelf as a map from a book id to
 *   the corresponding book. The variable bookshelf should be
 *   initialized with some books.
 * - In the main function:
 *   + Take a command-line argument as a book id
 *   + Search for the book of the id in bookshelf
 *   + If the book of the id is found, print the book details
 *   + Otherwise, print an error message
 *
 * Hints:
 * - Use strconv.Atoi to convert a string to an int
 *     strconv.Atoi(s string) (int, error)
 */

// -- <code begin> --
package main

import (
	"fmt"
	"os"
	"strconv"
)

type book struct {
	id        int
	name      string
	publisher string
	year      int
}

func main() {

	bookshelf := map[int]book{
		1: {id: 1, name: "Golang", publisher: "Google", year: 2009},
		2: {id: 2, name: "Python", publisher: "Guido van Rossum", year: 1991},
		3: {id: 3, name: "Java", publisher: "Sun Microsystems", year: 1995},
	}

	if len(os.Args) < 2 {
		fmt.Println("請提供書籍 ID 作為參數")
		return
	}
	for i := 1; i < len(os.Args); i++ {
		bookID, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if book, found := bookshelf[bookID]; found {
			fmt.Printf("id: %d\nname: %s\npublisher: %s\nyear: %d\n\n", book.id, book.name, book.publisher, book.year)
		} else {
			fmt.Println("Not found id:", bookID)
		}
	}
}

// -- <code end> --
