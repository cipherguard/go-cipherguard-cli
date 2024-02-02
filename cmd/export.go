package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/keepass"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports Passbolt Data",
	Long:  `Exports Passbolt Data`,
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.AddCommand(keepass.KeepassExportCmd)
}
