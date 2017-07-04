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

	request := userv2.ListUserRequest{
		Limit: 10,
		Offset: 0,
		Sort: userv2.UsernameSortField,
		Order: userv2.AscSortType,
		Query: map[string]string{
			userv2.UsernameQueryField: "John",
		},
	}

	response, err := userCli.List(request)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, user := range response.Users {
			fmt.Println(user)
		}
	}
}