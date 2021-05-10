package swapdao

import (
	"poly-bridge/conf"
	"testing"
)

var (
	configFile = "./../../conf/config_devnet.json"

	testdao *SwapDao
)

func TestMain(m *testing.M) {
	cfg := conf.NewConfig(configFile)
	testdao = NewSwapDao(cfg.DBConfig, cfg.Backup)
	m.Run()
}

func TestUpdateProfileTokenName(t *testing.T) {
	oldName := "seascape11"
	newName := "seascape"
	testdao.UpdateNFTProfileTokenName(oldName, newName)
}
