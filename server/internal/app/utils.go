package app

import (
	"log"
	"math/rand"
	"os"
	"photographer/internal/api"
	"time"
	// "alber/pkg/api"
)

const (
	DBPort      = 8000
	ApiPort     = 8010
	AuthPort    = 8020
	BillingPort = 8030
)

func RandomStringFromCharsetAndLength(charset string, length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func CreateLogged(logType string) *log.Logger {
	wd, _ := os.Getwd()
	logFile, _ := os.OpenFile(wd+"/logs/"+logType+"/log_"+time.Now().Format("2006-01-02")+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	return log.New(logFile, "\033[31m[ERROR]\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)
}

// DoBackup make backup every 30 min
func (app *Application) DoBackup() error {
	// cmd := exec.Command("cp", `db/alber.db`, `db/alber_backup.db`)
	// return cmd.Run()
}

// CheckPerMin call SessionGC per minute that delete expired sessions and do db backup
func (app *Application) CheckPerMin() {
	timer := time.NewTicker(1 * time.Minute)
	for {
		// manage timer
		<-timer.C
		timer.Reset(1 * time.Minute)

		// change conf app
		app.CurrentRequestCount = 0
		app.CurrentMin++

		// do general actions
		if app.CurrentMin == 60*24 {
			app.CurrentMin = 0
		}
		if app.CurrentMin%30 == 0 {
			if e := app.DoBackup(); e == nil {
				app.Log.Println("backup created!")
			} else {
				app.Log.Println(e)
			}
		}
		if e := api.SessionGC(); e != nil {
			app.Log.Println(e)
		}

		// remove expired codes
		go func() {
			for code, v := range app.UsersCode {
				if v.ExpireMin == app.CurrentMin {
					app.m.Lock()
					delete(app.UsersCode, code)
					app.m.Unlock()
				}
			}
		}()
	}
}
