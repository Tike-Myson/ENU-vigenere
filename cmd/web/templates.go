package main

type cryptTemplate struct {
	Key string
	Plaintext string
	Ciphertext string
}

type bruteforceTemplate struct {
	KeyLength int
	Ciphertext string
	Result []resultBoard
}

type resultBoard struct {
	Id int
	Key string
	Percent string
}