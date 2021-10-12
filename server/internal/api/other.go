package api

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"photographer/internal/orm"
)

func Images(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method == "POST" {
		return nil, errors.New("wrong method")
	}

	ID, e := strconv.Atoi(r.FormValue("id"))
	if e != nil {
		return nil, errors.New("не корректный id")
	}

	mainQ := orm.SQLSelectParams{
		Table:   "Images AS i",
		What:    "i.*",
		Options: orm.DoSQLOption("i.parselID=?", "", "", ID),
	}
	if datas := orm.GeneralGet(mainQ, nil, orm.Photo{}); datas != nil {
		return datas, nil
	}
	return nil, errors.New("н/д")
}

// CreateImage create one image
func CreateImage(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	userID := GetUserIDfromReq(w, r)
	if userID == -1 {
		return nil, errors.New("не зарегистрированы в сети")
	}

	// link, name := r.PostFormValue("link"), r.PostFormValue("filename")
	// i := &orm.Photo{
	// 	Source: link, Name: name,
	// 	CustID: uint(userID),
	// }

	// if _, e = i.Create(); e != nil {
	// 	wd, _ := os.Getwd()
	// 	os.Remove(wd + i.Source)
	// 	return nil, errors.New("не удалось прикрепить фото")
	// }
	return nil, nil
}

// RemoveImage remove one image
func RemoveImage(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// get general ids
	userID := GetUserIDfromReq(w, r)
	if userID == -1 {
		return nil, errors.New("не зарегистрированы в сети")
	}
	imgID, e := strconv.Atoi(r.PostFormValue("id"))
	if e != nil {
		return nil, errors.New("не корректный id  фото")
	}

	wd, _ := os.Getwd()
	os.Remove(wd + r.PostFormValue("src"))

	return nil, orm.DeleteByParams(orm.SQLDeleteParams{
		Table:   "Images",
		Options: orm.DoSQLOption("id=? AND userID=?", "", "", imgID, userID),
	})
}
