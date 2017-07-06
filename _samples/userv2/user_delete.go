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

	request := userv2.DeleteUserRequest{
		Identifier: &userv2.Identifier{UserName:"user-name@test.com"},
	}

	response, err := userCli.Delete(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}