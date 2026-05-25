package main_test

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"testing"
)

var build bool

func init() {
	flag.BoolVar(&build, "build", true, "run build step")
}

func TestMain(m *testing.M) {
	flag.Parse()
	if build {
		cmd := exec.Command("go", "tool", "build")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	os.Exit(m.Run())
}

var WasmtimeAllowAllArgs = []string{
	"--dir", "/::/",
	"--wasi", "cli",
	"--wasi", "http",
	"--wasi", "tls",
	"--wasi", "inherit-network",
	"--wasi", "allow-ip-name-lookup",
	"--wasi", "tcp",
	"--wasi", "udp",
	"--wasi", "inherit-env",
	"--wasi", "inherit-stdin",
	"--wasi", "inherit-stdout",
	"--wasi", "inherit-stderr",
}

func WasmtimeInvoke(path string, expression string) (string, error) {
	cmd := exec.Command("wasmtime", append(slices.Clone(WasmtimeAllowAllArgs), "--invoke", expression, path)...)
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%v failed: %w\n%s", cmd, err, stderr)
	}
	if bytes.HasSuffix(out, []byte("\n")) {
		out = out[:len(out)-1]
		if bytes.HasSuffix(out, []byte("\r")) {
			out = out[:len(out)-1]
		}
	}
	return string(out), nil
}

func Wasmtime(path string, args ...string) *exec.Cmd {
	return exec.Command("wasmtime", append(append(slices.Clone(WasmtimeAllowAllArgs), path), args...)...)
}

func TestBuild(t *testing.T) {
	input := `build("v1.2.3-rc4+5.6.7")`
	output, err := WasmtimeInvoke("./mod.wasm", input)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v=%v", input, output)
		if got, want := output, `"+5.6.7"`; got != want {
			t.Errorf("%s: got=%v, want=%v", input, got, want)
		}
	}
}

func TestGosumcheck(t *testing.T) {
	cmd := Wasmtime("./gosumcheck.wasm", "--help")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := errors.AsType[*exec.ExitError](err); ok {
			if got, want := exitError.ExitCode(), 1; got != want {
				t.Errorf("got=%v, want=%v", got, want)
			}
			if got, want := out, []byte(`usage: gosumcheck [-h H] [-k key] [-u url] [-v] go.sum...`+"\n"); !bytes.Equal(got, want) {
				t.Errorf("got=%v, want=%v", got, want)
			}
		} else {
			t.Error(err)
		}
	} else {
		t.Error("expected exit code 1 with message")
	}
}
