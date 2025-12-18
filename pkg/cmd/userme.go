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

var usersMeRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Fetches the complete and dynamically updated profile information for the\ncurrently authenticated user, encompassing personal details, security status,\ngamification level, loyalty points, and linked identity attributes.",
	Flags:           []cli.Flag{},
	Action:          handleUsersMeRetrieve,
	HideHelpCommand: true,
}

var usersMeUpdate = cli.Command{
	Name:  "update",
	Usage: "Updates selected fields of the currently authenticated user's profile\ninformation.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "address",
			BodyPath: "address",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Updated full name of the user.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "phone",
			Usage:    "Updated primary phone number of the user.",
			BodyPath: "phone",
		},
		&requestflag.Flag[any]{
			Name:     "preferences",
			Usage:    "User's personalized preferences for the platform.",
			BodyPath: "preferences",
		},
	},
	Action:          handleUsersMeUpdate,
	HideHelpCommand: true,
}

func handleUsersMeRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.Users.Me.Get(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me retrieve", obj, format, transform)
}

func handleUsersMeUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserMeUpdateParams{}

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
	_, err = client.Users.Me.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me update", obj, format, transform)
}
