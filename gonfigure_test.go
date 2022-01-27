package gonfigure

import "testing"

func TestCheckComment(t *testing.T) {
	if !checkComment("# this is a comment line") {
		t.Errorf("[checkComment] unable to parse comment")
	}
}

func TestCheckSection(t *testing.T) {
	if !checkSection("[ValidSection]") {
		t.Errorf("[checkSection] unable to parse section")
	}
}

func TestGetSectionName(t *testing.T) {
	sectionName := getSectionName("[ValidSection]")
	if sectionName != "ValidSection" {
		t.Errorf("[getSectionName] unable to get correct section name")
	}
}

func TestIsEmptyLine(t *testing.T) {
	if !isEmptyLine("") {
		t.Errorf("[isEmptyLine] invalid response")
	}
}

func TestIsLetter(t *testing.T) {
	if !isLetter('p') {
		t.Errorf("[isLetter] invalid response")
	}

	if isLetter('3') {
		t.Errorf("[isLetter] invalid response")
	}
}

func TestGetParameter(t *testing.T) {
	line := "parameter=value"
	param, value := getParameter(line)
	if param != "parameter" {
		t.Errorf("[getParameter] unable to get correct parameter name")
	}
	if value != "value" {

		t.Errorf("[getParameter] unable to get correct parameter value")
	}
}

func TestGetSections(t *testing.T) {
	baseValue := []string{"Section"}
	ini := make(map[string]map[string]string)
	ini["Section"] = map[string]string{}

	returnValue := GetSections(ini)

	if returnValue[0] != baseValue[0] {
		t.Errorf("[GetSections] unable to get correct sections array")
	}

}

func TestGetParametersFromSection(t *testing.T) {
	baseValue := []string{"username"}
	ini := make(map[string]map[string]string)
	ini["Section"] = map[string]string{
		"username": "password",
	}
	parameters, err := GetParametersFromSection(ini, "Section")
	if parameters[0] != baseValue[0] {
		t.Errorf("[GetParametersFromSection] unable to get correct parameters array")
	}
	if err != nil {
		t.Errorf("[GetParametersFromSection] invalid error returned from the function")
	}
}

func TestGetParameterValue(t *testing.T) {
	ini := make(map[string]map[string]string)
	ini["Section"] = map[string]string{
		"username": "password",
	}
	value, err := GetParameterValue(ini, "Section", "username")
	if value != "password" {
		t.Errorf("[GetParameterValue] unable to get correct parameter value")
	}
	if err != nil {
		t.Errorf("[GetParameterValue] invalid error returned from the function")
	}
}

func TestInsertSection(t *testing.T) {
	ini := InsertSection(InitialiseINIobj(), "Section")
	baseIni := make(map[string]map[string]string)
	baseIni["Section"] = map[string]string{}

	if _, ok := ini["Section"]; !ok {
		t.Errorf("[InsertSection] section not inserted")
	}
}
func TestWriteParameterToSection(t *testing.T) {
	ini := make(map[string]map[string]string)
	ini["Section"] = map[string]string{}
	outputIni, err := WriteParameterToSection(ini, "Section", "username", "password")
	if err != nil {
		t.Errorf("[WriteParameterToSection] invalid error returned")
	}
	pop, ok := outputIni["Section"]["username"]
	if !ok {
		t.Errorf("[WriteParameterToSection] parameter name not inserted")
	}
	if pop != "password" {
		t.Errorf("[WriteParameterToSection] parameter value not inserted")
	}
}
