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

var accountsOverdraftSettingsRetrieveOverdraftSettings = cli.Command{
	Name:  "retrieve-overdraft-settings",
	Usage: "Retrieves the current overdraft protection settings for a specific account.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
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
		&requestflag.Flag[string]{
			Name: "account-id",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "fee-preference",
			BodyPath: "feePreference",
		},
		&requestflag.Flag[string]{
			Name:     "linked-savings-account-id",
			BodyPath: "linkedSavingsAccountId",
		},
		&requestflag.Flag[bool]{
			Name:     "link-to-savings",
			BodyPath: "linkToSavings",
		},
		&requestflag.Flag[float64]{
			Name:     "protection-limit",
			BodyPath: "protectionLimit",
		},
	},
	Action:          handleAccountsOverdraftSettingsUpdateOverdraftSettings,
	HideHelpCommand: true,
}

func handleAccountsOverdraftSettingsRetrieveOverdraftSettings(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Accounts.OverdraftSettings.GetOverdraftSettings(ctx, cmd.Value("account-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts:overdraft-settings retrieve-overdraft-settings", obj, format, transform)
}

func handleAccountsOverdraftSettingsUpdateOverdraftSettings(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.AccountOverdraftSettingUpdateOverdraftSettingsParams{}

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
		cmd.Value("account-id").(string),
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
