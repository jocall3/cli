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

var usersMeBiometricsDeregister = cli.Command{
	Name:            "deregister",
	Usage:           "Removes all enrolled biometric data associated with the user's account for\nsecurity reasons.",
	Flags:           []cli.Flag{},
	Action:          handleUsersMeBiometricsDeregister,
	HideHelpCommand: true,
}

var usersMeBiometricsEnroll = cli.Command{
	Name:  "enroll",
	Usage: "Initiates the enrollment process for biometric authentication (e.g.,\nfingerprint, facial scan) to enable secure and convenient access to sensitive\nfeatures. Requires a biometric signature for initial proof.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "biometric-signature",
			Usage:    "Base64 encoded representation of the biometric template or proof.",
			BodyPath: "biometricSignature",
		},
		&requestflag.Flag[string]{
			Name:     "biometric-type",
			Usage:    "The type of biometric data being enrolled.",
			BodyPath: "biometricType",
		},
		&requestflag.Flag[any]{
			Name:     "device-id",
			Usage:    "The ID of the device on which the biometric is being enrolled.",
			BodyPath: "deviceId",
		},
		&requestflag.Flag[any]{
			Name:     "device-name",
			Usage:    "Optional: A friendly name for the device, if not already linked.",
			BodyPath: "deviceName",
		},
	},
	Action:          handleUsersMeBiometricsEnroll,
	HideHelpCommand: true,
}

var usersMeBiometricsStatus = cli.Command{
	Name:            "status",
	Usage:           "Retrieves the current status of biometric enrollments for the authenticated\nuser.",
	Flags:           []cli.Flag{},
	Action:          handleUsersMeBiometricsStatus,
	HideHelpCommand: true,
}

var usersMeBiometricsVerify = cli.Command{
	Name:  "verify",
	Usage: "Performs real-time biometric verification to authorize sensitive actions or\naccess protected resources, using a one-time biometric signature.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "biometric-signature",
			Usage:    "Base64 encoded representation of the one-time biometric proof for verification.",
			BodyPath: "biometricSignature",
		},
		&requestflag.Flag[string]{
			Name:     "biometric-type",
			Usage:    "The type of biometric data being verified.",
			BodyPath: "biometricType",
		},
		&requestflag.Flag[any]{
			Name:     "device-id",
			Usage:    "The ID of the device initiating the biometric verification.",
			BodyPath: "deviceId",
		},
	},
	Action:          handleUsersMeBiometricsVerify,
	HideHelpCommand: true,
}

func handleUsersMeBiometricsDeregister(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
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

	return client.Users.Me.Biometrics.Deregister(ctx, options...)
}

func handleUsersMeBiometricsEnroll(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.UserMeBiometricEnrollParams{}

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
	_, err = client.Users.Me.Biometrics.Enroll(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me:biometrics enroll", obj, format, transform)
}

func handleUsersMeBiometricsStatus(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Users.Me.Biometrics.Status(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me:biometrics status", obj, format, transform)
}

func handleUsersMeBiometricsVerify(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.UserMeBiometricVerifyParams{}

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
	_, err = client.Users.Me.Biometrics.Verify(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me:biometrics verify", obj, format, transform)
}
