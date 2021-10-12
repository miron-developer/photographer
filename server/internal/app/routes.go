package app

import "net/http"

func (app *Application) SetRoutes() http.Handler {
	appMux := http.NewServeMux()
	appMux.HandleFunc("/", app.HIndex)

	// sign
	signMux := http.NewServeMux()
	signMux.HandleFunc("/", app.HIndex)
	signMux.HandleFunc("/up", app.HSignUp)
	signMux.HandleFunc("/in", app.HSignIn)
	signMux.HandleFunc("/re", app.HResetPassword)
	signMux.HandleFunc("/out", app.HLogout)
	signMux.HandleFunc("/status", app.HCheckUserLogged)
	appMux.Handle("/sign/", http.StripPrefix("/sign", signMux))

	// api routes
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/", app.HApiIndex)
	apiMux.HandleFunc("/user", app.HUser)
	apiMux.HandleFunc("/users", app.HUsers)
	apiMux.HandleFunc("/search", app.HSearch)
	apiMux.HandleFunc("/images", app.HClippedImages)
	appMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	// save
	saveMux := http.NewServeMux()
	saveMux.HandleFunc("/", app.HApiIndex)
	appMux.Handle("/s/", http.StripPrefix("/s", saveMux))

	// edit
	editMux := http.NewServeMux()
	editMux.HandleFunc("/", app.HApiIndex)
	editMux.HandleFunc("/user", app.HChangeProfile)
	appMux.Handle("/e/", http.StripPrefix("/e", editMux))

	// remove
	removeMux := http.NewServeMux()
	removeMux.HandleFunc("/", app.HApiIndex)
	removeMux.HandleFunc("/image", app.HRemoveImage)
	appMux.Handle("/r/", http.StripPrefix("/r", removeMux))

	// static react get
	static := http.FileServer(http.Dir("website/static"))
	appMux.Handle("/static/", http.StripPrefix("/static/", static))

	// assets get
	assets := http.FileServer(http.Dir("assets"))
	appMux.Handle("/assets/", http.StripPrefix("/assets/", assets))

	// middlewares
	muxHanlder := app.AccessLogMiddleware(appMux)
	muxHanlder = app.SecureHeaderMiddleware(muxHanlder)
	return muxHanlder
}
