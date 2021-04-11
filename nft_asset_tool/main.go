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
	"os"
	"runtime"
	"strings"
	"time"

	log "github.com/astaxie/beego/logs"
	"github.com/polynetwork/poly-nft-bridge/dao/crosschaindao"
	"github.com/polynetwork/poly-nft-bridge/utils/files"
	"github.com/urfave/cli"
)

var (
	cfgPath string
)

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "Poly NFT Bridge asset tool"
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2020 The Ontology Authors"
	app.Flags = []cli.Flag{
		LogLevelFlag,
		ConfigPathFlag,
		AssetFlag,
	}
	app.Commands = []cli.Command{
		CmdAddAsset,
		CmdDelAsset,
	}
	app.Before = beforeCommands
	return app
}

func main() {
	app := setupApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// action execute after commands
func beforeCommands(ctx *cli.Context) (err error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//logDir := ctx.GlobalString(getFlagName(LogDirFlag))
	//logFormat := fmt.Sprintf(`{"filename":"%s/deploy.log", "perm": "0777"}`, logDir)
	loglevel := ctx.GlobalUint64(getFlagName(LogLevelFlag))
	logFormat := fmt.Sprintf(`{"level:":"%d"}`, loglevel)
	if err := log.SetLogger("console", logFormat); err != nil {
		return fmt.Errorf("set logger failed, err: %v", err)
	}

	return nil
}

func handleAddAsset(ctx *cli.Context) error {
	log.Info("start to add NFT asset...")

	// load config instance
	cfg := new(AddAssetConfig)
	cfgPath := ctx.String(getFlagName(ConfigPathFlag))
	if err := files.ReadJsonFile(cfgPath, cfg); err != nil {
		return fmt.Errorf("read config json file, err: %v", err)
	}

	nowTime := time.Now().Unix()
	for _, assetBasic := range cfg.AssetBasics {
		for _, token := range assetBasic.Assets {
			token.Hash = slimHash(token.Hash)
		}
		assetBasic.Time = nowTime
	}

	dao := crosschaindao.NewCrossChainDao(cfg.Server, cfg.Backup, cfg.DBConfig)
	if err := dao.RemoveAssets(cfg.RemoveAssets); err != nil {
		return err
	}
	if err := dao.AddAssets(cfg.AssetBasics); err != nil {
		return err
	}
	if err := dao.RemoveAssetMaps(cfg.RemoveAssetMaps); err != nil {
		return err
	}

	return nil
}

func handleDelAsset(ctx *cli.Context) error {
	return nil
}

func slimHash(hash string) string {
	data := strings.TrimPrefix(hash, "0x")
	return strings.ToLower(data)
}
