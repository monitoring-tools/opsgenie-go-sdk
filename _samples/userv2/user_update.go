package main

import (
	"fmt"

	//"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/userv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	//cli.SetAPIKey(constants.APIKey)
	cli.SetAPIKey("f17414c6-4c4b-42f3-9e61-be7c3c65685a")

	userCli, _ := cli.UserV2()

	request := userv2.UpdateUserRequest{
		Identifier: &userv2.Identifier{
			ID: "1b4d54a9-7f07-44ef-9995-b2c0b1adc674",
			Expand: userv2.ContactExpandableField,
		},
		FullName: "Lex Luthor",
		Role: userv2.Role{Name:userv2.AdminRole},
		SkypeUsername: "lex.luthor",
		Tags: userv2.Tags{"updated"},
		Details: userv2.Details{"test": []string{"updated"}},
		Locale: "de_CH",
		Timezone: "US/Arizona",
		UserAddress: userv2.UserAddress{
			City: "Phoenix",
			State: "Arizona",
		},
	}

	response, err := userCli.Update(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}