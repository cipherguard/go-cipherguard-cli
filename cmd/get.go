package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/folder"
	"github.com/khulnasoft/go-cipherguard-cli/group"
	"github.com/khulnasoft/go-cipherguard-cli/resource"
	"github.com/khulnasoft/go-cipherguard-cli/user"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets a Passbolt Entity",
	Long:    `Gets a Passbolt Entity`,
	Aliases: []string{"read"},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().BoolP("json", "j", false, "Output JSON")
	getCmd.AddCommand(resource.ResourceGetCmd)
	getCmd.AddCommand(folder.FolderGetCmd)
	getCmd.AddCommand(group.GroupGetCmd)
	getCmd.AddCommand(user.UserGetCmd)

}
