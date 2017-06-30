package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	notificationCli, _ := cli.Notificationv2()

	response, err := notificationCli.Delete(notificationv2.DeleteNotificationRequest{
		Identifier: &notificationv2.Identifier{
			UserID: "4ca72bed-c1a9-40f8-a140-44cc957bf31a",
			RuleID: "88c00129-6e68-40c5-a812-757dc140018d",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Result: %s, Request ID: %s\n", response.Result, response.RequestID)
	}
}
