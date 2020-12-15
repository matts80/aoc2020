package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
)

func main() {
	// file object
	f, err := os.Open("day4_input.txt")
	if err != nil {
		panic(err)
	}

	// create a new buffered reader
	rdr := bufio.NewReader(f)

	var (
		buf, rec string
		validPassports int
	)

	requiredFields := []string{"ecl", "pid", "eyr", "byr", "iyr", "hgt", "hcl"}

	// read up until a blank line
	for {
		valid := true

		buf, err = rdr.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("%d valid passports\n", validPassports)
			break
		}
		buf = strings.Replace(buf, "\n", " ", 1)

		// concat the lines until we get to a blank line
		if buf !=  " " {
			rec += buf
		} else {
			// when we hit a blank line we have reached the end of a record
			fields := strings.Join(strings.Split(rec, ":"), " ")

			for _, requiredField := range requiredFields {
				if !strings.Contains(fields, requiredField) {
					valid = false
				}
			}

			rec = ""

			if valid == true {
				validPassports += 1
			}
		}
	}

}
