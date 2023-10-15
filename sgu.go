package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.MkdirAll("", os.ModeDir))
}
