package setting

import (
	"log"
	"time"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	AUTO_CREATE_TABLE bool

	PageSize int
	JwtSecret string
)

//rpc
var(
	SERVER_HOST string
	SERVER_PORT int
	ISMINER bool
	TOKEN string
	USER        string
	PASSWD      string
	USESSL      bool
)
//cron
var(
	CRONRULE string
	CRONRULE1 string
	CRONRULE_TRANS string
	CRONRULE_CONFIRM string
)



func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	loadRpc()
	loadCron()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

func loadRpc(){
	sec, err := Cfg.GetSection("rpc")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	SERVER_HOST = sec.Key("SERVER_HOST").MustString("")
	SERVER_PORT = sec.Key("SERVER_PORT").MustInt(9888)
	TOKEN = sec.Key("TOKEN").MustString("")
	USER = sec.Key("USER").MustString("")
	PASSWD = sec.Key("PASSWD").MustString("")
	USESSL = sec.Key("USESSL").MustBool(false)
	ISMINER = sec.Key("ISMINER").MustBool(true)
}

func loadCron()  {
	sec, err := Cfg.GetSection("cron")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	CRONRULE = sec.Key("CRONRULE").MustString("0 0/5 * * * ?")
	CRONRULE_TRANS = sec.Key("CRONRULE_TRANS").MustString("0 0/3 * * * ?")
	CRONRULE_CONFIRM = sec.Key("CRONRULE_CONFIRM").MustString("0 0/2 * * * ?")
}


func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}