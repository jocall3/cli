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

var corporateRiskFraudRulesCreate = cli.Command{
	Name:  "create",
	Usage: "Creates a new custom AI-powered fraud detection rule, allowing organizations to\ndefine specific criteria, risk scores, and automated responses to evolving\nthreat landscapes.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "action",
			Usage:    "Action to take when a fraud rule is triggered.",
			BodyPath: "action",
		},
		&requestflag.Flag[any]{
			Name:     "criteria",
			Usage:    "Criteria that define when a fraud rule should trigger.",
			BodyPath: "criteria",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Detailed description of what the rule detects.",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Name of the new fraud rule.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "severity",
			Usage:    "Severity level when this rule is triggered.",
			BodyPath: "severity",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Initial status of the rule.",
			BodyPath: "status",
		},
	},
	Action:          handleCorporateRiskFraudRulesCreate,
	HideHelpCommand: true,
}

var corporateRiskFraudRulesUpdate = cli.Command{
	Name:  "update",
	Usage: "Updates an existing custom AI-powered fraud detection rule, modifying its\ncriteria, actions, or status.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "rule-id",
		},
		&requestflag.Flag[any]{
			Name:     "action",
			Usage:    "Action to take when a fraud rule is triggered.",
			BodyPath: "action",
		},
		&requestflag.Flag[any]{
			Name:     "criteria",
			Usage:    "Criteria that define when a fraud rule should trigger.",
			BodyPath: "criteria",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Updated description of what the rule detects.",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Updated name of the fraud rule.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "severity",
			Usage:    "Updated severity level.",
			BodyPath: "severity",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Updated status of the rule.",
			BodyPath: "status",
		},
	},
	Action:          handleCorporateRiskFraudRulesUpdate,
	HideHelpCommand: true,
}

var corporateRiskFraudRulesList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of AI-powered fraud detection rules currently active for the\norganization, including their parameters, thresholds, and associated actions\n(e.g., flag, block, alert).",
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
	Action:          handleCorporateRiskFraudRulesList,
	HideHelpCommand: true,
}

var corporateRiskFraudRulesDelete = cli.Command{
	Name:  "delete",
	Usage: "Deletes a specific custom AI-powered fraud detection rule.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "rule-id",
		},
	},
	Action:          handleCorporateRiskFraudRulesDelete,
	HideHelpCommand: true,
}

func handleCorporateRiskFraudRulesCreate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleNewParams{}

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
	_, err = client.Corporate.Risk.Fraud.Rules.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:risk:fraud:rules create", obj, format, transform)
}

func handleCorporateRiskFraudRulesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("rule-id") && len(unusedArgs) > 0 {
		cmd.Set("rule-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleUpdateParams{}

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
	_, err = client.Corporate.Risk.Fraud.Rules.Update(
		ctx,
		cmd.Value("rule-id").(any),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:risk:fraud:rules update", obj, format, transform)
}

func handleCorporateRiskFraudRulesList(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleListParams{}

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
	_, err = client.Corporate.Risk.Fraud.Rules.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "corporate:risk:fraud:rules list", obj, format, transform)
}

func handleCorporateRiskFraudRulesDelete(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("rule-id") && len(unusedArgs) > 0 {
		cmd.Set("rule-id", unusedArgs[0])
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

	return client.Corporate.Risk.Fraud.Rules.Delete(ctx, cmd.Value("rule-id").(any), options...)
}
