package main

import (
	"fmt"

	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	notificationCli, _ := cli.Notificationv2()

	response, err := notificationCli.List(notificationv2.ListNotificationRequest{
		Identifier: &notificationv2.Identifier{
			UserID: "4ca72bed-c1a9-40f8-a140-44cc957bf31a",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i, notification := range response.Notifications {
			fmt.Printf("%d. %s\n", i, notification.Name)
		}
	}
}
