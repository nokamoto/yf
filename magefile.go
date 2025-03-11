//go:build mage
// +build mage

package main

import (
	"log/slog"
	"os"

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
	mg.SerialDeps(Tidy, Imports, Fmt, Test)
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
