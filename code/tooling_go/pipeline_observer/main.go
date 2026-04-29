package main

import "fmt"
import "os"

func main() {
	path := "./examples"// TODO: ask path on first start of tool. NiceToHave: Make it cofigurable
	content, err := os.ReadDir(path)
	content = filterForXML(content)
	if err != nil {
		fmt.Println("Error reading files: ", err)
		return
	}
	for _, entry := range content {
		fmt.Println(entry.Name())
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
