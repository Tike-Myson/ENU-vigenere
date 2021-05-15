package main

import "net/http"

func (app *application) encryptForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "./ui/html/encrypt.html")
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (app *application) decryptForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "./ui/html/decrypt.html")
}

func (app *application) bruteforceForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "./ui/html/bruteforce.html")
}

func (app *application) encrypt(w http.ResponseWriter, r *http.Request) {

}

func (app *application) decrypt(w http.ResponseWriter, r *http.Request) {

}

func (app *application) bruteforce(w http.ResponseWriter, r *http.Request) {

}
