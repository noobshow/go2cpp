// Multiple return
package main

import (
	"fmt"
)

func addsub(x int) (a int, b int) {
	return x+2, x-2
}

func main() {
	y,z := addsub(4)
	fmt.Println("y=", y, "z=", z)
}
