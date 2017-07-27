package main

import (
	"flag"
	"fmt"
	"io"
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

	//r := io.TeeReader(file, out)
	//b, err := ioutil.ReadAll(r)
	//l := len(b)

	l, err := io.Copy(out, file)
	if err != nil {
		panic("Unable to read file.")
	}

	if *v {
		fmt.Printf("%d bytes read.\n", l)
	}
	fmt.Println("🐱  🐱  🐱  🐱  🐱  🐱  Done  🐱  🐱  🐱  🐱  🐱  🐱")

}
