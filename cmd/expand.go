package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/open-policy-agent/opa/rego"
)



func init() {
  rootCmd.AddCommand(expandCmd)
}

var expandCmd = &cobra.Command{
	Use:   "expand",
	Short: "Expand a provided CIDR block into its IP addresses",
	Long: `Expand a provided CIDR block into its IP addresses.

Example:
cidr expand 192.168.0.0/24
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cidr := args[0]

		res, err := expand(cidr)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	},
}

func expand(cidr string) (string, error) {
	ctx := context.Background()

	module := `
		package cidr
		expand[ips] {
			basecidr := input.basecidr
			ips := net.cidr_expand(basecidr)
		}
	`

	reg := rego.New(
		rego.Query("data.cidr.expand"),
		rego.Module("expand", module),
		rego.Input(
			map[string]interface{}{
				"basecidr": cidr,
			},
		),
	)

	rs, err := reg.Eval(ctx)
	if err != nil {
		return "", err
	}

	type cidrIps struct {
		CIDR string      `json:"cidr"`
		IPs  interface{} `json:"ips"`
	}

	res, err := json.Marshal(
		cidrIps{
			CIDR: cidr,
			IPs:  rs[0].Expressions[0].Value,
		},
	)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
