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

var goalsCreate = cli.Command{
	Name:  "create",
	Usage: "Creates a new long-term financial goal, with optional AI plan generation.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Name of the new financial goal.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "target-amount",
			Usage:    "The target monetary amount for the goal.",
			BodyPath: "targetAmount",
		},
		&requestflag.Flag[any]{
			Name:     "target-date",
			Usage:    "The target completion date for the goal.",
			BodyPath: "targetDate",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of financial goal.",
			BodyPath: "type",
		},
		&requestflag.Flag[[]any]{
			Name:     "contributing-account",
			Usage:    "Optional: List of account IDs initially contributing to this goal.",
			BodyPath: "contributingAccounts",
		},
		&requestflag.Flag[any]{
			Name:     "generate-ai-plan",
			Usage:    "If true, AI will automatically generate a strategic plan for the goal.",
			BodyPath: "generateAIPlan",
		},
		&requestflag.Flag[any]{
			Name:     "initial-contribution",
			Usage:    "Optional: Initial amount to contribute to the goal.",
			BodyPath: "initialContribution",
		},
		&requestflag.Flag[string]{
			Name:     "risk-tolerance",
			Usage:    "Desired risk tolerance for investments related to this goal.",
			BodyPath: "riskTolerance",
		},
	},
	Action:          handleGoalsCreate,
	HideHelpCommand: true,
}

var goalsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Retrieves detailed information for a specific financial goal, including current\nprogress, AI strategic plan, and related insights.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "goal-id",
		},
	},
	Action:          handleGoalsRetrieve,
	HideHelpCommand: true,
}

var goalsUpdate = cli.Command{
	Name:  "update",
	Usage: "Updates the parameters of an existing financial goal, such as target amount,\ndate, or contributing accounts. This may trigger an AI plan recalculation.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "goal-id",
		},
		&requestflag.Flag[[]any]{
			Name:     "contributing-account",
			Usage:    "Updated list of account IDs contributing to this goal.",
			BodyPath: "contributingAccounts",
		},
		&requestflag.Flag[any]{
			Name:     "generate-ai-plan",
			Usage:    "If true, AI will recalculate and update the strategic plan for the goal.",
			BodyPath: "generateAIPlan",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Updated name of the financial goal.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "risk-tolerance",
			Usage:    "Updated risk tolerance for investments related to this goal.",
			BodyPath: "riskTolerance",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Updated status of the goal's progress.",
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "target-amount",
			Usage:    "The updated target monetary amount for the goal.",
			BodyPath: "targetAmount",
		},
		&requestflag.Flag[any]{
			Name:     "target-date",
			Usage:    "The updated target completion date for the goal.",
			BodyPath: "targetDate",
		},
	},
	Action:          handleGoalsUpdate,
	HideHelpCommand: true,
}

var goalsList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all financial goals defined by the user, including their\nprogress and associated AI plans.",
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
	Action:          handleGoalsList,
	HideHelpCommand: true,
}

var goalsDelete = cli.Command{
	Name:  "delete",
	Usage: "Deletes a specific financial goal from the user's profile.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name: "goal-id",
		},
	},
	Action:          handleGoalsDelete,
	HideHelpCommand: true,
}

func handleGoalsCreate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.GoalNewParams{}

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
	_, err = client.Goals.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "goals create", obj, format, transform)
}

func handleGoalsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("goal-id") && len(unusedArgs) > 0 {
		cmd.Set("goal-id", unusedArgs[0])
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
	_, err = client.Goals.Get(ctx, interface{}(cmd.Value("goal-id").(any)), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "goals retrieve", obj, format, transform)
}

func handleGoalsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("goal-id") && len(unusedArgs) > 0 {
		cmd.Set("goal-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.GoalUpdateParams{}

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
	_, err = client.Goals.Update(
		ctx,
		interface{}(cmd.Value("goal-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "goals update", obj, format, transform)
}

func handleGoalsList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.GoalListParams{}

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
	_, err = client.Goals.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "goals list", obj, format, transform)
}

func handleGoalsDelete(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("goal-id") && len(unusedArgs) > 0 {
		cmd.Set("goal-id", unusedArgs[0])
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

	return client.Goals.Delete(ctx, interface{}(cmd.Value("goal-id").(any)), options...)
}
