package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Capitalize(input string) (output string) {
	var input_array = strings.Split(input, "")

	sep := regexp.MustCompile("[^A-Za-z0-9]")

	for i := 0; i < len(input); i++ {
		var should_be_capital bool

		if i == 0 || sep.MatchString(input_array[i-1]) {
			should_be_capital = true
		}

		if should_be_capital {

			output += strings.ToUpper(input_array[i])

		} else {

			output += strings.ToLower(input_array[i])

		}

	}

	return output
}

func main() {
	args := os.Args[1:]
	input := strings.Join(args, " ")

	fmt.Println(Capitalize(input))
}
