package main

import (
	"log"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
