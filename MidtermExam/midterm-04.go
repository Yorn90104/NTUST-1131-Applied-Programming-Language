package main

// Score: 7
// Note:
// - importing a new package is not allowed
// - math.Abs should be applied to (x + y - z), not (x + y)
/*
(10 points)
Make the function sumeq return true if x + y is equal to z and false otherwise.
The expected output is:
```
true
true
false
```

You are only allowed to add code between `// -- <code begin> --` and
`// -- <code end> --`.
*/

import (
  "fmt"
  "math"
)
  

func sumeq(x float64, y float64, z float64) bool {
  // -- <code begin> --
  const threshold = 0.0000000001
  return math.Abs(x + y) - z < threshold && z - math.Abs(x + y) < threshold 
  // -- <code end> --
}

func main() {
  fmt.Println(sumeq(0.1, 0.2, 0.3))
  fmt.Println(sumeq(0.12, 0.21, 0.33))
  fmt.Println(sumeq(0.375, 0.125, 0.5000001))
}
