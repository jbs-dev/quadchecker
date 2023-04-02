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
	// Extract the first and second command-line arguments, which represent the width (x) and height (y) of the figure.
	xStr := os.Args[1]
	yStr := os.Args[2]

	// Convert the width and height strings to integers.
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)

	// If both the width and height are greater than 0, proceed with the figure drawing.
	if x > 0 && y > 0 {
		// Loop through each row of the figure.
		for i := 1; i <= y; i++ {
			// If the current row is the first row, draw the top border.
			if i == 1 {
				// Loop through each column of the figure.
				for j := 1; j <= x; j++ {
					// If the current column is the first column, print an 'A' character.
					if j == 1 {
						z01.PrintRune('A')
						// If the current column is the last column, print a 'C' character.
					} else if j == x {
						z01.PrintRune('C')
						// Otherwise, print a 'B' character.
					} else {
						z01.PrintRune('B')
					}
				}
				// After processing each row, print a newline character.
				z01.PrintRune('\n')
				// If the current row is the last row, draw the bottom border.
			} else if i == y {
				// Loop through each column of the figure.
				for j := 1; j <= x; j++ {
					// If the current column is the first column, print an 'A' character.
					if j == 1 {
						z01.PrintRune('A')
						// If the current column is the last column, print a 'C' character.
					} else if j == x {
						z01.PrintRune('C')
						// Otherwise, print a 'B' character.
					} else {
						z01.PrintRune('B')
					}
				}
				// After processing each row, print a newline character.
				z01.PrintRune('\n')
				// If the current row is any other row, draw the sides of the figure.
			} else {
				// Loop through each column of the figure.
				for j := 1; j <= x; j++ {
					// If the current column is the first or the last column, print a 'B' character.
					if j == x || j == 1 {
						z01.PrintRune('B')
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
