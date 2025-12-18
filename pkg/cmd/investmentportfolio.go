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

var investmentsPortfoliosCreate = cli.Command{
	Name:  "create",
	Usage: "Creates a new investment portfolio, with options for initial asset allocation.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "currency",
			BodyPath: "currency",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "risk-tolerance",
			BodyPath: "riskTolerance",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "ai-auto-allocate",
			BodyPath: "aiAutoAllocate",
		},
		&requestflag.Flag[float64]{
			Name:     "initial-investment",
			BodyPath: "initialInvestment",
		},
		&requestflag.Flag[string]{
			Name:     "linked-account-id",
			BodyPath: "linkedAccountId",
		},
	},
	Action:          handleInvestmentsPortfoliosCreate,
	HideHelpCommand: true,
}

var investmentsPortfoliosRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Retrieves detailed information for a specific investment portfolio, including\nholdings, performance, and AI insights.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "portfolio-id",
		},
	},
	Action:          handleInvestmentsPortfoliosRetrieve,
	HideHelpCommand: true,
}

var investmentsPortfoliosUpdate = cli.Command{
	Name:  "update",
	Usage: "Updates high-level details of an investment portfolio, such as name or risk\ntolerance.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "portfolio-id",
		},
		&requestflag.Flag[string]{
			Name:     "ai-rebalancing-frequency",
			BodyPath: "aiRebalancingFrequency",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "risk-tolerance",
			BodyPath: "riskTolerance",
		},
	},
	Action:          handleInvestmentsPortfoliosUpdate,
	HideHelpCommand: true,
}

var investmentsPortfoliosList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a summary of all investment portfolios linked to the user's account.",
	Flags: []cli.Flag{
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
	},
	Action:          handleInvestmentsPortfoliosList,
	HideHelpCommand: true,
}

var investmentsPortfoliosRebalance = cli.Command{
	Name:  "rebalance",
	Usage: "Triggers an AI-driven rebalancing process for a specific investment portfolio\nbased on a target risk tolerance or strategy.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "portfolio-id",
		},
		&requestflag.Flag[string]{
			Name:     "target-risk-tolerance",
			BodyPath: "targetRiskTolerance",
		},
		&requestflag.Flag[bool]{
			Name:     "confirmation-required",
			BodyPath: "confirmationRequired",
		},
		&requestflag.Flag[bool]{
			Name:     "dry-run",
			Usage:    "If true, returns the proposed changes without executing them.",
			BodyPath: "dryRun",
		},
	},
	Action:          handleInvestmentsPortfoliosRebalance,
	HideHelpCommand: true,
}

func handleInvestmentsPortfoliosCreate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioNewParams{}

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
	_, err = client.Investments.Portfolios.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "investments:portfolios create", obj, format, transform)
}

func handleInvestmentsPortfoliosRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("portfolio-id") && len(unusedArgs) > 0 {
		cmd.Set("portfolio-id", unusedArgs[0])
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
	_, err = client.Investments.Portfolios.Get(ctx, cmd.Value("portfolio-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "investments:portfolios retrieve", obj, format, transform)
}

func handleInvestmentsPortfoliosUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("portfolio-id") && len(unusedArgs) > 0 {
		cmd.Set("portfolio-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioUpdateParams{}

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
	_, err = client.Investments.Portfolios.Update(
		ctx,
		cmd.Value("portfolio-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "investments:portfolios update", obj, format, transform)
}

func handleInvestmentsPortfoliosList(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioListParams{}

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
	_, err = client.Investments.Portfolios.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "investments:portfolios list", obj, format, transform)
}

func handleInvestmentsPortfoliosRebalance(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("portfolio-id") && len(unusedArgs) > 0 {
		cmd.Set("portfolio-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioRebalanceParams{}

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
	_, err = client.Investments.Portfolios.Rebalance(
		ctx,
		cmd.Value("portfolio-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "investments:portfolios rebalance", obj, format, transform)
}
