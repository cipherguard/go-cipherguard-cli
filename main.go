package main

import (
	"github.com/khulnasoft/go-cipherguard-cli/cmd"
)

func main() {
	cmd.SetVersionInfo(version, commit, date, dirty)
	cmd.Execute()
}
