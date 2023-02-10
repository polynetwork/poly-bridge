package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"poly-bridge/conf"
	"poly-bridge/models"
)

/* Steps
 * - createTables
 * - migrateBridgeBasicTables
 */
func zionSetUp(cfg *conf.Config) {
	dbCfg := cfg.DBConfig
	Logger := logger.Default
	if dbCfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
		dbCfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}
	//poly, err := gorm.Open(mysql.Open(dbCfg.User+":"+dbCfg.Password+"@tcp("+dbCfg.URL+")/"+
	//	dbCfg.PolyScheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	//if err != nil {
	//	panic(err)
	//}

	step := os.Getenv("STEP")
	if step == "" {
		panic("Invalid step")
	}

	switch step {
	case "createTables":
		createZionTables(db)
	case "migrateBridgeBasicTables":
		//migrateBridgeBasicTables(poly, db)
	default:
		logs.Error("Invalid step %s", step)
	}
}

func createZionTables(db *gorm.DB) {
	db.DisableForeignKeyConstraintWhenMigrating = true
	err := db.Debug().AutoMigrate(
		&models.AirDropInfo{},
		&models.AssetStatistic{},
		&models.ChainFee{},
		&models.ChainStatistic{},
		&models.Chain{},
		&models.DstSwap{},
		&models.DstTransaction{},
		&models.DstTransfer{},
		&models.LockTokenStatistic{},
		&models.NFTProfile{},
		&models.NftUser{},
		//&models.poly_details
		&models.PolyTransaction{},
		&models.PriceMarket{},
		&models.SrcSwap{},
		&models.SrcTransaction{},
		&models.SrcTransfer{},
		&models.TimeStatistic{},
		&models.TokenBasic{},
		&models.TokenMap{},
		&models.TokenPriceAvg{},
		&models.TokenStatistic{},
		&models.Token{},
		//&models.wrapper_details
		&models.WrapperTransaction{},
	)
	fmt.Println(err)
	//checkError(err, "Creating tables")
}
