package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// newRowSlice returns a new slice with the 128 row bits
func newRowSlice() []int {
	var s []int
	for i := 0; i <= 128; i++ {
		s = append(s, i)
	}
	return s
}

// newColSlice returns a new slice with the 8 column bits
func newColSlice() []int {
	var s []int
	for i := 0; i <= 7; i++ {
		s = append(s, i)
	}
	return s
}

// recursively partition the partitionScheme based on the binary key
func binaryPartition(partitionScheme string, partitionKey int, plane []int) []int {
	if len(plane) == 1 {
		return plane
	}

	switch partitionScheme[partitionKey] {
	case 'B', 'R':
		// upper half
		return binaryPartition(partitionScheme, partitionKey+1, plane[len(plane)/2:])
	case 'F', 'L':
		// lower half
		return binaryPartition(partitionScheme, partitionKey+1, plane[:len(plane)/2])
	}
	return nil
}

func main() {
	filename := "input.txt"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// create a new buffered reader
	rdr := bufio.NewReader(f)

	var seatId int
	var seats []int
	for {
		partitionScheme, err := rdr.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("Max Seat ID = %d\n", seatId)
			break
		}
		partitionScheme = strings.Replace(partitionScheme, "\n", "", 1)

		rowNumber := binaryPartition(partitionScheme, 0, newRowSlice())[0]
		colNumber := binaryPartition(partitionScheme, 7, newColSlice())[0]

		if rowNumber*8+colNumber > seatId {
			seatId = rowNumber*8 + colNumber
		}

		seats = append(seats, rowNumber*8+colNumber)
	}

	// sort the seatId's
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	for i := 0; i <= len(seats)-1; i++ {
		if seats[i+1] - seats[i] > 1 {
			fmt.Printf("Your Set ID is %d\n", seats[i]+1)
			break
		}
	}

}
