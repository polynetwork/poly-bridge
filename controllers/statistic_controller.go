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

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"poly-bridge/models"
)

type StatisticController struct {
	beego.Controller
}

func (c *StatisticController) ExpectTime() {
	var expectTimeReq models.ExpectTimeReq
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &expectTimeReq); err != nil {
		c.Data["json"] = models.MakeErrorRsp(fmt.Sprintf("request parameter is invalid!"))
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.ServeJSON()
	}
	var expectTime models.TimeStatistic
	db.Where("src_chain_id = ? and dst_chain_id = ?", expectTimeReq.SrcChainId, expectTimeReq.DstChainId).First(&expectTime)
	c.Data["json"] = models.MakeExpectTimeRsp(expectTime.SrcChainId, expectTime.DstChainId, expectTime.Time)
	c.ServeJSON()
}
