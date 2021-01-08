/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package conf

import (
	"encoding/json"
	"fmt"
	"poly-bridge/conf"
	"poly-bridge/models"
)

type DeployConfig struct {
	Chains      []*models.Chain
	ChainFees   []*models.ChainFee
	TokenBasics []*models.TokenBasic
	TokenMaps   []*models.TokenMap
	DBConfig    *conf.DBConfig
}

func NewDeployConfig(filePath string) *DeployConfig {
	fileContent, err := conf.ReadFile(filePath)
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
