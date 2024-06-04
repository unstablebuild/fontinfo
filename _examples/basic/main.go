package main

import (
	"fmt"

	"github.com/unstablebuild/fontinfo"
)

func main() {

	fonts, err := fontinfo.List()
	if err != nil {
		panic(err)
	}

	for _, font := range fonts {
		fmt.Printf("Family=%s Path=%s\n", font.Family, font.Path)
	}
}
