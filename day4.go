package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
	"strconv"
)

var Reset   = "\033[0m"
var Red     = "\033[31m"
var Green   = "\033[32m"

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
			fmt.Printf(Green + "\n%d valid passports\n" + Reset, validPassports)
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

			if valid == true {
				// do some data validation

				// tokenize into key:value
				tokens := strings.Fields(fields)
				for i := 0; i <= len(tokens)-1; i += 2 {
					v := tokens[i+1]

					switch tokens[i] {
					case "byr":
						year, err := strconv.Atoi(v)
						if err != nil {
							valid = false
						}
						// four digits; at least 1920 and at most 2002
						if year < 1920 || year > 2002 {
							//fmt.Printf(Red + "%v: byr (%s) not between 1920 and 2002 \u2717\n" + Reset, tokens, v)
							valid = false
						}
					case "iyr":
						year, err := strconv.Atoi(v)
						if err != nil {
							valid = false
						}
						// four digits; at least 2010 and at most 2020
						if year < 2010 || year > 2020 {
							//fmt.Printf(Red + "%v: iyr (%s) not between 2010 and 2020 \u2717\n" + Reset, tokens, v)
							valid = false
						}
					case "eyr":
						year, err := strconv.Atoi(v)
						if err != nil {
							valid = false
						}
						// four digits; at least 2020 and at most 2030
						if year < 2020 || year > 2030 {
							//fmt.Printf(Red + "%v: eyr (%s) not between 2020 and 2030 \u2717\n" + Reset, tokens, v)
							valid = false
						}
					case "hgt":
						// a number followed by either cm or in:
						// 	  If cm, the number must be at least 150 and at most 193.
						//    If in, the number must be at least 59 and at most 76.
						if strings.Contains(v, "cm") {
							hgt := strings.Split(v, "cm")[0]
							height, _  := strconv.Atoi(hgt)
							if height < 150 || height > 193 {
								//fmt.Printf(Red + "%v: height (%d) not between 150 and 193 \u2717\n" + Reset, tokens, height)
								valid = false
							}
						}

						if strings.Contains(v, "in") {
							hgt := strings.Split(v, "in")[0]
							height, _ := strconv.Atoi(hgt)
							if height < 59 || height > 76 {
								//fmt.Printf(Red + "%v: height (%d) not between 59 and 76 \u2717\n" + Reset, tokens, v)
								valid = false
							}
						}

						if !strings.Contains(v, "in") && !strings.Contains(v, "cm") {
							//fmt.Printf(Red + "%v: height %s doesn't contain cm or in \u2717\n" + Reset, tokens, v)
							valid = false
						}
					case "hcl":
						// a # followed by exactly six characters 0-9 or a-f
						hcl := strings.Split(v, "")

						if hcl[0] != "#" {
							//fmt.Printf(Red + "%v: hcl does not start with pound (%s) \u2717\n" + Reset, tokens, hcl[0])
							valid = false
						}

						if len(hcl[1:]) != 6 {
							//fmt.Printf(Red + "%v: hcl is not 6 characters long (length = %d) \u2717\n" + Reset, hcl, len(hcl[1:]))
							valid = false
						}

						if strings.ContainsAny(strings.Join(hcl[1:], ""), "ghijklmnopqrstuvwxyz") {
							//fmt.Printf(Red + "%v: contains a letter above f \u2717\n" + Reset, hcl[1:])
							valid = false
						}
					case "ecl":
						// exactly one of: amb blu brn gry grn hzl oth
						ecl := "amb blu brn gry grn hzl oth"
						if !strings.ContainsAny(v, ecl) {
							//fmt.Printf(Red + "%v: does not contain %s \u2717\n" + Reset, tokens, ecl)
							valid = false
						}
					case "pid":
						// a nine-digit number, including leading zeroes
						if len(strings.Split(v, "")) != 9 {
							//fmt.Printf(Red + "%v: %s is not 9 digits \u2717\n" + Reset, tokens, v)
							valid = false
						}

						if strings.ContainsAny(v, "abcdefghijklmnopqrstuvwxyz") {
							//fmt.Printf(Red + "%v: %s contains a letter\n" + Reset, tokens, v)
							valid = false
						}
					default:
					}
				}
			}

			if valid == true {
				fmt.Printf(Green + "%s \u2713\n" + Reset, fields)
				validPassports += 1
			}
			rec = ""
		}
	}
}
