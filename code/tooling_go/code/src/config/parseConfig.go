package config


type Config struct {
	TestDiscoveryPath Command `json:"testDiscoveryPath"`
	TestExecutionPath Command `json:"testExecutionPath"`
}

type Command struct {
	Command string `json:"command"`
	Args []string `json:"args"`
}
