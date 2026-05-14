package main

import (
	"flag"
	"log"
	"log/slog"
	"maps"
	"os/exec"
	"slices"
	"sync/atomic"
)

var TargetArg string

func Parse() {
	flag.Parse()
	TargetArg = flag.Arg(0)
}

var Targets = map[string]func() error{
	"mod":        BuildMod,
	"gosumcheck": BuildGosumcheck,
}

var didWkgWitFetch atomic.Bool

func BuildMod() error {
	if didWkgWitFetch.CompareAndSwap(false, true) {
		cmd := exec.Command("wkg", "wit", "fetch")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		slog.Info("run", "cmd", cmd)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	cmd := exec.Command("tinygo", "build", "-target", "wasip2", "-wit-package", "./wit/", "-wit-world", "gosumcheck", "golang.org/x/mod/gosumcheck")
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	slog.Info("run", "cmd", cmd)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func BuildGosumcheck() error {
	if didWkgWitFetch.CompareAndSwap(false, true) {
		cmd := exec.Command("wkg", "wit", "fetch")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		slog.Info("run", "cmd", cmd)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	cmd := exec.Command("tinygo", "build", "-target", "wasip2", "-wit-package", "./wit/", "-wit-world", "mod", "-o", "mod.wasm", ".")
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	slog.Info("run", "cmd", cmd)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	Parse()
	if TargetArg == "" {
		for _, build := range Targets {
			err := build()
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		build, ok := Targets[TargetArg]
		if !ok {
			log.Fatalf("%q not in %v", TargetArg, slices.Collect(maps.Keys(Targets)))
		}
		err := build()
		if err != nil {
			log.Fatal(err)
		}
	}
}
