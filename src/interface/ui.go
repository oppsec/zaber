package ui

import (
	"github.com/fatih/color"
)

func GetBanner() {

	banner_color := color.New(color.FgCyan).Add(color.Bold)

	banner_color.Println(`
	CVE-2019-9670 exploit
	version: 1.2
	`)

}