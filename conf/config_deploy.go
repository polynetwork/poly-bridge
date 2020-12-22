package conf

import (
	"encoding/json"
	"fmt"
	"poly-swap/models"
)

type DeployConfig struct {
	Chains      []*models.Chain
	ChainFees   []*models.ChainFee
	TokenBasics []*models.TokenBasic
	TokenMaps   []*models.TokenMap
	DBConfig    *DBConfig
}

func NewDeployConfig(filePath string) *DeployConfig {
	fileContent, err := ReadFile(filePath)
	if err != nil {
		fmt.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	config := &DeployConfig{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		fmt.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	return config
}
