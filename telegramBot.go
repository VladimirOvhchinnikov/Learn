package main

import (
	"flag"
	"log"
)

func MustToken() string {
	token := flag.String("token", "", "token for tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal()
	}

	return *token
}
