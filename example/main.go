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
	fmt.Println(gonfigure.GetSections(dat))
	parameters, err := gonfigure.GetParametersFromSection(dat, "Developer")
	if err != nil {
		panic(err)
	}
	fmt.Println(parameters)
	val, err := gonfigure.GetParameterValue(dat, "Developer", "name")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}
