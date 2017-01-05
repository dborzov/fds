package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var secs int64

func loadfile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("cant read the file")
		os.Exit(1)
	}
	return dat
}

func diff(org, current []byte) (bool, []byte) {
	res := make([]byte, len(current))
	changed := false
	for i := range current {
		if i < len(org) && current[i] == org[i] {
			res[i] = '-'
		} else {
			changed = true
			res[i] = current[i]
		}
	}
	return changed, res
}

func main() {
	secs = time.Now().Unix()
	prev := loadfile(os.Args[1])
	fmt.Printf("[%5d]%v", 0, string(prev))
	for {
		cur := loadfile(os.Args[1])
		changed, diffs := diff(prev, cur)
		if changed {
			prev = cur
			fmt.Printf("[%5d]%v\n", time.Now().Unix()-secs, string(diffs))
		}
	}
}
