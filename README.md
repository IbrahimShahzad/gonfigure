# Gonfigure

[![Go Reference](https://pkg.go.dev/badge/github.com/IbrahimShahzad/gonfigure.svg)](https://pkg.go.dev/github.com/IbrahimShahzad/gonfigure)

Reads ini files in golang.

## Reading INI Files
### Load file

- File can be loaded from `gonfigure.LoadINI()` function. 

- It reads the given `ini` file and returns a `INIobject` which can be later used to extract sections and parameter values.

- Consider the following code for simply loading the ini file

```go
objINI, err := gonfigure.LoadINI("/pathToFile.ini")
if err != nil {
	fmt.Println(err)
	panic(err)
}
```

### Sections

- Array of sections can be returned by calling the `gonfigure.GetSections()` function. 

- It requires 
  
  - the `INIobject` as argument.

- Returns array of section names and error.
  
  - `nil` error is returned in case of success

- Consider the following example.

```go
	fmt.Println(gonfigure.GetSections(dat))
```
### Parameters

- Array of parameter names for a given section can be returned by calling the `gonfigure.GetParametersFromSection()` function. 

- It requires
  - the `INIobject` as argument.
  
  - section Name


- Returns array of parameters names and error.

  - `nil` error is returned in case of success

- Consider the following example.

```go
	parameters, err := gonfigure.GetParametersFromSection(dat, "Developer")
	if err != nil {
		panic(err)
	}
```
### Parameter Values

- Parameter Value of parameter names for a given section can be returned by calling the `gonfigure.GetParametersFromSection()` function. 

- It requires
  - the `INIobject` as argument.
  
  - section Name
  
  - parameter Name

- Returns parameter value and error.

  - `nil` error is returned in case of success

- Consider the following example.

```go
	val, err := gonfigure.GetParameterValue(dat, "Developer", "name")
	if err != nil {
		panic(err)
	}
```

### Example code

Following does all the steps mentioned in the Reading INI Files section

```go

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
```

- The input file used is `single_section.ini` in `test_file` directory

- The output of the program is as follows:

```shell
> go run main.go
Section: sectionA
Param: parameterName Value: parameterValue
Param: parameterName1 Value: parameterValue1
Param: parameterName2 Value: parameterValue2
Param: name Value: john
Section: sectionB
Param: parameterName3 Value: parameterValue1
Section: Developer
Param: age Value: 23
Param: name Value: ibrahim
```
