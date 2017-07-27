package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	in := flag.String("input", "Input File", "specify input file.")
	eo := flag.Bool("error", false, "whether the output should be printed to stderr")
	v := flag.Bool("v", false, "print extra information about the file")
	flag.Parse()

	if _, err := os.Stat(*in); os.IsNotExist(err) {
		panic("File does not exist.")
	}

	file, err := os.Open(*in)

	out := os.Stdout
	if *eo {
		out = os.Stderr
	}

	r := io.TeeReader(file, out)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic("Unable to read file.")
	}

	if *v {
		fmt.Printf("%d bytes read.\n", len(b))
	}
	fmt.Println("🐱  🐱  🐱  🐱  🐱  🐱  Done  🐱  🐱  🐱  🐱  🐱  🐱")

}
