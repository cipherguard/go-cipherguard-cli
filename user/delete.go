package user

import (
	"context"
	"fmt"

	"github.com/khulnasoft/go-cipherguard-cli/util"
	"github.com/khulnasoft/go-cipherguard/helper"
	"github.com/spf13/cobra"
)

// UserDeleteCmd Deletes a User
var UserDeleteCmd = &cobra.Command{
	Use:   "user",
	Short: "Deletes a Cipherguard User",
	Long:  `Deletes a Cipherguard User`,
	RunE:  UserDelete,
}

func UserDelete(cmd *cobra.Command, args []string) error {
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

	helper.DeleteUser(ctx, client, resourceID)
	if err != nil {
		return fmt.Errorf("Deleting User: %w", err)
	}
	return nil
}
