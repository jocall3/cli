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

var corporatePerformSanctionScreening = cli.Command{
	Name:  "perform-sanction-screening",
	Usage: "Executes a real-time screening of an individual or entity against global\nsanction lists and watchlists.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-type",
			BodyPath: "entityType",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			BodyPath: "country",
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:     "date-of-birth",
			BodyPath: "dateOfBirth",
		},
	},
	Action:          handleCorporatePerformSanctionScreening,
	HideHelpCommand: true,
}

func handleCorporatePerformSanctionScreening(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporatePerformSanctionScreeningParams{}

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
	_, err = client.Corporate.PerformSanctionScreening(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate perform-sanction-screening", obj, format, transform)
}
