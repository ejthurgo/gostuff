package main

import (
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("STARTING")

	tgResize := NewTestGroup("ResizeGroup")

	for i := 0; i < 20; i++ {
		log.Printf("Adding resize test to group: %v\n", i)
		tgResize.Add(&TestOpResize{id: i})
	}
	tgResize.Run()

	log.Println("FINISHED")
}
