package main

import "fmt"
import "os"
import "encoding/xml"
import "path/filepath"

func main() {
	path := "./examples"  // TODO: ask path on first start of tool. NiceToHave: Make it configurable
	content, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading files: ", err)
		return
	}
	content = filterForXML(content)
	for _, entry := range content {
		filePath := filepath.Join(path, entry.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error unmarshaling XML: ", err)
			continue
		}

		var testsuite Testsuite
		err = xml.Unmarshal(data, &testsuite)
		if err != nil {
			fmt.Println("Error unmarshaling XML: ", err)
			continue
		}

		fmt.Println("Found: ", testsuite.Testcases[len(testsuite.Testcases)-1].Failure)  // Debug
	}
}

func filterForXML(files []os.DirEntry) []os.DirEntry {
	var xmlFiles []os.DirEntry
	for _, file := range files {
		if file.Name()[len(file.Name())-4:] == ".xml" {
			xmlFiles = append(xmlFiles, file)
		}
	}
	return xmlFiles
}
