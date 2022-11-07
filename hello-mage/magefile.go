//go:build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

func Exec(os string, arch string, projectBuildDir string, projectName string) error {
	goos := fmt.Sprintf("GOOS=%s", os)
	goarch := fmt.Sprintf("GOARCH=%s", arch)
	outputLocation := fmt.Sprintf("%s/%s-%s/%s", projectBuildDir, os, arch, projectName)
	fmt.Println(outputLocation)
	return sh.Run("env", goos, goarch, "CGO_ENABLED=0", "go", "build", "-o", outputLocation)
}
