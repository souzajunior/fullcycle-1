package main

import (
	"log"

	"simulator-fc1/config"
)

func init() {
	config.LoadConfig()

	log.SetFlags(log.Lshortfile)
}

func main() {

}
