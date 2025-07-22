package main

import (
	"log"

	"github.com/shadowdara/checkstaticfiles/core"
)

func main() {
	CheckFiles()
}

func CheckFiles() {
	log.Println("Checking required files...")
	core.Main()
	log.Println("Checking finished successfully!")
}

func CreateFiles() {
	log.Println("Creating required files...")
	log.Println("Creating finished successfully!")
}
