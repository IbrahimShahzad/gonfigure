# Gonfigure

[![Go Reference](https://pkg.go.dev/badge/github.com/IbrahimShahzad/gonfigure.svg)](https://pkg.go.dev/github.com/IbrahimShahzad/gonfigure)

Reads and Writes ini files in golang.

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
	fmt.Println(gonfigure.GetSections(iniObj))
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
	parameters, err := gonfigure.GetParametersFromSection(iniObj, "Developer")
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
	val, err := gonfigure.GetParameterValue(iniObj, "Developer", "name")
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

	"github.com/IbrahimShahzad/gonfigure"
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
		fmt.Printf("\n")
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

## Writing INI Files

### Creating an INI object

- Create an empty ini object using `gonfigure.InitialiseINIobj()` function.

- The following returns an empty ini object which we can use to fill up

```go
newObj := gonfigure.InitialiseINIobj()
```

### Adding a section

- A new section can be added to the ini object using `gonfigure.InsertSection()` function. The function requires

  - INI object

  - Section Name as string that needs to be added

- Consider the following implementation

```go
newObj = gonfigure.InsertSection(newObj, "Ansible")
```

### Adding Key/Values to sections

- Parameters can be added to the section by calling `gonfigure.WriteParameterToSection()` function. The function requires

  - INI object

  - Section name as string
  
  - Key as string
  
  - Value as string

- Note that the `section` must exist prior to calling this function

- Consider the following implementation

```go
	newObj, err = gonfigure.WriteParameterToSection(newObj, "Ansible", "username", "value")
	if err != nil {
		panic(err)
	}
	fmt.Println(newObj)
```
### Write INI obj to File

- A filled ini object can be written to the file by calling the `gonfigure.WriteINIFile()` function. The function requires

  - Ini object

  - Path to file as string

- Consider the following implementation

```go
	err:= WriteINIFile(newObj , "./writeExample.ini")
	if err!= nil {
		panic()
	}
```

### Example code

- Following code inserts a section, parameter name and value to a file

```go
package main

include (
	"fmt"

	"gihtub.com/IbrahimShahzad/gonfigure"
)
func main(){
	newObj := gonfigure.InitialiseINIobj()
	newObj = gonfigure.InsertSection(newObj, "Ansible")
	newObj, err = gonfigure.WriteParameterToSection(newObj, "Ansible", "username", "value")
	if err != nil {
		panic(err)
	}
	fmt.Println(newObj)
	err = gonfigure.WriteINIFile(newObj, "./write.ini")
	if err != nil {
		panic(err)
	}
}
```
- Writes the following file:

```ini
[ansible]
username=value
```