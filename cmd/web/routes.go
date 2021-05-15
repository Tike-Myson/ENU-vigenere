package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/encrypt", standardMiddleware.ThenFunc(app.encryptForm))
	mux.Get("/decrypt", standardMiddleware.ThenFunc(app.decryptForm))
	mux.Get("/bruteforce", standardMiddleware.ThenFunc(app.bruteforceForm))
	mux.Post("/encrypt", standardMiddleware.ThenFunc(app.encrypt))
	mux.Post("/decrypt", standardMiddleware.ThenFunc(app.decrypt))
	mux.Post("/decrypt", standardMiddleware.ThenFunc(app.bruteforce))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
