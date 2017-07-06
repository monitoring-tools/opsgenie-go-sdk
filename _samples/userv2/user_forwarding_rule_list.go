package main

import (
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/userv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	userCli, _ := cli.UserV2()

	request := userv2.ListUserForwardingRulesRequest{
		Identifier: &userv2.Identifier{UserName:"22275390-e95b-47df-8c13-6edc469c0b26"},
	}

	response, err := userCli.ListForwardingRules(request)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, rule := range response.ForwardingRules {
			fmt.Println(rule)
		}
	}
}