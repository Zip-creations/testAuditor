package parseinput

import "encoding/xml"


func ParseInput(content []byte) (TestAuditorInput, error) {
	var input TestAuditorInput
	return input, xml.Unmarshal(content, &input)
}
