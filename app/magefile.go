//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Runs go mod download and then installs the binary.
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	args := []string{
		"build",
		"-o",
		"../webview-example.exe",
		`-ldflags`,
		"-H windowsgui",
	}
	return sh.Run("go", args...)
}
