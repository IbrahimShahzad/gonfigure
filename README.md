# gonfigure
read and write ini files

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