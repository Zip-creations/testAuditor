package jUnit

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	input "github.com/Zip-creations/optimize_CI_deterministic_builds/src/internal/input"
)

type TestKey struct {
	Classname string
	Name      string
}

func ParseJUnitTestcaseKeys(reports input.Reports) (map[TestKey]struct{}, error) {
	seen := make(map[TestKey]struct{})
	for _, report := range reports.Reports {
		content := strings.TrimSpace(report.Content)
		if content == "" {
			continue
		}
		if err := parseJUnitTestcaseKeys(content, seen); err != nil {
			// If one report is broken, skip it and continue with the others.
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
	return seen, nil
}

func parseJUnitTestcaseKeys(content string, seen map[TestKey]struct{}) error {
	decoder := xml.NewDecoder(strings.NewReader(content))
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		startElement, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		if startElement.Name.Local != "testcase" {
			continue
		}
		var key TestKey
		for _, attr := range startElement.Attr {
			switch attr.Name.Local {
			case "classname":
				key.Classname = attr.Value
			case "name":
				key.Name = attr.Value
			}
		}
		// Don't add the key if one or both of the key identification attributes is missing in the source.
		// Testcases without those attributes set will be treaten as not found (or "not already executed")
		if key.Classname == "" || key.Name == "" {
			continue
		}
		seen[key] = struct{}{}
	}
}
