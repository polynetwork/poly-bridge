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

package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"poly-bridge/bridge_tools/conf"
	"poly-bridge/cacheRedis"
	serverconf "poly-bridge/conf"
	"runtime"
	"strings"
)

var (
	logLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: 1,
	}

	configPathFlag = cli.StringFlag{
		Name:  "cliconfig",
		Usage: "tools config file `<path>`",
		Value: "./bridge_tools/conf/config_transactions.json",
	}

	configServerPathFlag = cli.StringFlag{
		Name:  "configserver",
		Usage: "Server config file `<path>`",
	}

	methodFlag = cli.StringFlag{
		Name:  "method",
		Usage: "Bridge tool method",
		Value: "",
	}

	logDirFlag = cli.StringFlag{
		Name:  "logdir",
		Usage: "log directory",
		Value: "./Log-bridge_tools/",
	}

	cmdFlag = cli.UintFlag{
		Name:  "cmd",
		Usage: "which command? 1:init poly bridge 2:dump status 3:update token information 4:update bridge 5:update transactions",
		Value: 2,
	}
	dyingTokensFlag = cli.StringFlag{
		Name:  "tokenbasicname",
		Usage: "marked a token as dying by tokenbasicname",
		Value: "",
	}
	dyingTokensRisingRateFlag = cli.IntFlag{
		Name:  "proxyfee",
		Usage: "ratio for dying token",
		Value: 500,
	}
)

//getFlagName deal with short flag, and return the flag name whether flag name have short name
func getFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "poly-bridge deploy Service"
	app.Action = startServer
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2019 The Ontology Authors"
	app.Flags = []cli.Flag{
		logLevelFlag,
		configPathFlag,
		configServerPathFlag,
		logDirFlag,
		cmdFlag,
		methodFlag,
		dyingTokensFlag,
		dyingTokensRisingRateFlag,
	}
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func startServer(ctx *cli.Context) {
	cmd := ctx.GlobalInt(getFlagName(cmdFlag))
	method := ctx.GlobalString("method")
	if method != "" {
		executeMethod(method, ctx)
		return
	}
	if cmd == 1 {
		configFile := ctx.GlobalString(getFlagName(configPathFlag))
		config := conf.NewDeployConfig(configFile)
		if config == nil {
			fmt.Printf("startServer - read config failed!")
			return
		}
		configserverFile := ctx.GlobalString(getFlagName(configServerPathFlag))
		serverconfig := serverconf.NewConfig(configserverFile)
		startDeploy(config, serverconfig)
		dumpStatus(config.DBConfig)
	} else if cmd == 2 {
		configFile := ctx.GlobalString(getFlagName(configPathFlag))
		config := conf.NewDeployConfig(configFile)
		if config == nil {
			fmt.Printf("startServer - read config failed!")
			return
		}
		dumpStatus(config.DBConfig)
	} else if cmd == 3 {
		configFile := ctx.GlobalString(getFlagName(configPathFlag))
		config := conf.NewDeployConfig(configFile)
		if config == nil {
			fmt.Printf("startServer - read config failed!")
			return
		}
		configserverFile := ctx.GlobalString(getFlagName(configServerPathFlag))
		serverconfig := serverconf.NewConfig(configserverFile)
		startUpdateToken(config, serverconfig)
		dumpStatus(config.DBConfig)
	} else if cmd == 4 {
		configFile := ctx.GlobalString(getFlagName(configPathFlag))
		config := conf.NewUpdateConfig(configFile)
		if config == nil {
			fmt.Printf("startServer - read config failed!")
			return
		}
		configserverFile := ctx.GlobalString(getFlagName(configServerPathFlag))
		serverconfig := serverconf.NewConfig(configserverFile)
		startUpdate(config, serverconfig)
		dumpAffectedRows(config, config.DBConfig)
		//dumpStatus(config.DBConfig)
	} else if cmd == 5 {
		configFile := ctx.GlobalString(getFlagName(configPathFlag))
		config := conf.NewTransactionsConfig(configFile)
		if config == nil {
			fmt.Printf("startServer - read config failed!")
			return
		}
		startTransactions(config)
	} else if cmd == 6 {
		merge()
	} else if cmd == 7 {
		configServerFile := ctx.GlobalString(getFlagName(configServerPathFlag))
		serverConfig := serverconf.NewConfig(configServerFile)
		if serverConfig == nil {
			fmt.Printf("startServer - read config failed!")
			return
		}
		cacheRedis.Init()
		tokenBasicName := ctx.GlobalString(getFlagName(dyingTokensFlag))
		if tokenBasicName == "" {
			fmt.Printf("please input token name")
			return
		}
		dyingTokensRisingRate := ctx.GlobalInt(getFlagName(dyingTokensRisingRateFlag))
		SetDyingToken(tokenBasicName, dyingTokensRisingRate)
	} else if cmd == 8 {
		configServerFile := ctx.GlobalString(getFlagName(configServerPathFlag))
		serverConfig := serverconf.NewConfig(configServerFile)
		if serverConfig == nil {
			fmt.Printf("startServer - read config failed!")
			return
		}
		cacheRedis.Init()
		tokenBasicName := ctx.GlobalString(getFlagName(dyingTokensFlag))
		if tokenBasicName == "" {
			fmt.Printf("please input token name")
			return
		}
		RemoveDyingToken(tokenBasicName)
	}
}

func main() {
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
