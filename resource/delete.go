package resource

import (
	"context"
	"fmt"

	"github.com/khulnasoft/go-cipherguard-cli/util"
	"github.com/spf13/cobra"
)

// ResourceDeleteCmd Deletes a Resource
var ResourceDeleteCmd = &cobra.Command{
	Use:   "resource",
	Short: "Deletes a Cipherguard Resource",
	Long:  `Deletes a Cipherguard Resource`,
	RunE:  ResourceDelete,
}

func ResourceDelete(cmd *cobra.Command, args []string) error {
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

	client.DeleteResource(ctx, resourceID)
	if err != nil {
		return fmt.Errorf("Deleting Resource: %w", err)
	}
	return nil
}
