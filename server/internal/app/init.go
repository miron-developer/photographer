package app

import (
	"log"
	"os"
	"sync"
	"time"
	// "photographer/internal/orm"
)

// Code struct for app
type Code struct {
	ExpireMin int
	Value     interface{}
}

// AppConfig - app's configurations
type AppConfig struct {
	PORT              string
	MOBIZON_API_KEY   string
	MAX_REQUEST_COUNT string
}

// Application - app config and items
type Application struct {
	m                   sync.Mutex
	Log                 *log.Logger
	CurrentRequestCount int
	CurrentMin          int // how many minuts pass after start/day
	UsersCode           map[string]*Code
	Config              *AppConfig
}

const (
	DBPort      = 8000
	ApiPort     = 8010
	AuthPort    = 8020
	BillingPort = 8030
)

func CheckFatal(eLogger *log.Logger, e error) {
	if e != nil {
		eLogger.Fatal(e)
	}
}

// InitProg initialise
func InitProg() *Application {
	wd, _ := os.Getwd()
	logFile, _ := os.OpenFile(wd+"/logs/log_"+time.Now().Format("2006-01-02")+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	log := log.New(logFile, "\033[31m[ERROR]\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)
	// iLog := log.New(logFile, "\033[34m[INFO]\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("loggers is done!")

	log.Println("creating/configuring database")
	// CheckFatal(log, orm.InitDB(log))
	log.Println("database completed!")

	log.Println("configuring app")
	// config, e := GetConfigs()
	// CheckFatal(log, e)
	log.Println("configuring done")

	return &Application{
		Log:                 log,
		CurrentRequestCount: 0,
		CurrentMin:          0,
		UsersCode:           map[string]*Code{},
		// Config:              config,
	}
}
