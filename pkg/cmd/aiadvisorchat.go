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

var aiAdvisorChatRetrieveHistory = cli.Command{
	Name:  "retrieve-history",
	Usage: "Fetches the full conversation history with the Quantum AI Advisor for a given\nsession or user.",
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
		&requestflag.Flag[any]{
			Name:      "session-id",
			Usage:     "Optional: Filter history by a specific session ID. If omitted, recent conversations will be returned.",
			QueryPath: "sessionId",
		},
	},
	Action:          handleAIAdvisorChatRetrieveHistory,
	HideHelpCommand: true,
}

var aiAdvisorChatSendMessage = cli.Command{
	Name:  "send-message",
	Usage: "Initiates or continues a sophisticated conversation with Quantum, the AI\nAdvisor. Quantum can provide advanced financial insights, execute complex tasks\nvia an expanding suite of intelligent tools, and learn from user interactions to\noffer hyper-personalized guidance.",
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "function-response",
			Usage:    "Optional: The output from a tool function that the AI previously requested to be executed.",
			BodyPath: "functionResponse",
		},
		&requestflag.Flag[any]{
			Name:     "message",
			Usage:    "The user's textual input to the AI Advisor.",
			BodyPath: "message",
		},
		&requestflag.Flag[any]{
			Name:     "session-id",
			Usage:    "Optional: Session ID to continue a conversation. If omitted, a new session is started.",
			BodyPath: "sessionId",
		},
	},
	Action:          handleAIAdvisorChatSendMessage,
	HideHelpCommand: true,
}

func handleAIAdvisorChatRetrieveHistory(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AIAdvisorChatGetHistoryParams{}

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
	_, err = client.AI.Advisor.Chat.GetHistory(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:advisor:chat retrieve-history", obj, format, transform)
}

func handleAIAdvisorChatSendMessage(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AIAdvisorChatSendMessageParams{}

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
	_, err = client.AI.Advisor.Chat.SendMessage(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:advisor:chat send-message", obj, format, transform)
}
