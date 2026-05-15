package datatypes


type Config struct {
	TestDiscoveryPath string `json:"testDiscoveryPath"`
	JUnitXMLDirectory string `json:"jUnitXMLDirectory"`
	OutputPath string `json:"outputPath"`
}
