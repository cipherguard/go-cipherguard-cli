package user

import (
	"context"
	"fmt"

	"github.com/khulnasoft/go-cipherguard-cli/util"
	"github.com/khulnasoft/go-cipherguard/helper"
	"github.com/spf13/cobra"
)

// UserUpdateCmd Updates a Cipherguard User
var UserUpdateCmd = &cobra.Command{
	Use:   "user",
	Short: "Updates a Cipherguard User",
	Long:  `Updates a Cipherguard User`,
	RunE:  UserUpdate,
}

func init() {
	UserUpdateCmd.Flags().String("id", "", "id of User to Update")
	UserUpdateCmd.Flags().StringP("firstname", "f", "", "User FirstName")
	UserUpdateCmd.Flags().StringP("lastname", "l", "", "User LastName")
	UserUpdateCmd.Flags().StringP("role", "r", "", "User Role")

	UserUpdateCmd.MarkFlagRequired("id")
}

func UserUpdate(cmd *cobra.Command, args []string) error {
	id, err := cmd.Flags().GetString("id")
	if err != nil {
		return err
	}
	firstname, err := cmd.Flags().GetString("firstname")
	if err != nil {
		return err
	}
	lastname, err := cmd.Flags().GetString("lastname")
	if err != nil {
		return err
	}
	role, err := cmd.Flags().GetString("role")
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

	err = helper.UpdateUser(
		ctx,
		client,
		id,
		role,
		firstname,
		lastname,
	)
	if err != nil {
		return fmt.Errorf("Updating User: %w", err)
	}
	return nil
}
