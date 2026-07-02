package parseinput

import "encoding/xml"


func ReadInput(content string) (TestAuditorInput, error) {
	var input TestAuditorInput
	err := xml.Unmarshal([]byte(content), &input)
	if err != nil {
		return input, err
	}
	return input, nil
}
