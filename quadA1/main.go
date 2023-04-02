package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	// Check if the number of command-line arguments is exactly 3 (including the program name).
	if len(os.Args) != 3 {
		return
	}
	// Extract the first and second command-line arguments, which represent the width (x) and height (y) of the rectangle.
	xStr := os.Args[1]
	yStr := os.Args[2]

	// Convert the width and height strings to integers.
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)

	// If both the width and height are greater than 0, proceed with the rectangle drawing.
	if x > 0 && y > 0 {
		// Loop through each row of the rectangle.
		for i := 1; i <= y; i++ {
			// If the current row is the first or the last row, draw the top or bottom border.
			if i == 1 || i == y {
				// Loop through each column of the rectangle.
				for j := 1; j <= x; j++ {
					// If the current column is the first or the last column, print a corner character ('o').
					if j == 1 || j == x {
						z01.PrintRune('o')
						// Otherwise, print a horizontal border character ('-').
					} else {
						z01.PrintRune('-')
					}
				}
				// After processing each row, print a newline character.
				z01.PrintRune('\n')
				// If the current row is any other row, draw the sides of the rectangle.
			} else {
				// Loop through each column of the rectangle.
				for j := 1; j <= x; j++ {
					// If the current column is the first or the last column, print a vertical border character ('|').
					if j == x || j == 1 {
						z01.PrintRune('|')
						// Otherwise, print a space character.
					} else {
						z01.PrintRune(' ')
					}
				}
				// After processing each row, print a newline character.
				z01.PrintRune('\n')
			}
		}
	}
}
