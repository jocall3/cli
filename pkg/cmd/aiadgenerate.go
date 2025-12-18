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

var aiAdsGenerateAdvanced = cli.Command{
	Name:  "advanced",
	Usage: "Submits a highly customized request to generate a video ad, allowing\nfine-grained control over artistic style, aspect ratio, voiceover, background\nmusic, target audience, and call-to-action elements for professional-grade\nproductions.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "aspect-ratio",
			BodyPath: "aspectRatio",
		},
		&requestflag.Flag[string]{
			Name:     "audience-target",
			BodyPath: "audienceTarget",
		},
		&requestflag.Flag[[]string]{
			Name:     "brand-asset",
			BodyPath: "brandAssets",
		},
		&requestflag.Flag[[]string]{
			Name:     "brand-color",
			BodyPath: "brandColors",
		},
		&requestflag.Flag[any]{
			Name:     "call-to-action",
			BodyPath: "callToAction",
		},
		&requestflag.Flag[int64]{
			Name:     "length-seconds",
			BodyPath: "lengthSeconds",
		},
		&requestflag.Flag[string]{
			Name:     "style",
			BodyPath: "style",
		},
		&requestflag.Flag[string]{
			Name:     "voiceover-style",
			BodyPath: "voiceoverStyle",
		},
		&requestflag.Flag[string]{
			Name:     "voiceover-text",
			BodyPath: "voiceoverText",
		},
	},
	Action:          handleAIAdsGenerateAdvanced,
	HideHelpCommand: true,
}

var aiAdsGenerateStandard = cli.Command{
	Name:  "standard",
	Usage: "Submits a request to generate a high-quality video ad using the advanced Veo 2.0\ngenerative AI model. This is an asynchronous operation, suitable for standard ad\ncontent creation.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "aspect-ratio",
			BodyPath: "aspectRatio",
		},
		&requestflag.Flag[[]string]{
			Name:     "brand-color",
			BodyPath: "brandColors",
		},
		&requestflag.Flag[int64]{
			Name:     "length-seconds",
			BodyPath: "lengthSeconds",
		},
		&requestflag.Flag[string]{
			Name:     "style",
			BodyPath: "style",
		},
	},
	Action:          handleAIAdsGenerateStandard,
	HideHelpCommand: true,
}

func handleAIAdsGenerateAdvanced(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParams{}

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
	_, err = client.AI.Ads.Generate.Advanced(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:ads:generate advanced", obj, format, transform)
}

func handleAIAdsGenerateStandard(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateStandardParams{}

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
	_, err = client.AI.Ads.Generate.Standard(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:ads:generate standard", obj, format, transform)
}
