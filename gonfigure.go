// Author: IbrahimShahzad
//
package gonfigure

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Object to hold INI data as key-value maps
type INIobject map[string]map[string]string

type section struct {
	sectionCount int
	sectionName  string
}

// checks whether the line is a comment line or data line
//
// Checks whether the given line starts with # symbol
//
// Args
//	line as string
//
// Returns
//	True in case of comment-line
//	False in case of non-comment-line
func checkComment(line string) bool {
	if strings.HasPrefix(strings.TrimSpace(line), "#") {
		return true
	}
	return false
}

// checks whether the line is a valid "section" line or not
//
// Expects a "section" line to start with  "[ " and  end with "]"
//
// Args
//	line as string
//
// Returns
//	True in case of section-line
//	False in case of non-section-line
func checkSection(line string) bool {
	if start := strings.HasPrefix(line, "["); !start {
		return false
	}
	if stop := strings.HasSuffix(line, "]"); !stop {
		panic("Invalid Section")
	}
	return true
}

// returns section name from the valid section line
//
// Args
//	section header line as string
//
// Returns
//	section name as string
func getSectionName(sectionHeader string) string {
	return strings.TrimRight(strings.TrimLeft(sectionHeader, "["), "]")
}

// checks whether the line is empty or not
//
// Empty lines are skipped while reading the ini file
//
// Args
//	line as string
//
// Returns
//	True in case of empty line
//	False in case of non-empty line
func isEmptyLine(line string) bool {
	if strings.TrimSpace(line) == "" {
		return true
	}
	return false
}

// Checks whether the given letter is alphabetic
//
// Args
// 	first letter of line as rune
//
// Returns
//	True for alphabetic rune
//	False for non-alphabetic rune
func isLetter(letter rune) bool {
	return ('a' <= letter && letter <= 'z') || ('A' <= letter && letter <= 'Z')
}

// Splits the parameter line into key and value
//
// Args
// 	line as string
//
// Returns
// 	Key and value strings
func getParameter(line string) (string, string) {
	val := strings.Split(line, "=")
	return val[0], val[1]
}

// Get Sections From INIobject
//
// Args
// 	INI file object,
//
// Returns
//	Array of section names (strings)
//
// The call can be made as following:
//
// 		sections := gonfigure.GetSections(iniObj)
//
func GetSections(mapINI INIobject) []string {
	sections := make([]string, len(mapINI))
	iterator := 0
	for key := range mapINI {
		sections[iterator] = key
		iterator++
	}
	return sections
}

// Get Parameters From A Given Section
//
// Args
// 	INI file object,
//	section Name,
//
// Returns
//	Array of parameter names (strings)
//	Error
//
// The call can be made as following:
//
// 		parameters, err := gonfigure.GetParametersFromSection(iniObj, "sectionA")
//		if err != nil {
//			panic(err)
//		}
//
func GetParametersFromSection(mapINI INIobject, sectionName string) ([]string, error) {
	if _, ok := mapINI[sectionName]; !ok {
		return []string{""}, fmt.Errorf("Section [%v] does not exist ", sectionName)
	}
	parameters := make([]string, len(mapINI[sectionName]))
	iterator := 0
	for key := range mapINI[sectionName] {
		parameters[iterator] = key
		iterator++
	}
	return parameters, nil
}

// Get Parameters From A Given Section
//
// Args
// 	INI file object,
//	section name as string,
//	parameter name as string
//
// Returns
//	parameter value as strings
//	Error
//
// The call can be made as following:
//
// 		parameters, err := gonfigure.GetParameterValue(iniObj, "sectionA","parameterName")
//		if err != nil {
//			panic(err)
//		}
//
func GetParameterValue(mapINI INIobject, sectionName string, parameterName string) (string, error) {
	if _, ok := mapINI[sectionName]; !ok {
		return "", fmt.Errorf("Section [%v] does not exist ", sectionName)
	}
	if _, ok := mapINI[sectionName][parameterName]; !ok {
		return "", fmt.Errorf("Parameter [%v] does not exist ", parameterName)
	}
	return mapINI[sectionName][parameterName], nil
}

// Reads the ini file and loads/returns the INI object
//
// Args
//	path to file as string
//
// Returns
//	INIobj
//	Error
//
// The call can be made as following:
//
// 		iniObj, err := gonfigure.LoadINI("example_file.ini")
//		if err != nil {
//			panic(err)
//		}
//
func LoadINI(pathToFile string) (INIobject, error) {
	stemp := section{
		sectionCount: 0,
	}
	// The main ini object to return
	objINI := make(INIobject)

	file, err := os.Open(pathToFile)
	if err != nil {
		return objINI, fmt.Errorf("Cannot Load INI File: %v", pathToFile)
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount = lineCount + 1

		if isEmptyLine(line) {
			continue
		}
		if checkComment(line) {
			continue
		}

		if checkSection(line) {
			stemp.sectionName = getSectionName(line)
			stemp.sectionCount = stemp.sectionCount + 1
			objINI[stemp.sectionName] = map[string]string{}
			continue
		}
		if !isLetter(rune(line[0])) {
			return objINI, fmt.Errorf("Cannot parse INI File. Invalid parameter at line: %v", lineCount)
		} else {
			key, value := getParameter(line)
			objINI[stemp.sectionName][key] = value
			continue
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return objINI, nil
}

// Creates an empty INI object
//
// Returns
//	INIobj
//
// The call can be made as following:
//
// 		iniObj, err := gonfigure.InitialiseINIobj()
//
func InitialiseINIobj() INIobject {
	iniObj := make(INIobject)
	return iniObj
}

// Insert Section in INI object
//
// Args
//	INIobject
//	Section name as string
//
// Returns
//	INIobj
//	Error
//
// The call can be made as following:
//
// 		iniObj, err := gonfigure.InsertSection(iniObj,"ExampleSection")
//		if err != nil {
//			panic(err)
//		}
//
func InsertSection(iniObj INIobject, sectionName string) INIobject {
	iniObj[sectionName] = map[string]string{}
	return iniObj
}

// Writes Parameter to Section
//
// Args
//	INIobject
//	Section name as string
//	Parameter Name as string
//	Parameter Value as string
//	INIobject
//
// Returns
//	INIobj
//	Error
//
// The Section You are trying to write to MUST exist. See `InsertSection()` function
//
// The call can be made as following:
//
// 		iniObj, err := gonfigure.WriteParameterToSection(iniObj,"ExampleSection","username","password")
//		if err != nil {
//			panic(err)
//		}
//
func WriteParameterToSection(iniObj INIobject, sectionName string, parameter string, value string) (INIobject, error) {
	if _, ok := iniObj[sectionName]; !ok {
		return iniObj, fmt.Errorf("Section [%v] does not exist ", sectionName)
	}
	iniObj[sectionName][parameter] = value
	return iniObj, nil
}

// Writes INI section to File
//
// Args
//	output File pointer
//	Section name as string
//
// Returns
//	Error
func writeINISectionToFile(outputFile *os.File, sectionName string) error {
	outputString := "[" + strings.ToLower(sectionName) + "]\n"
	_, err := outputFile.WriteString(outputString)
	return err
}

// Writes INI parameter name and value to File
//
// Args
//	output File pointer
//	parameter name as string
//	parameter value as string
//
// Returns
//	Error
func writeINIParameterToFile(outputFile *os.File, parameterName string, parameterValue string) error {
	outputString := parameterName + "=" + parameterValue + "\n"
	_, err := outputFile.WriteString(outputString)
	return err
}

// Writes INI file
//
// Args
//	INI object
//	path to file
//
// Returns
//	Error
func WriteINIFile(iniObj INIobject, pathToFile string) error {
	// Create an output File
	outputFile, err := os.Create(pathToFile)
	if err != nil {
		return err
	}

	// Close File after function returns
	defer outputFile.Close()

	for _, section := range GetSections(iniObj) {
		// Write Section Header
		err := writeINISectionToFile(outputFile, section)
		if err != nil {
			return err
		}

		// Write Parameters
		parameters, err := GetParametersFromSection(iniObj, section)
		if err != nil {
			return err
		}
		for _, param := range parameters {
			val, err := GetParameterValue(iniObj, section, param)
			if err != nil {
				return err
			}
			err = writeINIParameterToFile(outputFile, param, val)
			if err != nil {
				return err
			}
		}
	}
}
