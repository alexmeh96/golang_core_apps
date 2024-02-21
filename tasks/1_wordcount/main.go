//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	m := make(map[string]int64)

	for _, arg := range args {
		f, err := os.Open(arg)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if val, ok := m[scanner.Text()]; ok {
				m[scanner.Text()] = val + 1
			} else {
				m[scanner.Text()] = 1
			}
		}
	}

	for key, value := range m {
		if value >= 2 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}
