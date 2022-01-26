package gonfigure

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkComment(line string) bool {
	if strings.HasPrefix(strings.TrimSpace(line), "#") {
		return true
	}
	return false
}

func checkSection(line string) bool {
	if start := strings.HasPrefix(line, "["); !start {
		return false
	}
	if stop := strings.HasSuffix(line, "]"); !stop {
		panic("Invalid Section")
	}
	return true
}

func getSectionName(sectionHeader string) string {
	return strings.TrimRight(strings.TrimLeft(sectionHeader, "["), "]")
}

func isEmptyLine(line string) bool {
	if strings.TrimSpace(line) == "" {
		return true
	}
	return false
}

func isLetter(letter rune) bool {
	return ('a' <= letter && letter <= 'z') || ('A' <= letter && letter <= 'Z')
}
func getParameter(line string) (string, string) {
	val := strings.Split(line, "=")
	return val[0], val[1]
}
func GetSections(mapINI map[string]map[string]string) []string {
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
// 		INI file object,
//		section Name,
//
// Returns
//		Array of parameter names (strings)
//		Error
//
// The call can be made as following:
//
// 		parameters, err := gonfigure.GetParametersFromSection(dat, "sectionA")
//		if err != nil {
//			panic(err)
//		}
//
func GetParametersFromSection(mapINI map[string]map[string]string, sectionName string) ([]string, error) {
	if _, ok := mapINI[sectionName]; !ok {
		return []string{""}, fmt.Errorf("Section [%v] does not exist", sectionName)
	}
	parameters := make([]string, len(mapINI[sectionName]))
	iterator := 0
	for key := range mapINI[sectionName] {
		parameters[iterator] = key
		iterator++
	}
	return parameters, nil
}

func GetParameterValue(mapINI map[string]map[string]string, sectionName string, parameterName string) (string, error) {
	if _, ok := mapINI[sectionName]; !ok {
		return "", fmt.Errorf("Section [%v] does not exist", sectionName)
	}
	if _, ok := mapINI[sectionName][parameterName]; !ok {
		return "", fmt.Errorf("Parameter [%v] does not exist", parameterName)
	}
	return mapINI[sectionName][parameterName], nil
}

func LoadINI(pathToFile string) (map[string]map[string]string, error) {
	stemp := Section{
		sectionCount: 0,
	}
	//stemp.parameters = make(map[string]string)
	globalMap := make(map[string]map[string]string)

	file, err := os.Open(pathToFile)
	if err != nil {
		return globalMap, fmt.Errorf("Cannot Load INI File:%v", pathToFile)
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
			globalMap[stemp.sectionName] = map[string]string{}
			continue
		}
		if !isLetter(rune(line[0])) {
			return globalMap, fmt.Errorf("Cannot parse INI File. invalid parameter at line:%v", lineCount)
		} else {
			key, value := getParameter(line)
			globalMap[stemp.sectionName][key] = value
			continue
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return globalMap, nil
}
