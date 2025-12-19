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

var aiOracleSimulationsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Retrieves the full, detailed results of a specific financial simulation by its\nID.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "simulation-id",
		},
	},
	Action:          handleAIOracleSimulationsRetrieve,
	HideHelpCommand: true,
}

var aiOracleSimulationsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all financial simulations previously run by the user,\nincluding their status and summaries.",
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
	Action:          handleAIOracleSimulationsList,
	HideHelpCommand: true,
}

var aiOracleSimulationsDelete = cli.Command{
	Name:  "delete",
	Usage: "Deletes a previously run financial simulation and its results.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "simulation-id",
		},
	},
	Action:          handleAIOracleSimulationsDelete,
	HideHelpCommand: true,
}

func handleAIOracleSimulationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("simulation-id") && len(unusedArgs) > 0 {
		cmd.Set("simulation-id", unusedArgs[0])
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
	_, err = client.AI.Oracle.Simulations.Get(ctx, interface{}(cmd.Value("simulation-id").(any)), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:oracle:simulations retrieve", obj, format, transform)
}

func handleAIOracleSimulationsList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AIOracleSimulationListParams{}

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
	_, err = client.AI.Oracle.Simulations.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:oracle:simulations list", obj, format, transform)
}

func handleAIOracleSimulationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("simulation-id") && len(unusedArgs) > 0 {
		cmd.Set("simulation-id", unusedArgs[0])
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

	return client.AI.Oracle.Simulations.Delete(ctx, interface{}(cmd.Value("simulation-id").(any)), options...)
}
