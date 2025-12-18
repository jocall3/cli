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

var corporateAnomaliesList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a comprehensive list of AI-detected financial anomalies across\ntransactions, payments, and corporate cards that require immediate review and\npotential action to mitigate risk and ensure compliance.",
	Flags: []cli.Flag{
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "end-date",
			Usage:     "The end date for the query range (inclusive).",
			QueryPath: "endDate",
		},
		&requestflag.Flag[string]{
			Name:      "entity-type",
			Usage:     "Filter anomalies by the type of financial entity they are related to.",
			QueryPath: "entityType",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "The number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "severity",
			Usage:     "Filter anomalies by their AI-assessed severity level.",
			QueryPath: "severity",
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "start-date",
			Usage:     "The start date for the query range (inclusive).",
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
		&requestflag.Flag[string]{
			Name: "anomaly-id",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:     "resolution-notes",
			BodyPath: "resolutionNotes",
		},
	},
	Action:          handleCorporateAnomaliesUpdateStatus,
	HideHelpCommand: true,
}

func handleCorporateAnomaliesList(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyListParams{}

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
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("anomaly-id") && len(unusedArgs) > 0 {
		cmd.Set("anomaly-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyUpdateStatusParams{}

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
		cmd.Value("anomaly-id").(string),
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
