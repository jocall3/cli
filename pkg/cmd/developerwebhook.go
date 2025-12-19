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

var developersWebhooksCreate = cli.Command{
	Name:  "create",
	Usage: "Establishes a new webhook subscription, allowing a developer application to\nreceive real-time notifications for specified events (e.g., new transaction,\naccount update) via a provided callback URL.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "callback-url",
			Usage:    "The URL to which webhook events will be sent.",
			BodyPath: "callbackUrl",
		},
		&requestflag.Flag[[]any]{
			Name:     "event",
			Usage:    "List of event types to subscribe to.",
			BodyPath: "events",
		},
		&requestflag.Flag[any]{
			Name:     "secret",
			Usage:    "Optional: A custom shared secret for verifying webhook payloads. If omitted, one will be generated.",
			BodyPath: "secret",
		},
	},
	Action:          handleDevelopersWebhooksCreate,
	HideHelpCommand: true,
}

var developersWebhooksUpdate = cli.Command{
	Name:  "update",
	Usage: "Modifies an existing webhook subscription, allowing changes to the callback URL,\nsubscribed events, or activation status.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "subscription-id",
		},
		&requestflag.Flag[any]{
			Name:     "callback-url",
			Usage:    "Updated URL where webhook events will be sent.",
			BodyPath: "callbackUrl",
		},
		&requestflag.Flag[[]any]{
			Name:     "event",
			Usage:    "Updated list of event types subscribed to.",
			BodyPath: "events",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Updated status of the webhook subscription.",
			BodyPath: "status",
		},
	},
	Action:          handleDevelopersWebhooksUpdate,
	HideHelpCommand: true,
}

var developersWebhooksList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all active webhook subscriptions for the authenticated\ndeveloper application, detailing endpoint URLs, subscribed events, and current\nstatus.",
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
	Action:          handleDevelopersWebhooksList,
	HideHelpCommand: true,
}

var developersWebhooksDelete = cli.Command{
	Name:  "delete",
	Usage: "Deletes an existing webhook subscription, stopping all future event\nnotifications to the specified callback URL.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "subscription-id",
		},
	},
	Action:          handleDevelopersWebhooksDelete,
	HideHelpCommand: true,
}

func handleDevelopersWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.DeveloperWebhookNewParams{}

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
	_, err = client.Developers.Webhooks.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "developers:webhooks create", obj, format, transform)
}

func handleDevelopersWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.DeveloperWebhookUpdateParams{}

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
	_, err = client.Developers.Webhooks.Update(
		ctx,
		interface{}(cmd.Value("subscription-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "developers:webhooks update", obj, format, transform)
}

func handleDevelopersWebhooksList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.DeveloperWebhookListParams{}

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
	_, err = client.Developers.Webhooks.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "developers:webhooks list", obj, format, transform)
}

func handleDevelopersWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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

	return client.Developers.Webhooks.Delete(ctx, interface{}(cmd.Value("subscription-id").(any)), options...)
}
