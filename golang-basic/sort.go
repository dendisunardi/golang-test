package main

import (
	"fmt"
	"sort"
)

func main() {
	files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
	target := "Makefile"
	sort.Strings(files)
	i := sort.Search(len(files),
		func(i int) bool { return files[i] >= target })
	if i < len(files) && files[i] == target {
		fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
	}

}
