package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

func main() {

	var (
		line string
		trees int
		start int = 3
	)

	// variable for our different paths
	var paths = []struct{
		right int
		down  int
	}{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	// file object
	f, err := os.Open("day3_input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// create a new buffered reader
	rdr := bufio.NewReader(f)

	// read the first line and throw it away
	// since we know we're going to advance spaces at the begining 
	line, _ = rdr.ReadString('\n')
	line = strings.Replace(line, "\n", "", 1)

	// read through the file once for each path
	for _, path := range paths {
		start = path.right
		trees = 0
		for {
			// path.down tells us how many lines to read
			for readCounter := 0; readCounter < path.down; readCounter++ {
				line, err = rdr.ReadString('\n')
				if err == io.EOF {
					// EOF
					fmt.Printf("%d trees encountered\n", trees)
				}
				line = strings.ReplaceAll(line, "\n", "")
			}

			// if start has moved past the line
			// repeat the pattern
			if start > len(line)-1 {
				start = start - len(line)
			}

			if string(line[start]) == "#" {
				trees += 1
			}

			start += path.right
		}
	}
}
