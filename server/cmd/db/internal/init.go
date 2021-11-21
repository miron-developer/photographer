package internal

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"photographer/internal/app"
	"photographer/internal/consul"
)

// DB_SERVICE service's service
type DB_SERVICE struct {
	DBConn     *gorm.DB             // conn to db
	Log        *logrus.Logger       // logrus logger
	Config     *Config              // configs
	Client     *consul.ConsulClient // consul client
	Prommetric *DBPrommetric        // prometheus prometrics
}

// Init init service, settings and tables
func Init() *DB_SERVICE {
	var (
		e       error
		service = &DB_SERVICE{} // service ctx
	)

	// init logger
	if service.Log, e = app.CreateLogger("db", "log"); e != nil {
		log.Fatal(e)
	}

	// prometrics
	service.NewDBPrommetric()

	// configs
	service.NewConfig()

	// consul client
	if service.Client, e = consul.NewClient(service.Config.CONSUL_ADDR); e != nil {
		service.Log.Fatalln("consul access error: ", e)
	}

	// sync with consul config
	service.SyncWithConsul()

	// establish connection to service
	service.OpenConnectionToDB()

	// do migrations
	if e := service.makeMigrations(); e != nil {
		service.Log.Fatalln("migration error: ", e)
	}
	return service
}

// OpenConnectionToDB open connection to postgres
func (service *DB_SERVICE) OpenConnectionToDB() {
	var e error
	service.Log.Infoln("accessing database...")
	dsn := fmt.Sprintf(
		"host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
		service.Config.DB_HOST,
		service.Config.DB_PORT,
		service.Config.DB_NAME,
		service.Config.DB_USER,
		service.Config.DB_PASSWORD,
	)
	if service.DBConn, e = gorm.Open(postgres.Open(dsn), &gorm.Config{}); e != nil {
		service.Log.Fatalln(e)
	}
	service.Log.Infoln("done!")
}
