/*
	Mo dif y the echo prog ram to print the index and value of each of its arguments,
	on e per line.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
