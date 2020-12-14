package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {

	var (
		line string
		trees int
		start int = 3
	)

	// file object
	f, err := os.Open("day3_input.txt")
	if err != nil {
		panic(err)
	}

	// create a new buffered reader
	rdr := bufio.NewReader(f)

	// read the first line and throw it away
	// since we know we're going to advance 3 spaces
	line, _ = rdr.ReadString('\n')
	line = strings.Replace(line, "\n", "", 1)

	for {
		// read the next line
		line, err = rdr.ReadString('\n')
		line = strings.ReplaceAll(line, "\n", "")
		if err != nil {
			// EOF
			fmt.Printf("%d trees encountered\n", trees)
			return
		}

		// if start has moved past the line
		// repeat the pattern
		if start > len(line)-1 {
			start = start - len(line)
		}

		if string(line[start]) == "#" {
			trees += 1
		}

		start += 3
	}

}
