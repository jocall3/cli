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

var usersMeDevicesList = cli.Command{
	Name:  "list",
	Usage: "Retrieves a list of all devices linked to the user's account, including mobile\nphones, tablets, and desktops, indicating their last active status and security\nposture.",
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
	Action:          handleUsersMeDevicesList,
	HideHelpCommand: true,
}

var usersMeDevicesDeregister = cli.Command{
	Name:  "deregister",
	Usage: "Removes a specific device from the user's linked devices, revoking its access\nand requiring re-registration for future use. Useful for lost or compromised\ndevices.",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
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
			Usage:    "Type of the device being registered.",
			BodyPath: "deviceType",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Model of the device.",
			BodyPath: "model",
		},
		&requestflag.Flag[any]{
			Name:     "os",
			Usage:    "Operating system of the device.",
			BodyPath: "os",
		},
		&requestflag.Flag[any]{
			Name:     "biometric-signature",
			Usage:    "Optional: Base64 encoded biometric signature for initial enrollment (e.g., for Passkey registration).",
			BodyPath: "biometricSignature",
		},
		&requestflag.Flag[any]{
			Name:     "device-name",
			Usage:    "Optional: A friendly name for the device.",
			BodyPath: "deviceName",
		},
		&requestflag.Flag[any]{
			Name:     "push-token",
			Usage:    "Optional: Push notification token for the device.",
			BodyPath: "pushToken",
		},
	},
	Action:          handleUsersMeDevicesRegister,
	HideHelpCommand: true,
}

func handleUsersMeDevicesList(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserMeDeviceListParams{}

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
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
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

	return client.Users.Me.Devices.Deregister(ctx, cmd.Value("device-id").(any), options...)
}

func handleUsersMeDevicesRegister(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserMeDeviceRegisterParams{}

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
