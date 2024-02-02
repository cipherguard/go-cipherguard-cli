package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/folder"
	"github.com/khulnasoft/go-cipherguard-cli/group"
	"github.com/khulnasoft/go-cipherguard-cli/resource"
	"github.com/khulnasoft/go-cipherguard-cli/user"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Deletes a Passbolt Entity",
	Long:    `Deletes a Passbolt Entity`,
	Aliases: []string{"remove"},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(resource.ResourceDeleteCmd)
	deleteCmd.AddCommand(folder.FolderDeleteCmd)
	deleteCmd.AddCommand(group.GroupDeleteCmd)
	deleteCmd.AddCommand(user.UserDeleteCmd)

	deleteCmd.PersistentFlags().String("id", "", "ID of the Entity to Delete")
}
