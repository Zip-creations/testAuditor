package xmlOutput

import "fmt"
import "os"
import "encoding/xml"
import "path/filepath"


func WriteXMLToFile(report Testsuites, filePath string) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("Error creating directories for path %s:\n %w", dir, err)
	}
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshalling Report XML:\n %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
