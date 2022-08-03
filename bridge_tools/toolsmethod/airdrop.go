package toolsmethod

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/conf"
	"poly-bridge/models"
)

func AirDrop(cfg *conf.Config) {
	//runflag := os.Getenv("runflag")
	//if runflag == "" {
	//	panic(fmt.Sprintf("runflag is null "))
	//}
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
	err = db.Debug().AutoMigrate(
		&models.AirDropInfo{},
	)
	if err != nil {
		panic(fmt.Sprintf("AutoMigrate AirDropInfo err:%v", err))
	}
	err = db.Debug().AutoMigrate(
		&models.TokenPriceAvg{},
	)
	if err != nil {
		panic(fmt.Sprintf("AutoMigrate TokenPriceAvg err:%v", err))
	}
	var count int64
	db.Model(&models.TokenPriceAvg{}).Count(&count)
	if count == 0 {
		tokenBasics := make([]*models.TokenBasic, 0)
		tokenPriceAvgs := make([]*models.TokenPriceAvg, 0)
		db.Find(&tokenBasics)
		for _, v := range tokenBasics {
			tokenPriceAvgs = append(tokenPriceAvgs, &models.TokenPriceAvg{
				Name:        v.Name,
				PriceAvg:    v.Price,
				UpdateTime:  v.Time,
				PriceTotal:  v.Price,
				PriceNumber: 1,
				PriceTime:   v.Time,
			})
		}
		err = db.Save(tokenPriceAvgs).Error
		if err != nil {
			panic(fmt.Sprintf("Save tokenPriceAvgs err:%v", err))
		}
	}
}
