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

var usersMePreferencesRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Retrieves the user's deep personalization preferences, including AI\ncustomization settings, notification channel priorities, thematic choices, and\ndata sharing consents.",
	Flags:           []cli.Flag{},
	Action:          handleUsersMePreferencesRetrieve,
	HideHelpCommand: true,
}

var usersMePreferencesUpdate = cli.Command{
	Name:  "update",
	Usage: "Updates the user's deep personalization preferences, allowing dynamic control\nover AI behavior, notification delivery, thematic choices, and data privacy\nsettings.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "ai-interaction-mode",
			Usage:    "How the user prefers to interact with AI (proactive advice, balanced, or only on demand).",
			BodyPath: "aiInteractionMode",
		},
		&requestflag.Flag[any]{
			Name:     "data-sharing-consent",
			Usage:    "Consent status for sharing anonymized data for AI improvement and personalized offers.",
			BodyPath: "dataSharingConsent",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "notification-channels",
			Usage:    "Preferred channels for receiving notifications.",
			BodyPath: "notificationChannels",
		},
		&requestflag.Flag[any]{
			Name:     "preferred-language",
			Usage:    "Preferred language for the user interface.",
			BodyPath: "preferredLanguage",
		},
		&requestflag.Flag[any]{
			Name:     "theme",
			Usage:    "Preferred UI theme (e.g., Light-Default, Dark-Quantum).",
			BodyPath: "theme",
		},
		&requestflag.Flag[string]{
			Name:     "transaction-grouping",
			Usage:    "Default grouping preference for transaction lists.",
			BodyPath: "transactionGrouping",
		},
	},
	Action:          handleUsersMePreferencesUpdate,
	HideHelpCommand: true,
}

func handleUsersMePreferencesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Users.Me.Preferences.Get(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me:preferences retrieve", obj, format, transform)
}

func handleUsersMePreferencesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserMePreferenceUpdateParams{}

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
	_, err = client.Users.Me.Preferences.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me:preferences update", obj, format, transform)
}
