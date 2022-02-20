// print the names of all files in which each duplic ated line occ urs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesWithDuplicatedLines := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			input := bufio.NewScanner(f)
			for input.Scan() {
				filesWithDuplicatedLines[input.Text()] = arg
				counts[input.Text()]++
			}
			// countLines(f, counts, filesWithDuplicatedLines)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filesWithDuplicatedLines[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
