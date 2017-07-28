package main

import (
	"bufio"
	"fmt"
	"os"
)

// Usage:
//   echo "foo bar baz foo bar foo" | go run main.go
//   go run main.go Ctrl+D to escape
func main() {

	var counts map[string]int = map[string]int{}

	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	for scn.Scan() {
		incrementWord(counts, scn.Text())
	}

	fmt.Println(counts)

}

func incrementWord(m map[string]int, w string) int {
	for k, v := range m {
		if k == w {
			m[k] = v + 1
			return v + 1
		}
	}
	m[w] = 1
	return 1
}
