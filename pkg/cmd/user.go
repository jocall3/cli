// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/1231-cli/internal/apiquery"
	"github.com/stainless-sdks/1231-cli/internal/requestflag"
	"github.com/stainless-sdks/1231-go"
	"github.com/stainless-sdks/1231-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var usersLogin = cli.Command{
	Name:  "login",
	Usage: "Authenticates a user and creates a secure session, returning access tokens. May\nrequire MFA depending on user settings.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "mfa-code",
			Usage:    "Multi-factor authentication code, if required.",
			BodyPath: "mfaCode",
		},
	},
	Action:          handleUsersLogin,
	HideHelpCommand: true,
}

var usersRegister = cli.Command{
	Name:  "register",
	Usage: "Registers a new user account with , initiating the onboarding process. Requires\nbasic user details.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "email",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "phone",
			BodyPath: "phone",
		},
	},
	Action:          handleUsersRegister,
	HideHelpCommand: true,
}

func handleUsersLogin(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.UserLoginParams{}

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
	_, err = client.Users.Login(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users login", obj, format, transform)
}

func handleUsersRegister(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.UserRegisterParams{}

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
	_, err = client.Users.Register(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users register", obj, format, transform)
}
