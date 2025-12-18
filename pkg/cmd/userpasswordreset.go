// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jocall3/cli/internal/apiquery"
	"github.com/jocall3/cli/internal/requestflag"
	"github.com/jocall3/go"
	"github.com/jocall3/go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var usersPasswordResetConfirm = cli.Command{
	Name:  "confirm",
	Usage: "Confirms the password reset using the received verification code and sets a new\npassword.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "identifier",
			Usage:    "User's email or phone number used for verification.",
			BodyPath: "identifier",
		},
		&requestflag.Flag[any]{
			Name:     "new-password",
			Usage:    "The new password for the user account.",
			BodyPath: "newPassword",
		},
		&requestflag.Flag[any]{
			Name:     "verification-code",
			Usage:    "The verification code received via email or SMS.",
			BodyPath: "verificationCode",
		},
	},
	Action:          handleUsersPasswordResetConfirm,
	HideHelpCommand: true,
}

var usersPasswordResetInitiate = cli.Command{
	Name:  "initiate",
	Usage: "Starts the password reset flow by sending a verification code or link to the\nuser's registered email or phone.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "identifier",
			Usage:    "User's email or phone number for verification.",
			BodyPath: "identifier",
		},
	},
	Action:          handleUsersPasswordResetInitiate,
	HideHelpCommand: true,
}

func handleUsersPasswordResetConfirm(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserPasswordResetConfirmParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.PasswordReset.Confirm(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:password-reset confirm", obj, format, transform)
}

func handleUsersPasswordResetInitiate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserPasswordResetInitiateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.PasswordReset.Initiate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:password-reset initiate", obj, format, transform)
}
