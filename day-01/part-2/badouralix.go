package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func run(s string) string {
	var input []int
	set := make(map[int]bool)

	freq := 0

	for _, str := range strings.Split(s, "\n") {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		input = append(input, n)
	}

	for {
		for _, n := range input {
			freq += n

			if _, ok := set[freq]; ok {
				return strconv.Itoa(freq)
			} else {
				set[freq] = true
			}
		}
	}
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(run(string(input)))
}
