package main

import (
	"fmt"
	"gonfigure"
)

func main() {
	dat, err := gonfigure.LoadINI("../test_file/single_section.ini")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(dat)
}
