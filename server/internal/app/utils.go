package app

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"photographer/internal/api"

	"github.com/sirupsen/logrus"
)

var (
	NUMBER_CHARSET     = "0123456789"
	LOWER_CASE_CHARSET = "abcdefghijklmnopqrstuvwxyz"
	UPPER_CASE_CHARSET = strings.ToUpper(LOWER_CASE_CHARSET)
)

// GetWD is just wrapper on os.Getwd w\out error
func GetWD() string {
	wd, _ := os.Getwd()
	return wd
}

// RndStr return random string from charset with certain length
// 	charset = "abcdef01234ABCDEF"
// 	length = 8
// 	return "aE42De0b"
func RndStr(charset string, length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// CreateLogger create log.Logger in /logs/logDir/logName_time.log
// 	logDir = where to place logs
// 	logName = logs prefix name
func CreateLogger(logDir, logName string) (*logrus.Logger, error) {
	wd := GetWD()

	if e := os.MkdirAll(wd+"/logs/"+logDir, os.ModePerm); e != nil {
		return nil, fmt.Errorf("create log error: %v", e.Error())
	}

	logFile, e := os.OpenFile(wd+"/logs/"+logDir+"/"+logName+"_"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		return nil, e
	}

	log := logrus.New()
	log.Out = logFile
	return log, nil
}

// RemoveExpiredFiles remove files in src folder which before d days before
// 	d = amount of days
// 	src = source folder
func RemoveExpiredFiles(d int, src string) error {
	before := time.Now().Add(time.Hour * -24 * time.Duration(d))

	files, e := os.ReadDir(src)
	if e != nil {
		return e
	}

	for _, f := range files {
		i, _ := f.Info()
		if i.ModTime().Before(before) {
			os.Remove(src + f.Name())
		}
	}
	return nil
}

// DoOnEachTick make tick on each timeOut duration and call funcs functions
// 	timeOut = make every timeOut time.Duration tick
// 	funcs = what shoud do on each tick
func DoOnEachTick(timeOut time.Duration, funcs ...func(tickCount *int)) {
	timer := time.NewTicker(timeOut)
	tickCount := 0
	for {
		<-timer.C
		timer.Reset(timeOut)
		tickCount++

		wg := sync.WaitGroup{}
		wg.Add(len(funcs))
		for _, f := range funcs {
			go f(&tickCount)
		}
		wg.Wait()
	}
}

// ZeroAppsRqCount clean apps request count
func (app *Application) ZeroAppsRqCount(tickCount *int) {
	app.CurrentRequestCount = 0
}

// MakeBackup make backup on every 30 tickCount
func (app *Application) MakeBackup(tickCount *int) {
	if *tickCount%30 == 0 {
		cmd := exec.Command("cp", `db/photographer.db`, `db/photographer_backup.db`)
		if e := cmd.Run(); e == nil {
			app.Log.Println("backup created!")
		} else {
			app.Log.Println(e)
		}
	}
}

// ClearSessions clear expired sessions
func (app *Application) ClearSessions(tickCount *int) {
	if e := api.SessionGC(); e != nil {
		app.Log.Println(e)
	}
}

// ClearCodes clear expired apps codes
func (app *Application) ClearCodes(tickCount *int) {
	for code, v := range app.UsersCode {
		if v.ExpireMin == app.CurrentMin {
			app.m.Lock()
			delete(app.UsersCode, code)
			app.m.Unlock()
		}
	}
}

// ZeroTickCount to zero tickCount
func ZeroTickCount(tickCount *int) {
	if *tickCount == 60*24 {
		*tickCount = 0
	}
}
