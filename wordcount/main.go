package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var counts map[string]int = map[string]int{}

	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	for scn.Scan() {
		incrementWord(counts, scn.Text())
	}

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
