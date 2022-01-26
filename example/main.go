package main

import (
	"fmt"
	"gonfigure"
)

func main() {
	objINI, err := gonfigure.LoadINI("../test_file/single_section.ini")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, section := range gonfigure.GetSections(objINI) {
		fmt.Printf("Section: %v\n", section)
		parameters, err := gonfigure.GetParametersFromSection(objINI, section)
		if err != nil {
			panic(err)
		}
		for _, param := range parameters {
			val, err := gonfigure.GetParameterValue(objINI, section, param)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Param: %v Value: %v\n", param, val)
		}
	}
}
