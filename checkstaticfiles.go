package shadowdara_checkstaticfiles

import (
	"log"

	"github.com/shadowdara/checkstaticfiles/core"
)

func Checkfiles(data []byte, settings int) {
	log.Println("Checking required files...")
	core.Main(data, settings)
}
