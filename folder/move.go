package folder

import (
	"context"
	"fmt"

	"github.com/khulnasoft/go-cipherguard-cli/util"
	"github.com/khulnasoft/go-cipherguard/helper"
	"github.com/spf13/cobra"
)

// FolderMoveCmd Moves a Cipherguard Folder
var FolderMoveCmd = &cobra.Command{
	Use:   "folder",
	Short: "Moves a Cipherguard Folder into a Folder",
	Long:  `Moves a Cipherguard Folder into a Folder`,
	RunE:  FolderMove,
}

func init() {
	FolderMoveCmd.Flags().String("id", "", "id of Folder to Move")
	FolderMoveCmd.Flags().StringP("folderParentID", "f", "", "Folder in which to Move the Folder")

	FolderMoveCmd.MarkFlagRequired("id")
	FolderMoveCmd.MarkFlagRequired("folderParentID")
}

func FolderMove(cmd *cobra.Command, args []string) error {
	id, err := cmd.Flags().GetString("id")
	if err != nil {
		return err
	}
	folderParentID, err := cmd.Flags().GetString("folderParentID")
	if err != nil {
		return err
	}

	ctx := util.GetContext()

	client, err := util.GetClient(ctx)
	if err != nil {
		return err
	}
	defer client.Logout(context.TODO())
	cmd.SilenceUsage = true

	err = helper.MoveFolder(
		ctx,
		client,
		id,
		folderParentID,
	)
	if err != nil {
		return fmt.Errorf("Moving Folder: %w", err)
	}
	return nil
}
