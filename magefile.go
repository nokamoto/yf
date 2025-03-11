//go:build mage
// +build mage

package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = All

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))
}

// All runs all the tasks.
func All() {
	mg.SerialDeps(Tidy, Imports, Fmt, Test, Install, Examples)
}

// Tidy runs go mod tidy.
func Tidy() error {
	slog.Info("Running go mod tidy")
	return sh.RunV("go", "mod", "tidy")
}

// Test runs go test.
func Test() error {
	slog.Info("Running go test")
	return sh.RunV("go", "test", "./...")
}

// Fmt runs gofumpt.
func Fmt() error {
	slog.Info("Running gofumpt")
	if err := sh.RunV("go", "install", "mvdan.cc/gofumpt@latest"); err != nil {
		return err
	}
	return sh.RunV("gofumpt", "-w", ".")
}

// Imports runs goimports.
func Imports() error {
	slog.Info("Running goimports")
	if err := sh.RunV("go", "install", "golang.org/x/tools/cmd/goimports@latest"); err != nil {
		return err
	}
	return sh.RunV("goimports", "-w", ".")
}

// Install runs go install.
func Install() error {
	return sh.RunV("go", "install", "./cmd/yf")
}

// Examples updates the examples directory.
func Examples() error {
	slog.Info("Updating examples")
	const dir = "examples"
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Dir(path) != dir {
			return nil
		}

		out := filepath.Join(dir, "out", filepath.Base(path))
		slog.Info("yf", "path", path, "out", out)
		in, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read file: %w", err)
		}

		var stdout bytes.Buffer
		stdin := bytes.NewBuffer(in)
		yf := exec.Command("yf")
		yf.Stdin = stdin
		yf.Stderr = os.Stderr
		yf.Stdout = &stdout
		if err := yf.Run(); err != nil {
			return fmt.Errorf("run yf: %w", err)
		}
		if err := os.WriteFile(out, stdout.Bytes(), 0o644); err != nil {
			return fmt.Errorf("write file: %w", err)
		}
		return nil
	})
	return err
}
