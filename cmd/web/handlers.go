package main

import (
	"net/http"
	"strconv"
)

func (app *application) encryptForm(w http.ResponseWriter, r *http.Request) {
	t := app.render(w, r, "./ui/html/encrypt.html")
	t.Execute(w, nil)
}

func (app *application) homeForm(w http.ResponseWriter, r *http.Request) {
	t := app.render(w, r, "./ui/html/index.html")
	t.Execute(w, nil)
}

func (app *application) decryptForm(w http.ResponseWriter, r *http.Request) {
	t := app.render(w, r, "./ui/html/decrypt.html")
	t.Execute(w, nil)
}

func (app *application) bruteforceForm(w http.ResponseWriter, r *http.Request) {
	t := app.render(w, r, "./ui/html/bruteforce.html")
	t.Execute(w, nil)
}

func (app *application) encrypt(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	plaintext := r.FormValue("plaintext")

	if key == "" || plaintext == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var Resp cryptTemplate
	Resp.Key = key
	Resp.Plaintext = plaintext
	Resp.Ciphertext = Encipher(plaintext, key)

	t := app.render(w, r, "./ui/html/encrypt.html")
	err := t.Execute(w, Resp)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) decrypt(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	ciphertext := r.FormValue("ciphertext")

	if key == "" || ciphertext == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var Resp cryptTemplate
	Resp.Key = key
	Resp.Ciphertext = ciphertext
	Resp.Plaintext = Decipher(ciphertext, key)

	t := app.render(w, r, "./ui/html/decrypt.html")
	err := t.Execute(w, Resp)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) bruteforce(w http.ResponseWriter, r *http.Request) {
	keyLength, err := strconv.Atoi(r.FormValue("keyLength"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	ciphertext := r.FormValue("ciphertext")

	if keyLength == 0 || ciphertext == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var Resp bruteforceTemplate
	Resp = BruteForce(ciphertext, keyLength)
	Resp.KeyLength = keyLength
	Resp.Ciphertext = ciphertext

	t := app.render(w, r, "./ui/html/bruteforce.html")
	err = t.Execute(w, Resp)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
