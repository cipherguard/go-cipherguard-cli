package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/folder"
	"github.com/khulnasoft/go-cipherguard-cli/group"
	"github.com/khulnasoft/go-cipherguard-cli/resource"
	"github.com/khulnasoft/go-cipherguard-cli/user"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates a Cipherguard Entity",
	Long:    `Creates a Cipherguard Entity`,
	Aliases: []string{"new"},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().BoolP("json", "j", false, "Output JSON")
	createCmd.AddCommand(resource.ResourceCreateCmd)
	createCmd.AddCommand(folder.FolderCreateCmd)
	createCmd.AddCommand(group.GroupCreateCmd)
	createCmd.AddCommand(user.UserCreateCmd)
}
