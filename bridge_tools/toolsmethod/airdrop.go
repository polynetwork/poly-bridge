package toolsmethod

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"poly-bridge/conf"
	"poly-bridge/models"
	"time"
)

func AirDrop(cfg *conf.Config) {
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
	err = db.Model(&models.TokenPriceAvg{}).Count(&count).
		Error
	if err != nil {
		panic(fmt.Sprintf("Count TokenPriceAvg err:%v", err))
	}
	if count == 0 {
		currentTime := time.Now()
		nowDayStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()

		tokenBasics := make([]*models.TokenBasic, 0)
		tokenPriceAvgs := make([]*models.TokenPriceAvg, 0)
		db.Find(&tokenBasics)
		for _, v := range tokenBasics {
			tokenPriceAvgs = append(tokenPriceAvgs, &models.TokenPriceAvg{
				Name:        v.Name,
				PriceAvg:    v.Price,
				UpdateTime:  currentTime.Unix(),
				PriceTotal:  v.Price,
				PriceNumber: 1,
				PriceTime:   nowDayStartTime,
			})
		}
		err = db.Save(tokenPriceAvgs).Error
		if err != nil {
			panic(fmt.Sprintf("Save tokenPriceAvgs err:%v", err))
		}
	}
}

func UpdateAirDropAmount(cfg *conf.Config) {
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
	airDropInfos := make([]*models.AirDropInfo, 0)
	err = db.Find(&airDropInfos).Error
	if len(airDropInfos) == 0 {
		panic(fmt.Sprintf("Find airDropInfos err", err))
	}
	for _, v := range airDropInfos {
		v.Amount *= 100
	}
	err = db.Save(airDropInfos).Error
	if err != nil {
		panic(fmt.Sprintf("Save airDropInfos err", err))
	}

}
