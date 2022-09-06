package toolsmethod

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"poly-bridge/conf"
)

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type NftJson struct {
	Description string       `json:"description"`
	ExternalUrl string       `json:"external_url"` //wang url
	Image       string       `json:"image"`        //imag url
	Name        string       `json:"name"`         //nft name
	Attributes  []*Attribute `json:"attributes"`
}

func AirDropNft(cfg *conf.Config) {
	runflag := os.Getenv("runflag")
	if runflag == "" {
		panic(fmt.Sprintf("runflag is null "))
	}
	Logger := logger.Default
	dbCfg := cfg.DBConfig
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(fmt.Sprintf("database err", err))
	}

	nftCfg := cfg.NftConfig

	//create data
	if runflag == "0" {
		initAirDropNft(db)
	} else if runflag == "1" {
		createipfsjson(nftCfg, db)
	} else if runflag == "2" {
		signNft(nftCfg, db)
	} else if runflag == "2" {
		signOtherNft(nftCfg, db)
	} else if runflag == "-99" {
		db.Exec("DELETE FROM air_drop_nfts")
	}
}
