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

package http

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/cacheRedis"
	"poly-bridge/models"
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type ChainHealthController struct {
	web.Controller
}

func (c *ChainHealthController) Health() {
	var chainHealthReq models.ChainHealthReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &chainHealthReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}

	chainHealthRsp := &models.ChainHealthRsp{
		Result: make(map[uint64]bool),
	}
	logs.Info("chainHealthReq:%+v", chainHealthReq)
	for _, chainId := range chainHealthReq.ChainIds {
		chainHealthRsp.Result[chainId] = true
		var chainStatus basedef.ChainStatus
		dataStr, err := cacheRedis.Redis.Get(cacheRedis.ChainStatusPrefix + strconv.FormatUint(chainId, 10))
		if err == nil {
			err = json.Unmarshal([]byte(dataStr), &chainStatus)
		}
		if err != nil {
			logs.Error("chain %d get chain status error: %s", chainId, err)
		} else {
			if len(chainStatus.StatusTimeMap) > 0 && !chainStatus.Health {
				logs.Info("chain %d not health: %+v", chainId, chainStatus.StatusTimeMap)
				chainHealthRsp.Result[chainId] = false
			}
		}
	}
	logs.Info("chainHealthRsp:%+v", chainHealthRsp)
	c.Data["json"] = chainHealthRsp
	c.ServeJSON()
}
