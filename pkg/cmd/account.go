// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jocall3/1231-cli/internal/apiquery"
	"github.com/jocall3/1231-cli/internal/requestflag"
	"github.com/jocall3/go"
	"github.com/jocall3/go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var accountsLink = cli.Command{
	Name:  "link",
	Usage: "Begins the secure process of linking a new external financial institution (e.g.,\nanother bank, investment platform) to the user's profile, typically involving a\nthird-party tokenized flow.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "country-code",
			Usage:    "Two-letter ISO country code of the institution.",
			BodyPath: "countryCode",
		},
		&requestflag.Flag[any]{
			Name:     "institution-name",
			Usage:    "Name of the financial institution to link.",
			BodyPath: "institutionName",
		},
		&requestflag.Flag[any]{
			Name:     "provider-identifier",
			Usage:    "Optional: Specific identifier for a third-party linking provider (e.g., 'plaid', 'finicity').",
			BodyPath: "providerIdentifier",
		},
		&requestflag.Flag[any]{
			Name:     "redirect-uri",
			Usage:    "Optional: URI to redirect the user after completing the external authentication flow.",
			BodyPath: "redirectUri",
		},
	},
	Action:          handleAccountsLink,
	HideHelpCommand: true,
}

var accountsRetrieveDetails = cli.Command{
	Name:  "retrieve-details",
	Usage: "Retrieves comprehensive analytics for a specific financial account, including\nhistorical balance trends, projected cash flow, and AI-driven insights into\nspending patterns.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "account-id",
		},
	},
	Action:          handleAccountsRetrieveDetails,
	HideHelpCommand: true,
}

var accountsRetrieveMe = cli.Command{
	Name:  "retrieve-me",
	Usage: "Fetches a comprehensive, real-time list of all external financial accounts\nlinked to the user's profile, including consolidated balances and institutional\ndetails.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "limit",
			Usage:     "Maximum number of items to return in a single page.",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "offset",
			Usage:     "Number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
	},
	Action:          handleAccountsRetrieveMe,
	HideHelpCommand: true,
}

var accountsRetrieveStatements = cli.Command{
	Name:  "retrieve-statements",
	Usage: "Fetches digital statements for a specific account, allowing filtering by date\nrange and format.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "account-id",
		},
		&requestflag.Flag[any]{
			Name:      "month",
			Usage:     "Month for the statement (1-12).",
			QueryPath: "month",
		},
		&requestflag.Flag[any]{
			Name:      "year",
			Usage:     "Year for the statement.",
			QueryPath: "year",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Desired format for the statement. Use 'application/json' Accept header for download links.",
			Default:   "pdf",
			QueryPath: "format",
		},
	},
	Action:          handleAccountsRetrieveStatements,
	HideHelpCommand: true,
}

func handleAccountsLink(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AccountLinkParams{}

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
	_, err = client.Accounts.Link(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts link", obj, format, transform)
}

func handleAccountsRetrieveDetails(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Accounts.GetDetails(ctx, interface{}(cmd.Value("account-id").(any)), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts retrieve-details", obj, format, transform)
}

func handleAccountsRetrieveMe(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AccountGetMeParams{}

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
	_, err = client.Accounts.GetMe(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts retrieve-me", obj, format, transform)
}

func handleAccountsRetrieveStatements(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AccountGetStatementsParams{}

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
	_, err = client.Accounts.GetStatements(
		ctx,
		interface{}(cmd.Value("account-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts retrieve-statements", obj, format, transform)
}
