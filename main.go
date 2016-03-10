package main

import (
	"flag"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var (
	password string
	cost     int
)

func init() {
	flag.StringVar(&password, "p", "", "password to bcrypt")
	flag.IntVar(&cost, "c", 10, "cost to pass to bcrypt")
}

func main() {
	flag.Parse()
	if password == "" {
		log.Fatalf("[bcryptify] must provide a password")
	}
	b := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(b, cost)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	log.Printf("[bcryptify] generated hash\n%s\n", hash)
}
