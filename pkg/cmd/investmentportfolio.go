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

var investmentsPortfoliosCreate = cli.Command{
	Name:  "create",
	Usage: "Creates a new investment portfolio, with options for initial asset allocation.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "currency",
			Usage:    "ISO 4217 currency code of the portfolio.",
			BodyPath: "currency",
		},
		&requestflag.Flag[any]{
			Name:     "initial-investment",
			Usage:    "Initial amount to invest into the portfolio.",
			BodyPath: "initialInvestment",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Name for the new investment portfolio.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "risk-tolerance",
			Usage:    "Desired risk tolerance for this portfolio.",
			BodyPath: "riskTolerance",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "General type or strategy of the portfolio.",
			BodyPath: "type",
		},
		&requestflag.Flag[any]{
			Name:     "ai-auto-allocate",
			Usage:    "If true, AI will automatically allocate initial investment based on risk tolerance.",
			BodyPath: "aiAutoAllocate",
		},
		&requestflag.Flag[any]{
			Name:     "linked-account-id",
			Usage:    "Optional: ID of a linked account to fund the initial investment.",
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
		&requestflag.Flag[any]{
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
		&requestflag.Flag[any]{
			Name: "portfolio-id",
		},
		&requestflag.Flag[string]{
			Name:     "ai-rebalancing-frequency",
			Usage:    "Updated frequency for AI-driven rebalancing.",
			BodyPath: "aiRebalancingFrequency",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Updated name of the portfolio.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "risk-tolerance",
			Usage:    "Updated risk tolerance for this portfolio. May trigger rebalancing.",
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
	Action:          handleInvestmentsPortfoliosList,
	HideHelpCommand: true,
}

var investmentsPortfoliosRebalance = cli.Command{
	Name:  "rebalance",
	Usage: "Triggers an AI-driven rebalancing process for a specific investment portfolio\nbased on a target risk tolerance or strategy.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "portfolio-id",
		},
		&requestflag.Flag[string]{
			Name:     "target-risk-tolerance",
			Usage:    "The desired risk tolerance for rebalancing the portfolio.",
			BodyPath: "targetRiskTolerance",
		},
		&requestflag.Flag[any]{
			Name:     "confirmation-required",
			Usage:    "If true, user confirmation is required before executing actual trades after a dry run.",
			Default:  true,
			BodyPath: "confirmationRequired",
		},
		&requestflag.Flag[any]{
			Name:     "dry-run",
			Usage:    "If true, only simulate the rebalance without executing trades. Returns proposed trades.",
			Default:  true,
			BodyPath: "dryRun",
		},
	},
	Action:          handleInvestmentsPortfoliosRebalance,
	HideHelpCommand: true,
}

func handleInvestmentsPortfoliosCreate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.InvestmentPortfolioNewParams{}

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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Investments.Portfolios.Get(ctx, cmd.Value("portfolio-id").(any), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "investments:portfolios retrieve", obj, format, transform)
}

func handleInvestmentsPortfoliosUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("portfolio-id") && len(unusedArgs) > 0 {
		cmd.Set("portfolio-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.InvestmentPortfolioUpdateParams{}

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
		cmd.Value("portfolio-id").(any),
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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.InvestmentPortfolioListParams{}

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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("portfolio-id") && len(unusedArgs) > 0 {
		cmd.Set("portfolio-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.InvestmentPortfolioRebalanceParams{}

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
		cmd.Value("portfolio-id").(any),
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
