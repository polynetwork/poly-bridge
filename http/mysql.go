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
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"poly-bridge/basedef"
	"poly-bridge/conf"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var polyProxy map[string]bool

func Init() {
	config := conf.GlobalConfig.DBConfig
	Logger := logger.Default
	if conf.GlobalConfig.RunMode == "dev" {
		Logger = Logger.LogMode(logger.Info)
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.User, config.Password, config.URL, config.Scheme)
	var err error
	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	proxyMap := make(map[string]bool, 0)
	proxyConfigs := conf.GlobalConfig.ChainListenConfig
	for _, v := range proxyConfigs {
		proxyMap[strings.ToUpper(v.ProxyContract)] = true
		proxyMap[strings.ToUpper(basedef.HexStringReverse(v.ProxyContract))] = true
	}
	polyProxy = proxyMap
	if len(polyProxy) == 0 {
		panic("http init polyProxy err,polyProxy is nil")
	}
	logs.Info("http init polyProxy:", polyProxy)
}
