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

var usersMeDevicesList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all devices linked to the user's account, including mobile\nphones, tablets, and desktops, indicating their last active status and security\nposture.",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of items to return.",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "The number of items to skip before starting to collect the result set.",
			QueryPath: "offset",
		},
	},
	Action:          handleUsersMeDevicesList,
	HideHelpCommand: true,
}

var usersMeDevicesDeregister = cli.Command{
	Name:  "deregister",
	Usage: "Removes a specific device from the user's linked devices, revoking its access\nand requiring re-registration for future use. Useful for lost or compromised\ndevices.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "device-id",
		},
	},
	Action:          handleUsersMeDevicesDeregister,
	HideHelpCommand: true,
}

var usersMeDevicesRegister = cli.Command{
	Name:  "register",
	Usage: "Registers a new device for secure access and multi-factor authentication,\nassociating it with the user's profile. This typically initiates a biometric or\nMFA enrollment flow.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "device-type",
			BodyPath: "deviceType",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "os",
			BodyPath: "os",
		},
		&requestflag.Flag[string]{
			Name:     "biometric-signature",
			Usage:    "Base64 encoded proof of biometric capability.",
			BodyPath: "biometricSignature",
		},
		&requestflag.Flag[string]{
			Name:     "device-name",
			BodyPath: "deviceName",
		},
	},
	Action:          handleUsersMeDevicesRegister,
	HideHelpCommand: true,
}

func handleUsersMeDevicesList(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.UserMeDeviceListParams{}

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
	_, err = client.Users.Me.Devices.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me:devices list", obj, format, transform)
}

func handleUsersMeDevicesDeregister(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("device-id") && len(unusedArgs) > 0 {
		cmd.Set("device-id", unusedArgs[0])
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

	return client.Users.Me.Devices.Deregister(ctx, cmd.Value("device-id").(string), options...)
}

func handleUsersMeDevicesRegister(ctx context.Context, cmd *cli.Command) error {
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jamesburvelocallaghaniiicitibankdemobusinessinc.UserMeDeviceRegisterParams{}

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
	_, err = client.Users.Me.Devices.Register(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:me:devices register", obj, format, transform)
}
