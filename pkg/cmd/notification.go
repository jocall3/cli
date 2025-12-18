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

var notificationsListUserNotifications = cli.Command{
	Name:  "list-user-notifications",
	Usage: "Retrieves a paginated list of personalized notifications and proactive AI alerts\nfor the authenticated user, allowing filtering by status and severity.",
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
		&requestflag.Flag[string]{
			Name:      "severity",
			Usage:     "Filter notifications by AI-assigned severity level.",
			QueryPath: "severity",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter notifications by their read status.",
			QueryPath: "status",
		},
	},
	Action:          handleNotificationsListUserNotifications,
	HideHelpCommand: true,
}

var notificationsMarkAsRead = cli.Command{
	Name:  "mark-as-read",
	Usage: "Marks a specific user notification as read.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "notification-id",
		},
	},
	Action:          handleNotificationsMarkAsRead,
	HideHelpCommand: true,
}

func handleNotificationsListUserNotifications(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.NotificationListUserNotificationsParams{}

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
	_, err = client.Notifications.ListUserNotifications(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications list-user-notifications", obj, format, transform)
}

func handleNotificationsMarkAsRead(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("notification-id") && len(unusedArgs) > 0 {
		cmd.Set("notification-id", unusedArgs[0])
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
	_, err = client.Notifications.MarkAsRead(ctx, cmd.Value("notification-id").(any), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications mark-as-read", obj, format, transform)
}
