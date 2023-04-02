// This program reads input from the user and tries to match it with one of the quad functions (quadA, quadB, quadC, quadD, quadE).
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Declare a slice to store the input characters.
	var arr []rune

	// Initialize a reader to read input from the standard input (os.Stdin).
	reader := bufio.NewReader(os.Stdin)
	// Read input characters one by one until an error occurs (usually EOF or end of input).
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		// Add the read character to the slice.
		arr = append(arr, char)
	}
	// Convert the slice of characters to a string.
	str := string(arr)

	// Initialize variables x and y to store the dimensions of the input.
	x := 0
	y := 0
	// Calculate the dimensions (x = width, y = height) of the input.
	for _, char := range arr {
		if char != '\n' && y == 0 {
			x++
		}
		if char == '\n' {
			y++
		}
	}
	// If the dimensions are invalid, output "Not a quad function" and exit.
	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Check if the input matches any of the quad functions and output the result.
	if isEqual(str, x, y, 'o', 'o', 'o', 'o', '-', '|') {
		fmt.Printf("[quadA] [%v] [%v]\n", x, y)
		return
	}

	if isEqual(str, x, y, '/', '\\', '\\', '/', '*', '*') {
		fmt.Printf("[quadB] [%v] [%v]\n", x, y)
		return
	}

	n := 0
	if isEqual(str, x, y, 'A', 'A', 'C', 'C', 'B', 'B') {
		n++
		fmt.Printf("[quadC] [%v] [%v]", x, y)
	}
	if isEqual(str, x, y, 'A', 'C', 'A', 'C', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[quadD] [%v] [%v]", x, y)
	}
	if isEqual(str, x, y, 'A', 'C', 'C', 'A', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[quadE] [%v] [%v]", x, y)
	}

	if n > 0 {
		fmt.Println()
		return
	}

	// If none of the quad functions matched, output "Not a quad function".
	fmt.Println("Not a quad function")
}

// isEqual is a helper function that checks if the input (str) matches the quad function defined by the given parameters.
// x and y are dimensions of the input. c1, c2, c3, and c4 are the corner characters.
// hor is the horizontal character, and ver is the vertical character.
func isEqual(str string, x, y int, c1, c2, c3, c4, hor, ver rune) bool {
	// Declare a slice to store the expected characters of the quad function.
	var arrE []rune

	// Loop through each row of the expected quad function.
	for i := 0; i < y; i++ {
		// Loop through each column of the expected quad function.
		for j := 0; j < x; j++ {
			// If the current position is on the first row.
			if i == 0 {
				// If it's the first column, append the first corner character (c1).
				if j == 0 {
					arrE = append(arrE, c1)
					// If it's the last column, append the second corner character (c2).
				} else if j == x-1 {
					arrE = append(arrE, c2)
					// If it's any other column, append the horizontal character (hor).
				} else {
					arrE = append(arrE, hor)
				}
				// If the current position is on the last row.
			} else if i == y-1 {
				// If it's the first column, append the third corner character (c3).
				if j == 0 {
					arrE = append(arrE, c3)
					// If it's the last column, append the fourth corner character (c4).
				} else if j == x-1 {
					arrE = append(arrE, c4)
					// If it's any other column, append the horizontal character (hor).
				} else {
					arrE = append(arrE, hor)
				}
				// If the current position is on any other row.
			} else {
				// If it's the first or last column, append the vertical character (ver).
				if j == 0 || j == x-1 {
					arrE = append(arrE, ver)
					// If it's any other position, append a space character.
				} else {
					arrE = append(arrE, ' ')
				}
			}
		}
		// After processing each row, append a newline character.
		arrE = append(arrE, '\n')
	}
	// Convert the slice of expected characters to a string.
	strE := string(arrE)
	// Return true if the expected string (strE) matches the input string (str); otherwise, return false.
	return strE == str
}
