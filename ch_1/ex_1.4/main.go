// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// const length = len(files)
	occurences := make(map[string][]string)
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurences)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurences)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\n", occurences[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, occurences map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if (len(occurences[input.Text()]) == 0 ) || (occurences[input.Text()][len(occurences[input.Text()])-1] != f.Name()) {
			occurences[input.Text()] = append(occurences[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
