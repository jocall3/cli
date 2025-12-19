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

var corporateAnomaliesList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a comprehensive list of AI-detected financial anomalies across\ntransactions, payments, and corporate cards that require immediate review and\npotential action to mitigate risk and ensure compliance.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "end-date",
			Usage:     "End date for filtering results (inclusive, YYYY-MM-DD).",
			QueryPath: "endDate",
		},
		&requestflag.Flag[string]{
			Name:      "entity-type",
			Usage:     "Filter anomalies by the type of financial entity they are related to.",
			QueryPath: "entityType",
		},
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
			Usage:     "Filter anomalies by their AI-assessed severity level.",
			QueryPath: "severity",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			Usage:     "Start date for filtering results (inclusive, YYYY-MM-DD).",
			QueryPath: "startDate",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter anomalies by their current review status.",
			Default:   "New",
			QueryPath: "status",
		},
	},
	Action:          handleCorporateAnomaliesList,
	HideHelpCommand: true,
}

var corporateAnomaliesUpdateStatus = cli.Command{
	Name:  "update-status",
	Usage: "Updates the review status of a specific financial anomaly, allowing compliance\nofficers to mark it as dismissed, resolved, or escalate for further\ninvestigation after thorough AI-assisted and human review.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "anomaly-id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The new status for the financial anomaly.",
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "resolution-notes",
			Usage:    "Optional notes regarding the resolution or dismissal of the anomaly.",
			BodyPath: "resolutionNotes",
		},
	},
	Action:          handleCorporateAnomaliesUpdateStatus,
	HideHelpCommand: true,
}

func handleCorporateAnomaliesList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.CorporateAnomalyListParams{}

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
	_, err = client.Corporate.Anomalies.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:anomalies list", obj, format, transform)
}

func handleCorporateAnomaliesUpdateStatus(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("anomaly-id") && len(unusedArgs) > 0 {
		cmd.Set("anomaly-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.CorporateAnomalyUpdateStatusParams{}

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
	_, err = client.Corporate.Anomalies.UpdateStatus(
		ctx,
		interface{}(cmd.Value("anomaly-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:anomalies update-status", obj, format, transform)
}
