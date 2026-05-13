package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
)

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		for _, p := range flag.Args() {
			slog.Info("remove all", "path", p)
			err := os.RemoveAll(p)
			if err != nil {
				slog.Warn(err.Error())
			}
		}
	} else {
		log.Fatal("no arguments")
	}
}
