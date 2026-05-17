package xmlOutput

import "fmt"
import "os"
import "encoding/xml"


func WriteXMLToFile(report Testsuites, filePath string) error {
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshalling Report XML:\n %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
