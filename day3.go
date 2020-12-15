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
		start int 
		solution int = 1
	)

	// file object
	f, err := os.Open("day3_input.txt")
	if err != nil {
		panic(err)
	}

	// create a new buffered reader
	rdr := bufio.NewReader(f)

	// read the first line and throw it away
	// we will only advance forward not read anything from this line
	line, _ = rdr.ReadString('\n')
	line = strings.Replace(line, "\n", "", 1)

	var paths = []struct{
		right int
		down int
	}{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	for k, path := range paths {
		start = path.right
		trees = 0

		if k > 0 {
			// next path; reread the file

			// file object
			f, err := os.Open("day3_input.txt")
			if err != nil {
				panic(err)
			}
			rdr = bufio.NewReader(f)

			// read the first line and throw it away
			// we will only advance forward not read anything from this line
			line, _ = rdr.ReadString('\n')
			line = strings.Replace(line, "\n", "", 1)
		}

		Read:
		for {
			// read the next line(s)
			for readCounter := 0; readCounter < path.down; readCounter++ {
				line, err = rdr.ReadString('\n')
				if err == io.EOF{
					// EOF
					fmt.Printf("%d trees encountered\n", trees)
					solution *= trees
					f.Close()
					break Read
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
	fmt.Printf("Solution: %d\n", solution)
}
