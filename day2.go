package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type PasswordRule struct {
	MinOccurrence int
	MaxOccurrence int
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

		// get the min and max occurrence of the character
		occ := strings.Split(s[0], "-")
		min, _ := strconv.Atoi(occ[0])
		max, _ := strconv.Atoi(occ[1])

		// create a struct with the password rules
		rule = PasswordRule{
			MinOccurrence: min,
			MaxOccurrence: max,
			MustContain: s[1],
		}

		// evaluate whether the password meets the requirements
		charCount := strings.Count(s[2], s[1])
		if charCount >= rule.MinOccurrence && charCount <= rule.MaxOccurrence {
			validPasswords += 1
		}
	}
}
