package main

import (
	"os"

	"github.com/taylormonacelli/dentwatcher"
)

func main() {
	code := dentwatcher.Execute()
	os.Exit(code)
}
