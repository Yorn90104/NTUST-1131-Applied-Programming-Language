package main

/*
 * (10 points)
 * The struct kvpair is used to store a pair of a key (of type int)
 * and a value (of type string). A slice of kvpairs can be stored
 * as a map from int to string (i.e. map[int]string). Complete the
 * function toMap such that it stores the input kvpairs as a map m.
 *
 * The expected output of this program is:
 *   map[1:Kevin 3:Alice 6:Nightmare 9527:HuaAn]
 */

import "fmt"

type kvpair struct {
	key   int
	value string
}

func toMap(pairs []kvpair) map[int]string {
	m := make(map[int]string, 10)
	// -- <code begin> --
	for _, pair := range pairs {
		m[pair.key] = pair.value
	}
	// -- <code end> --
	return m
}

func main() {
	pairs := []kvpair{
		kvpair{1, "Kevin"},
		kvpair{3, "Alice"},
		kvpair{6, "Nightmare"},
		kvpair{9527, "HuaAn"},
	}
	fmt.Println(toMap(pairs))
}
