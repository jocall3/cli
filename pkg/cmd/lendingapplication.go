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

var lendingApplicationsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Retrieves the current status and detailed information for a submitted loan\napplication, including AI underwriting outcomes, approved terms, and next steps.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "application-id",
		},
	},
	Action:          handleLendingApplicationsRetrieve,
	HideHelpCommand: true,
}

var lendingApplicationsSubmit = cli.Command{
	Name:  "submit",
	Usage: "Submits a new loan application, which is instantly processed and underwritten by\nour Quantum AI, providing rapid decisions and personalized loan offers based on\nreal-time financial health data.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "loan-amount",
			Usage:    "The desired loan amount.",
			BodyPath: "loanAmount",
		},
		&requestflag.Flag[string]{
			Name:     "loan-purpose",
			Usage:    "The purpose of the loan.",
			BodyPath: "loanPurpose",
		},
		&requestflag.Flag[any]{
			Name:     "repayment-term-months",
			Usage:    "The desired repayment term in months.",
			BodyPath: "repaymentTermMonths",
		},
		&requestflag.Flag[any]{
			Name:     "additional-notes",
			Usage:    "Optional notes or details for the application.",
			BodyPath: "additionalNotes",
		},
		&requestflag.Flag[any]{
			Name:     "co-applicant",
			Usage:    "Optional: Details of a co-applicant for the loan.",
			BodyPath: "coApplicant",
		},
	},
	Action:          handleLendingApplicationsSubmit,
	HideHelpCommand: true,
}

func handleLendingApplicationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("application-id") && len(unusedArgs) > 0 {
		cmd.Set("application-id", unusedArgs[0])
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
	_, err = client.Lending.Applications.Get(ctx, cmd.Value("application-id").(any), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lending:applications retrieve", obj, format, transform)
}

func handleLendingApplicationsSubmit(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.LendingApplicationSubmitParams{}

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
	_, err = client.Lending.Applications.Submit(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lending:applications submit", obj, format, transform)
}
