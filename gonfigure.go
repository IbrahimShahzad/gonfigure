package gonfigure

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Section struct {
	sectionCount int
	sectionName  string
	//parameters   map[string]string
}

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
