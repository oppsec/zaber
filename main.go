package main

import (
	"log"
	"os"
	"github.com/oppsec/zaber/src/interface"
	"github.com/oppsec/zaber/src/zaber"
	"github.com/jessevdk/go-flags"
)

func error(err interface{}) {

	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
}

func main() {
	ui.GetBanner()

	var opts struct {
		Url string `short:"u" long:"url" description:"Definition: Argument used to pass target URL" required:"true"`
	}

	_, err := flags.Parse(&opts)
	error(err)

	exploit.TargetConnect(opts.Url)
}