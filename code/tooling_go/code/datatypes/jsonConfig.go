package datatypes


type Config struct {
	TestDiscoveryPath Command `json:"testDiscoveryPath"`
	JUnitXMLDirectory string `json:"jUnitXMLDirectory"`
	OutputPath string `json:"outputPath"`
}

type Command struct {
	Command string `json:"command"`
	Args []string `json:"args"`
}
