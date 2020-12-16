package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
	"strconv"
	"regexp"
)

var (
	Reset = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	X = "\n\t" + Red + "\u2717 " + Reset
	Check = Green + "\u2713 " + Reset
)

type validator func(string) bool

var requiredFields = []string{"ecl", "pid", "eyr", "byr", "iyr", "hgt", "hcl"}
var validEcl = "amb blu brn gry grn hzl oth"
var validPassport = map[string]validator {
	"byr": func(value string) bool {
		year, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf(X + "byr (%s) is not a number" + Reset, value)
			return false
		}

		if 1920 < year && year > 2002 {
			fmt.Printf(X + "byr (%s) is not between 1920 and 2002" + Reset, value)
			return false
		}
		return true
	},
	"iyr": func(value string) bool {
		year, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf(X + "iyr (%s) is not a number" + Reset, value)
			return false
		}

		if 2010 < year && year > 2020 {
			fmt.Printf(X + "iyr (%s) is not between 2010 and 2020" + Reset, value)
			return false
		}
		return true
	},
	"eyr": func(value string) bool {
		year, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf(X + "eyr (%s) is not a number" + Reset, value)
			return false
		}

		if 2020 < year && year > 2030 {
			fmt.Printf(X + "eyr (%s) is not between 2020 and 2030" + Reset, value)
			return false
		}
		return true
	},
	"hgt": func(value string) bool {
		return true
	},
	"hcl": func(value string) bool {
		return true
	},
	"ecl": func(value string) bool {
		if !strings.ContainsAny(value, validEcl) {
			fmt.Printf(X + "invalid eye color" + Reset)
			return false
		}
		return true
	},
	"pid": func(value string) bool {
		matched, _ := regexp.Match(`\d{9}`, []byte(value)) 
		if !matched {
			fmt.Printf(X + "pid (%s) is not a 9-digit number" + Reset, value)
			return false
		}
		return true
	},
}

func isValidPassport(passport map[string]string) bool {
	for key, validatorFunc:= range validPassport {
		value, ok := passport[key]
		if !ok {
			continue
		}
		if !validatorFunc(value) {
			return false
		}
	}
	return true
}

func main() {

	// file object
	f, err := os.Open("day4_test.txt")
	if err != nil {
		panic(err)
	}

	// create a new buffered reader
	rdr := bufio.NewReader(f)

	var (
		buf, rec string
		validPassports int
	)

	// read up until a blank line
	passport := make(map[string]string)
	for {
		valid := true

		buf, err = rdr.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("\n%d valid passports\n" , validPassports)
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

			// part 2
			if valid == true {
				// do data validation
				// build a map with key value pairs
				tokens := strings.Fields(fields)
				fmt.Printf("\nTesting %v \u21e8 ", tokens)

				for i := 0; i <= len(tokens)-1; i += 2{
					passport[tokens[i]] = tokens[i+1]

					if tokens[i] == "cid" {
						// ignore
						continue
					}

					// test for validity
					if !isValidPassport(passport) {
						valid = false
					}
					// remove the key for the next run
					delete(passport, tokens[i])
				}
			}
			if valid == true {
				fmt.Printf(Check + "\n" + Reset)
				validPassports += 1
			}
			rec = ""
		}
	}
}
