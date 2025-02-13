package main

// Score: 7
// Note:
// - apply case on constants, not integers
// - the number of cases can be reduced
// - the number of printing functions can be reduced
/*
(10 points)
Consider five programming languages: C, CPP, Go, Java, and OCaml. The types of
these languages are identified as the following.
- C: Imperative
- CPP: Objective, Imperative,
- Go: Imperative
- Java: Objective, Imperative,
- OCaml: Functional, Objective, Imperative

Declare the five programming languages as five distinct constants between the
first pair of `// -- <code begin> --` and `// -- <code end> --` using the
"const" keyword with an iota expression. The values of these constants do not
matter. They only need to be distinct.

Add a switch statement between the second pair of `// -- <code begin> --` and
`// -- <code end> --` such that the function langType prints the types of the
input language l. For example, if the input language l is Java, the function
langType will print:
```
- Objective
- Imperative
```

The expected output of this program is:
```
Go
- Imperative
Java
- Objective
- Imperative
OCaml
- Functional
- Objective
- Imperative
```

Any if statement is forbidden. The fewer switch cases and printing functions
you use, the higher score you get.
*/

import "fmt"

// -- <code begin> --
const(
  C = iota
  CPP
  Go
  Java
  OCaml

)
// -- <code end> --

func langType(l int) {
  // -- <code begin> --
  switch l{
  case 0:
    fmt.Println("- Imperative")
  case 1:
    fmt.Println("- Objective")
    fmt.Println("- Imperative")
  case 2:
    fmt.Println("- Imperative")
  case 3:
    fmt.Println("- Objective")
    fmt.Println("- Imperative")
  case 4:
    fmt.Println("- Functional")
    fmt.Println("- Objective")
    fmt.Println("- Imperative")
  }
  // -- <code end> --
}

func main() {
  fmt.Println("Go")
  langType(Go)

  fmt.Println("Java")
  langType(Java)

  fmt.Println("OCaml")
  langType(OCaml)
}
