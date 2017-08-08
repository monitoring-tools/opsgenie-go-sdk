package main

import (
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/userv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	userCli, _ := cli.UserV2()

	request := userv2.GetUserRequest{
		Identifier: &userv2.Identifier{
			ID:     "1b4d54a9-7f07-44ef-9995-b2c0b1adc674",
			Expand: userv2.ContactExpandableField,
		},
	}

	response, err := userCli.Get(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.User)
	}
}
