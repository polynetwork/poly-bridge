package toolsmethod

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"poly-bridge/conf"
	"strconv"
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
		colnumber := os.Getenv("colnumber")
		if colnumber == "" {
			panic(fmt.Sprintf("colnumber is null "))
		}
		number, err := strconv.Atoi(colnumber)
		if err != nil {
			panic(fmt.Sprintf("colnumber Atoi err %v", err))
		}
		createipfsjson(nftCfg, db, number)
	} else if runflag == "2" {
		path := os.Getenv("path")
		if path == "" {
			panic(fmt.Sprintf("path is null "))
		}
		pass := os.Getenv("pass")
		if pass == "" {
			fmt.Println("pass is test")
			path = "test"
		}
		addr := os.Getenv("addr")
		if addr == "" {
			panic(fmt.Sprintf("addr is null "))
		}
		pwd, err := getPrivateKey(path, pass, addr)
		if err != nil {
			panic(fmt.Sprintf("getPrivateKey err %v", err))
		}
		signNft(nftCfg, db, pwd)
	} else if runflag == "3" {
		path := os.Getenv("path")
		if path == "" {
			panic(fmt.Sprintf("path is null "))
		}
		pass := os.Getenv("pass")
		if pass == "" {
			fmt.Println("pass is test")
			path = "test"
		}
		addr := os.Getenv("addr")
		if addr == "" {
			panic(fmt.Sprintf("addr is null "))
		}
		pwd, err := getPrivateKey(path, pass, addr)
		if err != nil {
			panic(fmt.Sprintf("getPrivateKey err %v", err))
		}
		signOtherNft(nftCfg, db, pwd)
	} else if runflag == "4" {
		path := os.Getenv("path")
		if path == "" {
			panic(fmt.Sprintf("path is null "))
		}
		pass := os.Getenv("pass")
		if pass == "" {
			fmt.Println("pass is test")
			path = "test"
		}
		addr := os.Getenv("addr")
		if addr == "" {
			panic(fmt.Sprintf("addr is null "))
		}
		pwd, err := getPrivateKey(path, pass, addr)
		if err != nil {
			panic(fmt.Sprintf("getPrivateKey err %v", err))
		}
		useraddrs := os.Getenv("useraddrs")
		if addr == "" {
			panic(fmt.Sprintf("useraddrs is null "))
		}
		signCloNft(nftCfg, pwd, useraddrs)
	} else if runflag == "-99" {
		db.Exec("DELETE FROM air_drop_nfts")
	}
}
