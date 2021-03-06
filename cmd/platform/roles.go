// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.
package main

import (
	"errors"

	"github.com/mattermost/platform/app"
	"github.com/spf13/cobra"
)

var rolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "Management of user roles",
}

var makeSystemAdminCmd = &cobra.Command{
	Use:     "system_admin [users]",
	Short:   "Set a user as system admin",
	Long:    "Make some users system admins",
	Example: "  roles system_admin user1",
	RunE:    makeSystemAdminCmdF,
}

var makeMemberCmd = &cobra.Command{
	Use:     "member [users]",
	Short:   "Remove system admin privileges",
	Long:    "Remove system admin privileges from some users.",
	Example: "  roles member user1",
	RunE:    makeMemberCmdF,
}

func init() {
	rolesCmd.AddCommand(
		makeSystemAdminCmd,
		makeMemberCmd,
	)
}

func makeSystemAdminCmdF(cmd *cobra.Command, args []string) error {
	initDBCommandContextCobra(cmd)
	if len(args) < 1 {
		return errors.New("Enter at least one user.")
	}

	users := getUsersFromUserArgs(args)
	for i, user := range users {
		if user == nil {
			return errors.New("Unable to find user '" + args[i] + "'")
		}

		if _, err := app.UpdateUserRoles(user.Id, "system_admin system_user"); err != nil {
			return err
		}
	}

	return nil
}

func makeMemberCmdF(cmd *cobra.Command, args []string) error {
	initDBCommandContextCobra(cmd)
	if len(args) < 1 {
		return errors.New("Enter at least one user.")
	}

	users := getUsersFromUserArgs(args)
	for i, user := range users {
		if user == nil {
			return errors.New("Unable to find user '" + args[i] + "'")
		}

		if _, err := app.UpdateUserRoles(user.Id, "system_user"); err != nil {
			return err
		}
	}

	return nil
}
