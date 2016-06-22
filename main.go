package main

import (
	"github.com/amyboyd/editorconfig-cli/editorconfig"
	"os"
)

func main() {
	app := editorconfig.CreateCliApp()

	app.Run(os.Args)
}
