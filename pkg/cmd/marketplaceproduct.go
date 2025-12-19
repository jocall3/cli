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

var marketplaceProductsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a personalized, AI-curated list of products and services from the\nPlato AI marketplace, tailored to the user's financial profile, goals, and\nspending patterns. Includes options for filtering and advanced search.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ai-personalization-level",
			Usage:     "Filter by AI personalization level (e.g., low, medium, high). 'High' means highly relevant to user's specific needs.",
			QueryPath: "aiPersonalizationLevel",
		},
		&requestflag.Flag[string]{
			Name:      "category",
			Usage:     "Filter products by category (e.g., loans, insurance, credit_cards, investments).",
			QueryPath: "category",
		},
		&requestflag.Flag[any]{
			Name:      "limit",
			Usage:     "Maximum number of items to return in a single page.",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "min-rating",
			Usage:     "Minimum user rating for products (0-5).",
			QueryPath: "minRating",
		},
		&requestflag.Flag[any]{
			Name:      "offset",
			Usage:     "Number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
	},
	Action:          handleMarketplaceProductsList,
	HideHelpCommand: true,
}

var marketplaceProductsSimulateImpact = cli.Command{
	Name:  "simulate-impact",
	Usage: "Uses the Quantum Oracle to simulate the long-term financial impact of purchasing\nor subscribing to a specific marketplace product, such as a loan, investment, or\ninsurance policy, on the user's overall financial health and goals.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "product-id",
		},
		&requestflag.Flag[any]{
			Name:     "simulation-parameters",
			Usage:    "Dynamic parameters specific to the product type (e.g., loan amount, investment term).",
			BodyPath: "simulationParameters",
		},
	},
	Action:          handleMarketplaceProductsSimulateImpact,
	HideHelpCommand: true,
}

func handleMarketplaceProductsList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.MarketplaceProductListParams{}

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
	_, err = client.Marketplace.Products.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "marketplace:products list", obj, format, transform)
}

func handleMarketplaceProductsSimulateImpact(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("product-id") && len(unusedArgs) > 0 {
		cmd.Set("product-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.MarketplaceProductSimulateImpactParams{}

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
	_, err = client.Marketplace.Products.SimulateImpact(
		ctx,
		interface{}(cmd.Value("product-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "marketplace:products simulate-impact", obj, format, transform)
}
