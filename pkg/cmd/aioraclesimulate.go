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

var aiOracleSimulateRunAdvanced = cli.Command{
	Name:  "run-advanced",
	Usage: "Engages the Quantum Oracle for highly complex, multi-variable simulations,\nallowing precise control over numerous financial parameters, market conditions,\nand personal events to generate deep, predictive insights and sensitivity\nanalysis.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			BodyPath: "prompt",
		},
		&requestflag.Flag[[]any]{
			Name:     "scenario",
			BodyPath: "scenarios",
		},
	},
	Action:          handleAIOracleSimulateRunAdvanced,
	HideHelpCommand: true,
}

var aiOracleSimulateRunStandard = cli.Command{
	Name:  "run-standard",
	Usage: "Submits a hypothetical scenario to the Quantum Oracle AI for standard financial\nimpact analysis. The AI simulates the effect on the user's current financial\nstate and provides a summary.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "A natural language prompt describing the scenario.",
			BodyPath: "prompt",
		},
		&requestflag.Flag[any]{
			Name:     "parameters",
			Usage:    "Structured parameters for the simulation.",
			BodyPath: "parameters",
		},
	},
	Action:          handleAIOracleSimulateRunStandard,
	HideHelpCommand: true,
}

func handleAIOracleSimulateRunAdvanced(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParams{}

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
	_, err = client.AI.Oracle.Simulate.RunAdvanced(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:oracle:simulate run-advanced", obj, format, transform)
}

func handleAIOracleSimulateRunStandard(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunStandardParams{}

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
	_, err = client.AI.Oracle.Simulate.RunStandard(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:oracle:simulate run-standard", obj, format, transform)
}
