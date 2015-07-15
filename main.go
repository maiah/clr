package main

import (
	"fmt"
	"os"
	"strconv"
)

const DEFAULT_LINE_COUNT = 200 // Default  lines to clear

/**
 * Clear your terminal 3x !
 */
func main() {
	lines := getLines(os.Args)
	clear(lines)
}

/**
 * Compute number of lines to clear.
 *
 * @param cargs Command line arguments
 */
func getLines(cargs []string) (lines int) {
	lines = DEFAULT_LINE_COUNT

	if len(cargs) > 1 {
		arg := cargs[1:][0]

		if len(arg) > 1 {
			firstChar := string([]rune(arg)[0])
			count := convertToInt(arg[1:len(arg)])

			switch firstChar {
			case "-":
				lines -= count
			case "+":
				lines += count
			default:
				lines = convertToInt(arg)
			}

		} else {
			lines = convertToInt(arg)
		}
	}

	return
}

/**
 * Clear terminal by lines.
 */
func clear(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Println()
	}
}

/**
 * Utility to convert string to int.
 */
func convertToInt(carg string) (count int) {
	lines, err := strconv.Atoi(carg)

	if err != nil {
		panic(err)
	}

	count = lines
	return
}
