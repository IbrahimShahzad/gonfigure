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
	parameters   map[string]string
}

type Parameter struct {
	Section
	parameterCount int
	parameterName  string
	parameterValue string
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

func LoadINI(pathToFile string) ([]byte, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return []byte(""), fmt.Errorf("Cannot Load INI File")
	}
	defer file.Close()

	stemp := Section{
		sectionCount: 0,
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if isEmptyLine(line) {
			fmt.Println("Found Empty Line")
			continue
		}
		if checkComment(line) {
			fmt.Println("Found a Comment")
			continue
		}

		if checkSection(line) {
			fmt.Printf("SECTION:\t")
			stemp.sectionName = getSectionName(line)
			stemp.sectionCount = stemp.sectionCount + 1
			fmt.Println(stemp.sectionName)
			continue
		}
		if !isLetter(rune(line[0])) {
			panic("Line starts with unknown!")
		}
		fmt.Println("def: ", line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("--------------")
	return []byte("hello"), nil
}
