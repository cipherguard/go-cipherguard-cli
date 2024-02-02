package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/keepass"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports Cipherguard Data",
	Long:  `Exports Cipherguard Data`,
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.AddCommand(keepass.KeepassExportCmd)
}
