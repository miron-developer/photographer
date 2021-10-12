package app

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"photographer/internal/api"
)

// SecureHeaderMiddleware set secure header option
func (app *Application) SecureHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1;mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		next.ServeHTTP(w, r)
	})
}

// AccessLogMiddleware logging request
func (app *Application) AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		max, _ := strconv.Atoi(app.Config.MAX_REQUEST_COUNT)
		if app.CurrentRequestCount < max {
			app.CurrentRequestCount++
			// loging
			// app.Log.Printf(logingReq(r))
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "service is overloaded", 529)
			app.Log.Println(errors.New("rate < curl"))
		}
	})
}

// Hindex for handle '/'
func (app *Application) HIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		wd, _ := os.Getwd()
		t, e := template.ParseFiles(wd + "/website/index.html")
		if e != nil {
			http.Error(w, "can't load this page", 500)
			app.Log.Println(e)
			return
		}
		t.Execute(w, nil)
	}
}

/* ------------------------------------------- API ------------------------------------------------ */

// HUser for handle '/api/'
func (app *Application) HApiIndex(w http.ResponseWriter, r *http.Request) {
	type Route struct {
		Path        string              `json:"route"`
		Description string              `json:"description"`
		Children    []Route             `json:"children"`
		Params      []map[string]string `json:"params"`
		Methods     []string            `json:"methods"`
	}

	data := api.API_RESPONSE{
		Err:  "",
		Code: 200,
		Data: []Route{
			{Path: "/user", Description: "получение одного пользователя(полное)", Methods: []string{"GET"}, Params: []map[string]string{{"k": "id", "v": "id пользователя"}}},
			{
				Path: "/users", Description: "получение id пользователей", Methods: []string{"GET"},
				Params: []map[string]string{{"k": "from", "v": "с какого индекса"}, {"k": "step", "v": "сколько взять"}},
			},
			{
				Path: "/parsels", Description: "посылки", Methods: []string{"GET"},
				Params: []map[string]string{
					{"k": "type", "v": "user/empty"}, {"k": "weight", "v": "масса"},
					{"k": "price", "v": "цена"}, {"k": "fromID", "v": "откуда"}, {"k": "toID", "v": "куда"},
					{"k": "from", "v": "с какого индекса"}, {"k": "step", "v": "сколько взять"},
				},
			},
			{
				Path: "/travelers", Description: "попутчики", Methods: []string{"GET"},
				Params: []map[string]string{
					{"k": "type", "v": "user(если твое)/empty(общий)"}, {"k": "fromID", "v": "откуда"}, {"k": "toID", "v": "куда"},
					{"k": "from", "v": "с какого индекса"}, {"k": "step", "v": "сколько взять"},
				},
			},
			{Path: "/images", Description: "прикрепленные фото", Methods: []string{"GET"}, Params: []map[string]string{{"k": "id", "v": "parsel id"}}},
			{Path: "/search", Description: "поиск(пока ищет города с таким названием)", Methods: []string{"GET"}, Params: []map[string]string{{"k": "q", "v": "поисковый запрос"}}},
			{Path: "/toptypes", Description: "поднятия", Methods: []string{"GET"}},
			{Path: "/travelTypes", Description: "типы путешествия", Methods: []string{"GET"}},
			{Path: "/countryCodes", Description: "коды стран", Methods: []string{"GET"}},
		},
	}

	api.DoJS(w, data)
}

// HUser for handle '/api/user/'
func (app *Application) HUser(w http.ResponseWriter, r *http.Request) {
	api.HApi(w, r, api.User)
}

// HUser for handle '/api/users/'
func (app *Application) HUsers(w http.ResponseWriter, r *http.Request) {
	api.HApi(w, r, api.Users)
}

// HSearch for handle '/api/search'
func (app *Application) HSearch(w http.ResponseWriter, r *http.Request) {
	api.HApi(w, r, api.Search)
}

// HClippedFiles for handle '/api/images'
func (app *Application) HClippedImages(w http.ResponseWriter, r *http.Request) {
	api.HApi(w, r, api.Images)
}

/* --------------------------------------------- Logical ---------------------------------- */
// ---------------------------------------------- Sign ---------------------------------------

// HcheckUserLogged for handle '/status'
func (app *Application) HCheckUserLogged(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := api.API_RESPONSE{
			Err:  "ok",
			Data: "",
			Code: 200,
		}

		userID := api.GetUserIDfromReq(w, r)
		if userID == -1 {
			api.SendErrorJSON(w, data, "нет зарегистрированы в сети")
			return
		}

		data.Data = map[string]int{"id": userID}
		api.DoJS(w, data)
	}
}

// HSignUp for handle '/sign/up'
func (app *Application) HSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := api.API_RESPONSE{
			Err:  "ok",
			Data: "",
			Code: 200,
		}

		successData, e := app.SignUp(w, r)
		if e != nil {
			api.SendErrorJSON(w, data, e.Error())
			return
		}
		data.Data = successData

		// delete unnecessary code
		app.m.Lock()
		delete(app.UsersCode, r.PostFormValue("code"))
		app.m.Unlock()

		api.DoJS(w, data)
	}
}

// HSignIn for handle '/sign/in'
func (app *Application) HSignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := api.API_RESPONSE{
			Err:  "ok",
			Data: "",
			Code: 200,
		}

		ID, e := app.SignIn(w, r)
		if e != nil {
			api.SendErrorJSON(w, data, e.Error())
			return
		}
		data.Data = map[string]int{"id": ID}
		api.DoJS(w, data)
	}
}

// HRestore for handle '/sign/re'
func (app *Application) HResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := api.API_RESPONSE{
			Err:  "ok",
			Data: "",
			Code: 200,
		}

		if e := app.ResetPassword(w, r); e != nil {
			api.SendErrorJSON(w, data, e.Error())
			return
		}

		// delete unnecessary code
		app.m.Lock()
		delete(app.UsersCode, r.PostFormValue("code"))
		app.m.Unlock()

		api.DoJS(w, data)
	}
}

// HLogout for handle '/sign/out'
func (app *Application) HLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := api.API_RESPONSE{
			Err:  "ok",
			Data: "",
			Code: 200,
		}

		userID := api.GetUserIDfromReq(w, r)
		if userID == -1 {
			api.SendErrorJSON(w, data, "не зарегистрированы в сети")
			return
		}

		if e := app.Logout(w, r); e != nil {
			api.SendErrorJSON(w, data, e.Error())
		}
		api.DoJS(w, data)
	}
}

// ------------------------------------------- Change ------------------------------------------

// HChangeProfile user data
func (app *Application) HChangeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data := api.API_RESPONSE{
			Err:  "ok",
			Data: "",
			Code: 200,
		}

		if _, ok := app.UsersCode[r.PostFormValue("code")]; !ok {
			api.SendErrorJSON(w, data, "wrong code")
			return
		}

		if e := api.ChangeProfile(w, r); e != nil {
			api.SendErrorJSON(w, data, e.Error())
			return
		}

		// delete unnecessary code
		app.m.Lock()
		delete(app.UsersCode, r.PostFormValue("code"))
		app.m.Unlock()

		api.DoJS(w, data)
	}
}

// ------------------------------------------- Save ------------------------------------------

// ------------------------------------------- Remove ----------------------------------------

// HRemoveImage save one image
func (app *Application) HRemoveImage(w http.ResponseWriter, r *http.Request) {
	api.HApi(w, r, api.RemoveImage)
}
