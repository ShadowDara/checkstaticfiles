package main

import (
	"log"

	"github.com/shadowdara/checkstaticfiles/core"
)

func main() {
	Checkfiles(CheckstaticfilesOutputJSONGz)
}

func Checkfiles(data []byte) {
	log.Println("Checking required files...")
	core.Main(data)
}
