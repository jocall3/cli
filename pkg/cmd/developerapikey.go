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

var developersAPIKeysCreate = cli.Command{
	Name:  "create",
	Usage: "Generates a new API key for the developer application with specified scopes and\nan optional expiration.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "A descriptive name for the API key.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]any]{
			Name:     "scope",
			Usage:    "List of permissions to grant to this API key.",
			BodyPath: "scopes",
		},
		&requestflag.Flag[any]{
			Name:     "expires-in-days",
			Usage:    "Optional: Number of days until the API key expires. If omitted, it will not expire.",
			BodyPath: "expiresInDays",
		},
	},
	Action:          handleDevelopersAPIKeysCreate,
	HideHelpCommand: true,
}

var developersAPIKeysList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of API keys issued to the authenticated developer application.",
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
	Action:          handleDevelopersAPIKeysList,
	HideHelpCommand: true,
}

var developersAPIKeysRevoke = cli.Command{
	Name:  "revoke",
	Usage: "Revokes an existing API key, disabling its access immediately.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "key-id",
		},
	},
	Action:          handleDevelopersAPIKeysRevoke,
	HideHelpCommand: true,
}

func handleDevelopersAPIKeysCreate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.DeveloperAPIKeyNewParams{}

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
	_, err = client.Developers.APIKeys.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "developers:api-keys create", obj, format, transform)
}

func handleDevelopersAPIKeysList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.DeveloperAPIKeyListParams{}

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
	_, err = client.Developers.APIKeys.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "developers:api-keys list", obj, format, transform)
}

func handleDevelopersAPIKeysRevoke(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("key-id") && len(unusedArgs) > 0 {
		cmd.Set("key-id", unusedArgs[0])
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

	return client.Developers.APIKeys.Revoke(ctx, interface{}(cmd.Value("key-id").(any)), options...)
}
