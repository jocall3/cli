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

var marketplaceOffersRedeem = cli.Command{
	Name:  "redeem",
	Usage: "Redeems a personalized, exclusive offer from the Plato AI marketplace, often\nresulting in a discount, special rate, or credit to the user's account.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "offer-id",
		},
		&requestflag.Flag[any]{
			Name:     "payment-account-id",
			Usage:    "Optional: The ID of the account to use for any associated payment or credit.",
			BodyPath: "paymentAccountId",
		},
	},
	Action:          handleMarketplaceOffersRedeem,
	HideHelpCommand: true,
}

func handleMarketplaceOffersRedeem(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("offer-id") && len(unusedArgs) > 0 {
		cmd.Set("offer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.MarketplaceOfferRedeemParams{}

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
	_, err = client.Marketplace.Offers.Redeem(
		ctx,
		interface{}(cmd.Value("offer-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "marketplace:offers redeem", obj, format, transform)
}
