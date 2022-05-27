package main

import (
	"os"

	"github.com/aveplen/silicon-funnel/internal/bot/config"
	"github.com/aveplen/silicon-funnel/internal/bot/server"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	ConfigPath *string `short:"c" long:"config" description:"Path to yaml config"`
}

func init() {
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		panic(err)
	}

	if opts.ConfigPath == nil {
		opts.ConfigPath = &args[0]
	}
}

func main() {
	cfg, err := config.ReadConfig(*opts.ConfigPath)
	if err != nil {
		panic(err)
	}

	server.Start(cfg)
}
