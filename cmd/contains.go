package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/open-policy-agent/opa/rego"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(containsCmd)
}

var containsCmd = &cobra.Command{
	Use:   "contains",
	Short: "Check if an IP address is part of a CIDR block",
	Long: `Check if an IP address is part of a CIDR block.
Usage:

cidr contains <CIDR> <IP>
cidr contains <CIDR> <TESTCIDR>

Example:
cidr contains 192.168.0.0/16 192.168.7.42
cidr contains 192.168.0.0/16 192.168.0.0/24`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cidr := args[0]
		ip := args[1]

		res, err := Contains(cidr, ip)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	},
}

func Contains(cidr string, ip string) (bool, error) {
	ctx := context.Background()

	module := `
		package cidr
		default contains = false
		contains = true {
			basecidr := input.basecidr
			testip := input.testip
			net.cidr_contains(basecidr, testip)
		}
	`

	reg := rego.New(
		rego.Query("data.cidr.contains"),
		rego.Module("contains", module),
		rego.Input(
			map[string]interface{}{
				"basecidr": cidr,
				"testip":   ip,
			},
		),
	)

	res, err := reg.Eval(ctx)
	if err != nil {
		return false, err
	}

	boolRes, err := strconv.ParseBool(fmt.Sprintf("%v", res[0].Expressions[0]))
	if err != nil {
		return false, err
	}

	return boolRes, nil
}
