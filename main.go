package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	lines := 200

	if len(os.Args) > 1 {
		arg := os.Args[1:][0]

		if len(arg) > 1 {
			firstChar := string([]rune(arg)[0])
			count, err := strconv.Atoi(arg[1:len(arg)])

			if err != nil {
				panic(err)
			}

			switch firstChar {
			case "-":
				lines -= count
			case "+":
				lines += count
			default:
				lines, err = strconv.Atoi(arg)
				if err != nil {
					panic(err)
				}
			}

		} else {
			count, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}

			lines = count
		}
	}

	for i := 0; i < lines; i++ {
		fmt.Println()
	}
}
