package sql

import(

	"fmt"
	"github.com/myxtype/filecoin-client/models"
	"github.com/myxtype/filecoin-client/pkg/setting"
	"log"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"


)

var db *gorm.DB

func InitSql()(db *gorm.DB)  {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
		AUTO_CREATE_TABLE bool
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	AUTO_CREATE_TABLE=sec.Key("AUTO_CREATE_TABLE").MustBool(false)

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	if AUTO_CREATE_TABLE {
		log.Println("-----Init DB----->" )
		var tables = []interface{}{
			&models.CmcTransation{},
			&models.Wallet{},
			&models.FilTransation{},

		}
		for _,table:=range tables{
			log.Printf("Automigrate table :%s ",reflect.TypeOf(table))
			if err:=db.AutoMigrate(table);err!=nil{
				log.Println("------创建表报错--->",reflect.TypeOf(table),err.Error)
			}
		}
	}
	defer db.Close()
	return
}



