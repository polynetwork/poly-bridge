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
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/urfave/cli"
	"math"
	"os"
	"os/signal"
	"poly-bridge/conf"
	"poly-bridge/crosschaindao"
	"poly-bridge/crosschainlisten"
	"poly-bridge/crosschainlisten/ethereumlisten"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var chainListen *crosschainlisten.CrossChainListen

var (
	logLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: 1,
	}

	configPathFlag = cli.StringFlag{
		Name:  "cliconfig",
		Usage: "Server config file `<path>`",
		Value: "./conf/config_mainnet.json",
	}

	logDirFlag = cli.StringFlag{
		Name:  "logdir",
		Usage: "log directory",
		Value: "./Log/",
	}

	chainFlag = cli.UintFlag{
		Name:  "chain",
		Usage: "Set chain. 2:Ethereum 8:Bsc",
		Value: 100000,
	}

	heightFlag = cli.UintFlag{
		Name:  "height",
		Usage: "Set chain. 2:Ethereum 8:Bsc",
		Value: 100000,
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
	app.Usage = "listen Service"
	app.Action = StartServer
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2019 The Ontology Authors"
	app.Flags = []cli.Flag{
		logLevelFlag,
		configPathFlag,
		logDirFlag,
		chainFlag,
		heightFlag,
	}
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func StartServer(ctx *cli.Context) {
	for true {
		startServer(ctx)
		sig := waitSignal()
		stopServer()
		if sig != syscall.SIGHUP {
			break
		} else {
			continue
		}
	}
}

func startServer(ctx *cli.Context) {
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/crosschain_listen.log"}`)
	configFile := ctx.GlobalString(getFlagName(configPathFlag))
	config := conf.NewConfig(configFile)
	if config == nil {
		logs.Error("startServer - read config failed!")
		return
	}
	{
		conf, _ := json.Marshal(config)
		logs.Info("%s\n", string(conf))
	}
	chain := ctx.GlobalUint64(getFlagName(chainFlag))
	height := ctx.GlobalUint64(getFlagName(heightFlag))

	db := crosschaindao.NewCrossChainDao(config.Server, config.Backup, config.DBConfig)
	if db == nil {
		panic("server is invalid")
	}
	chainListenConfig := config.GetChainListenConfig(chain)
	if chainListenConfig == nil {
		panic("chain is invalid")
	}

	chainHandler := ethereumlisten.NewEthereumChainListenBatch(chainListenConfig)
	if chainHandler == nil {
		panic("chain handler is invalid")
	}
	chainInfo, err := db.GetChain(chainHandler.GetChainId())
	if err != nil {
		panic(err)
	}
	if height != 0 {
		chainInfo.Height = height
	}
	chainHeight, err := chainHandler.GetLatestHeight()
	if err != nil || chainHeight == 0 {
		panic(err)
	}
	logs.Info("cross chain listen, chain: %s, dao: %s......", chainHandler.GetChainName(), db.Name())
	ticker := time.NewTicker(time.Second * time.Duration(chainHandler.GetChainListenSlot()))
	for {
		select {
		case <-ticker.C:
			var height, err = chainHandler.GetLatestHeight()
			if err != nil || height == 0 || height == math.MaxUint64 {
				logs.Error("listenChain - cannot get chain %s height, err: %s", chainHandler.GetChainName(), err)
				continue
			}
			if chainInfo.Height >= height-chainHandler.GetDefer() {
				continue
			}
			logs.Info("ListenChain - chain %s latest height is %d, listen height: %d", chainHandler.GetChainName(), height, chainInfo.Height)
			for chainInfo.Height < height-chainHandler.GetDefer() {
				start := chainInfo.Height + 1
				end := start + 2000
				if end > height-chainHandler.GetDefer() {
					end = height - chainHandler.GetDefer()
				}
				logs.Info("start handle block: %d, %d", start, end)
				wrapperTransactions, srcTransactions, polyTransactions, dstTransactions, err := chainHandler.HandleNewBlock(start, end)
				if err != nil {
					logs.Error("HandleNewBlock %d err: %v", start, err)
					break
				}
				chainInfo.Height = end
				err = db.UpdateEvents(wrapperTransactions, srcTransactions, polyTransactions, dstTransactions)
				if err != nil {
					logs.Error("UpdateEvents on block %d err: %v", chainInfo.Height, err)
					break
				}
				if db.UpdateChain(chainInfo) != nil {
					logs.Error("UpdateChain [chainId:%d, height:%d] err %v", chainInfo.ChainId, chainInfo.Height, err)
					chainInfo.Height -= end
				}
			}
		}
	}
}

func waitSignal() os.Signal {
	exit := make(chan os.Signal, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(sc)
	go func() {
		for sig := range sc {
			logs.Info("cross chain listen received signal:(%s).", sig.String())
			exit <- sig
			close(exit)
			break
		}
	}()
	sig := <-exit
	return sig
}

func stopServer() {
	chainListen.Stop()
}

func main() {
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
