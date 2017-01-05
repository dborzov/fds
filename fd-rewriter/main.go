package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDWR|os.O_APPEND, 0660)
	check(err)
	w := bufio.NewWriter(f)
	w.Write([]byte("Viva la resistance!"))
	w.Flush()
	for {

	}
}
