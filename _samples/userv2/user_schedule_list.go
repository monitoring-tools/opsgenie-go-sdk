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
	cli.SetAPIKey("bdd67d0e-92c6-47b5-87cf-f2a9940629e2")

	userCli, _ := cli.UserV2()

	request := userv2.ListUserSchedulesRequest{
		Identifier: &userv2.Identifier{UserName:"22275390-e95b-47df-8c13-6edc469c0b26"},
	}

	response, err := userCli.ListSchedules(request)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, schedule := range response.Schedules {
			fmt.Println(schedule)
		}
	}
}