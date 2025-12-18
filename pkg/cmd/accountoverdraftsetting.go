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

var accountsOverdraftSettingsRetrieveOverdraftSettings = cli.Command{
	Name:  "retrieve-overdraft-settings",
	Usage: "Retrieves the current overdraft protection settings for a specific account.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "account-id",
		},
	},
	Action:          handleAccountsOverdraftSettingsRetrieveOverdraftSettings,
	HideHelpCommand: true,
}

var accountsOverdraftSettingsUpdateOverdraftSettings = cli.Command{
	Name:  "update-overdraft-settings",
	Usage: "Updates the overdraft protection settings for a specific account, enabling or\ndisabling protection and configuring preferences.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "account-id",
		},
		&requestflag.Flag[any]{
			Name:     "enabled",
			Usage:    "Enable or disable overdraft protection.",
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "fee-preference",
			Usage:    "New preference for how overdraft fees are handled.",
			BodyPath: "feePreference",
		},
		&requestflag.Flag[any]{
			Name:     "linked-savings-account-id",
			Usage:    "New ID of the linked savings account, if `linkToSavings` is true. Set to null to unlink.",
			BodyPath: "linkedSavingsAccountId",
		},
		&requestflag.Flag[any]{
			Name:     "link-to-savings",
			Usage:    "Enable or disable linking to a savings account for overdraft coverage.",
			BodyPath: "linkToSavings",
		},
		&requestflag.Flag[any]{
			Name:     "protection-limit",
			Usage:    "New maximum amount for overdraft protection. Set to null to remove limit.",
			BodyPath: "protectionLimit",
		},
	},
	Action:          handleAccountsOverdraftSettingsUpdateOverdraftSettings,
	HideHelpCommand: true,
}

func handleAccountsOverdraftSettingsRetrieveOverdraftSettings(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Accounts.OverdraftSettings.GetOverdraftSettings(ctx, cmd.Value("account-id").(any), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts:overdraft-settings retrieve-overdraft-settings", obj, format, transform)
}

func handleAccountsOverdraftSettingsUpdateOverdraftSettings(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AccountOverdraftSettingUpdateOverdraftSettingsParams{}

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
	_, err = client.Accounts.OverdraftSettings.UpdateOverdraftSettings(
		ctx,
		cmd.Value("account-id").(any),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts:overdraft-settings update-overdraft-settings", obj, format, transform)
}
