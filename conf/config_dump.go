package conf

import (
	"encoding/json"
	"fmt"
)

type DumpConfig struct {
	DBConfig *DBConfig
}

func NewDumpConfig(filePath string) *DumpConfig {
	fileContent, err := ReadFile(filePath)
	if err != nil {
		fmt.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	config := &DumpConfig{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		fmt.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	return config
}
