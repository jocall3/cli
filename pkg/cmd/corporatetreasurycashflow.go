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

var corporateTreasuryCashFlowGetForecast = cli.Command{
	Name:  "get-forecast",
	Usage: "Retrieves an advanced AI-driven cash flow forecast for the organization,\nprojecting liquidity, identifying potential surpluses or deficits, and providing\nrecommendations for optimal treasury management.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "forecast-horizon-days",
			Usage:     "The number of days into the future for which to generate the cash flow forecast (e.g., 30, 90, 180).",
			Default:   90,
			QueryPath: "forecastHorizonDays",
		},
		&requestflag.Flag[any]{
			Name:      "include-scenario-analysis",
			Usage:     "If true, the forecast will include best-case and worst-case scenario analysis alongside the most likely projection.",
			QueryPath: "includeScenarioAnalysis",
		},
	},
	Action:          handleCorporateTreasuryCashFlowGetForecast,
	HideHelpCommand: true,
}

func handleCorporateTreasuryCashFlowGetForecast(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateTreasuryCashFlowGetForecastParams{}

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
	_, err = client.Corporate.Treasury.CashFlow.GetForecast(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:treasury:cash-flow get-forecast", obj, format, transform)
}
