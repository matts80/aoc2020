package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type PasswordRule struct {
	Pos1Occurrence int
	Pos2Occurrence int
	MustContain   string
}

func main() {
	// file object
	f, err := os.Open("day2_input.txt")
	if err != nil {
		panic(err)
	}

	// create a new buffered reader
	rdr := bufio.NewReader(f)

	var (
		line string
		rule PasswordRule
		validPasswords int
	)

	for {
		line, err = rdr.ReadString('\n')
		if err != nil {
			// EOF
			fmt.Printf("%d valid passwords\n", validPasswords)
			return
		}
		s := strings.Split(line, " ")

		// get rid of the newline
		s[2] = strings.Replace(s[2], "\n", "", 1)

		// get rid of the : in the second item
		s[1] = strings.Replace(s[1], ":", "", 1)

		// get the positional requirements
		occ := strings.Split(s[0], "-")
		pos1, _ := strconv.Atoi(occ[0])
		pos2, _ := strconv.Atoi(occ[1])

		// create a struct with the password rules
		rule = PasswordRule{
			Pos1Occurrence: pos1,
			Pos2Occurrence: pos2,
			MustContain: s[1],
		}

		// evaluate whether the password meets the requirements
		if string(s[2][rule.Pos1Occurrence-1]) == rule.MustContain && string(s[2][rule.Pos2Occurrence-1]) == rule.MustContain {
			// invaid: rule.MustContain is in both places
			continue
		}

		if string(s[2][rule.Pos1Occurrence-1]) == rule.MustContain || string(s[2][rule.Pos2Occurrence-1]) == rule.MustContain {
			// valid: it's in one or the other location
			validPasswords += 1
			continue
		}
	}
}
