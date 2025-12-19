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

var aiOracleSimulateRunAdvanced = cli.Command{
	Name:  "run-advanced",
	Usage: "Engages the Quantum Oracle for highly complex, multi-variable simulations,\nallowing precise control over numerous financial parameters, market conditions,\nand personal events to generate deep, predictive insights and sensitivity\nanalysis.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "prompt",
			Usage:    "A natural language prompt describing the complex, multi-variable scenario.",
			BodyPath: "prompt",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "scenario",
			BodyPath: "scenarios",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "global-economic-factors",
			Usage:    "Optional: Global economic conditions to apply to all scenarios.",
			BodyPath: "globalEconomicFactors",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "personal-assumptions",
			Usage:    "Optional: Personal financial assumptions to override defaults.",
			BodyPath: "personalAssumptions",
		},
	},
	Action:          handleAIOracleSimulateRunAdvanced,
	HideHelpCommand: true,
}

var aiOracleSimulateRunStandard = cli.Command{
	Name:  "run-standard",
	Usage: "Submits a hypothetical scenario to the Quantum Oracle AI for standard financial\nimpact analysis. The AI simulates the effect on the user's current financial\nstate and provides a summary.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "prompt",
			Usage:    "A natural language prompt describing the 'what-if' scenario.",
			BodyPath: "prompt",
		},
		&requestflag.Flag[any]{
			Name:     "parameters",
			Usage:    "Optional structured parameters to guide the simulation (e.g., duration, amount, risk tolerance).",
			BodyPath: "parameters",
		},
	},
	Action:          handleAIOracleSimulateRunStandard,
	HideHelpCommand: true,
}

func handleAIOracleSimulateRunAdvanced(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AIOracleSimulateRunAdvancedParams{}

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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AIOracleSimulateRunStandardParams{}

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
