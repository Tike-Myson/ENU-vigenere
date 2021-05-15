package main

type encryptTemplate struct {
	message string
}

type decryptTemplate struct {
	message string
}

type bruteforceTemplate struct {
	result []resultBoard
}

type resultBoard struct {
	id int
	key string
	percent float32
}