package group

import (
	"context"
	"fmt"

	"github.com/khulnasoft/go-cipherguard-cli/util"
	"github.com/spf13/cobra"
)

// GroupDeleteCmd Deletes a Group
var GroupDeleteCmd = &cobra.Command{
	Use:   "group",
	Short: "Deletes a Cipherguard Group",
	Long:  `Deletes a Cipherguard Group`,
	RunE:  GroupDelete,
}

func GroupDelete(cmd *cobra.Command, args []string) error {
	resourceID, err := cmd.Flags().GetString("id")
	if err != nil {
		return err
	}

	if resourceID == "" {
		return fmt.Errorf("No ID to Delete Provided")
	}

	ctx := util.GetContext()

	client, err := util.GetClient(ctx)
	if err != nil {
		return err
	}
	defer client.Logout(context.TODO())
	cmd.SilenceUsage = true

	client.DeleteGroup(ctx, resourceID)
	if err != nil {
		return fmt.Errorf("Deleting Group: %w", err)
	}
	return nil
}
