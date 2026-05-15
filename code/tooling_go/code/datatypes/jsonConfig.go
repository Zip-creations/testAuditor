package datatypes


type Config struct {
	TestDiscoveryPath Command `json:"testDiscoveryPath"`
	JUnitXMLDirectory string `json:"jUnitXMLDirectory"`
	OutputPath string `json:"outputPath"`
}

type Command struct {
	Name string `json:"name"`
	Args []string `json:"args"`
}
