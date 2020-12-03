package manager

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myxtype/filecoin-client/models"
	"github.com/myxtype/filecoin-client/pkg/setting"
	"log"
	"reflect"
)

var (
	dbinstance  *gorm.DB
)


func GetDbInstance() (instance *gorm.DB) {
	if dbinstance == nil {
		db, err :=initDb()
		if err != nil {
			log.Fatalln(err)
		}
		dbinstance = db
		return dbinstance
	}
	fmt.Printf("----db instance &----- %p", &dbinstance)
	return dbinstance
}



func initDb() (*gorm.DB,error) {
	sec, _err := setting.Cfg.GetSection("database")
	if _err != nil {
		return nil,_err
	}
	DBTYPE := sec.Key("TYPE").String()
	DBNAME := sec.Key("NAME").String()
	USER := sec.Key("USER").String()
	PASSWORD := sec.Key("PASSWORD").String()
	HOST := sec.Key("HOST").String()
	TABLEPREFIX := sec.Key("TABLE_PREFIX").String()
	AUTO_CREATE_TABLE:=sec.Key("AUTO_CREATE_TABLE").MustBool(false)

	gdb, err:= gorm.Open(DBTYPE, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER,
		PASSWORD,
		HOST,
		DBNAME))
	if err != nil {
		return nil,err
	}

	// 启用Logger，显示详细日志
	gdb.LogMode(true)
	gdb.SingularTable(false)
	gdb.DB().SetMaxIdleConns(10)
	gdb.DB().SetMaxOpenConns(50)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return TABLEPREFIX + defaultTableName
	}

	if AUTO_CREATE_TABLE {
		var tables = []interface{}{
			//&models.CmcTransation{},
			&models.Wallet{},
			//&models.CmcRecord{},
			&models.FilTransation{},
			&models.Auth{},
			&models.FilIntegral{},

		}
		for _, table := range tables {
			log.Printf("Automigrate table :%s ",reflect.TypeOf(table))
			if err = gdb.AutoMigrate(table).Error; err != nil {
				return nil,err
			}
		}
	}

	return gdb,nil
}
