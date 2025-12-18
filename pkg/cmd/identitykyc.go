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

var identityKYCRetrieveStatus = cli.Command{
	Name:            "retrieve-status",
	Usage:           "Retrieves the current status of the user's Know Your Customer (KYC) verification\nprocess.",
	Flags:           []cli.Flag{},
	Action:          handleIdentityKYCRetrieveStatus,
	HideHelpCommand: true,
}

var identityKYCSubmit = cli.Command{
	Name:  "submit",
	Usage: "Submits Know Your Customer (KYC) documentation, such as identity proofs and\naddress verification, for AI-accelerated compliance and identity verification,\ncrucial for higher service tiers and regulatory adherence.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "country-of-issue",
			Usage:    "The two-letter ISO country code where the document was issued.",
			BodyPath: "countryOfIssue",
		},
		&requestflag.Flag[any]{
			Name:     "document-number",
			Usage:    "The identification number on the document.",
			BodyPath: "documentNumber",
		},
		&requestflag.Flag[string]{
			Name:     "document-type",
			Usage:    "The type of KYC document being submitted.",
			BodyPath: "documentType",
		},
		&requestflag.Flag[any]{
			Name:     "expiration-date",
			Usage:    "The expiration date of the document (YYYY-MM-DD).",
			BodyPath: "expirationDate",
		},
		&requestflag.Flag[any]{
			Name:     "issue-date",
			Usage:    "The issue date of the document (YYYY-MM-DD).",
			BodyPath: "issueDate",
		},
		&requestflag.Flag[[]any]{
			Name:     "additional-document",
			Usage:    "Array of additional documents (e.g., utility bills) as base64 encoded images.",
			BodyPath: "additionalDocuments",
		},
		&requestflag.Flag[any]{
			Name:     "document-back-image",
			Usage:    "Base64 encoded image of the back of the document (if applicable).",
			BodyPath: "documentBackImage",
		},
		&requestflag.Flag[any]{
			Name:     "document-front-image",
			Usage:    "Base64 encoded image of the front of the document. Use 'application/json' with base64 string, or 'multipart/form-data' for direct file upload.",
			BodyPath: "documentFrontImage",
		},
	},
	Action:          handleIdentityKYCSubmit,
	HideHelpCommand: true,
}

func handleIdentityKYCRetrieveStatus(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.Identity.KYC.GetStatus(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "identity:kyc retrieve-status", obj, format, transform)
}

func handleIdentityKYCSubmit(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.IdentityKYCSubmitParams{}

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
	_, err = client.Identity.KYC.Submit(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "identity:kyc submit", obj, format, transform)
}
