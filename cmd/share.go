package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/folder"
	"github.com/khulnasoft/go-cipherguard-cli/resource"
	"github.com/spf13/cobra"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Shares a Passbolt Entity",
	Long:  `Shares a Passbolt Entity`,
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.AddCommand(resource.ResourceShareCmd)
	shareCmd.AddCommand(folder.FolderShareCmd)
}
