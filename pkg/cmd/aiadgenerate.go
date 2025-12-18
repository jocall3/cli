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
		&requestflag.Flag[any]{
			Name:     "length-seconds",
			Usage:    "Desired length of the video in seconds.",
			BodyPath: "lengthSeconds",
		},
		&requestflag.Flag[any]{
			Name:     "prompt",
			Usage:    "The textual prompt to guide the AI video generation.",
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "style",
			Usage:    "Artistic style of the video.",
			BodyPath: "style",
		},
		&requestflag.Flag[string]{
			Name:     "aspect-ratio",
			Usage:    "Aspect ratio of the video (e.g., 16:9 for widescreen, 9:16 for vertical shorts).",
			Default:  "16:9",
			BodyPath: "aspectRatio",
		},
		&requestflag.Flag[string]{
			Name:     "audience-target",
			Usage:    "Target audience for the ad, influencing tone and visuals.",
			BodyPath: "audienceTarget",
		},
		&requestflag.Flag[string]{
			Name:     "background-music-genre",
			Usage:    "Genre of background music.",
			BodyPath: "backgroundMusicGenre",
		},
		&requestflag.Flag[[]any]{
			Name:     "brand-asset",
			Usage:    "URLs to brand assets (e.g., logos, specific imagery) to be incorporated.",
			BodyPath: "brandAssets",
		},
		&requestflag.Flag[[]any]{
			Name:     "brand-color",
			Usage:    "Optional: Hex color codes to influence the video's aesthetic.",
			BodyPath: "brandColors",
		},
		&requestflag.Flag[any]{
			Name:     "call-to-action",
			Usage:    "Call-to-action text and URL to be displayed.",
			BodyPath: "callToAction",
		},
		&requestflag.Flag[[]any]{
			Name:     "keyword",
			Usage:    "Optional: Additional keywords to guide the AI's content generation.",
			BodyPath: "keywords",
		},
		&requestflag.Flag[string]{
			Name:     "voiceover-style",
			Usage:    "Style/tone for the AI voiceover.",
			BodyPath: "voiceoverStyle",
		},
		&requestflag.Flag[any]{
			Name:     "voiceover-text",
			Usage:    "Optional: Text for an AI-generated voiceover.",
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
		&requestflag.Flag[any]{
			Name:     "length-seconds",
			Usage:    "Desired length of the video in seconds.",
			BodyPath: "lengthSeconds",
		},
		&requestflag.Flag[any]{
			Name:     "prompt",
			Usage:    "The textual prompt to guide the AI video generation.",
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "style",
			Usage:    "Artistic style of the video.",
			BodyPath: "style",
		},
		&requestflag.Flag[string]{
			Name:     "aspect-ratio",
			Usage:    "Aspect ratio of the video (e.g., 16:9 for widescreen, 9:16 for vertical shorts).",
			Default:  "16:9",
			BodyPath: "aspectRatio",
		},
		&requestflag.Flag[[]any]{
			Name:     "brand-color",
			Usage:    "Optional: Hex color codes to influence the video's aesthetic.",
			BodyPath: "brandColors",
		},
		&requestflag.Flag[[]any]{
			Name:     "keyword",
			Usage:    "Optional: Additional keywords to guide the AI's content generation.",
			BodyPath: "keywords",
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
