package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sptbackend/config"
	"time"
)



type Order struct {
	Oid int `db:"oid"`
	Address string `db:"address"`
	Duration int `db:"duration"`
	Effective int `db:"effective"`
	Amount int64 `db:"amount"`
	Gettoday int64 `db:"gettoday"`
	CreateTime time.Time `db:"create_time"`
	Distribution int64 `db:"distribution"`
}

type PledgeType struct {
	Tid int `db:"tid"`
	Days int `db:"days"`
	Rate float32 `db:"rate"`
}

var DB *gorm.DB
func InitDatabase()  {
	sqlSource := config.Cfg.Database.DBUserName + ":" +config.Cfg.Database.DBPassword + "@tcp(" + config.Cfg.Database.DBHost + ":" + config.Cfg.Database.DBPort + ")/"+config.Cfg.Database.DBSchema+"?"+config.Cfg.Database.DBArgs
	db, err := gorm.Open("mysql", sqlSource)
	if err !=nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	DB =db
}

