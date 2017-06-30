package main

import (
	"fmt"
	"encoding/json"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
	//"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	//cli.SetAPIKey(constants.APIKey)
	cli.SetAPIKey("f17414c6-4c4b-42f3-9e61-be7c3c65685a")

	notificationCli, _ := cli.Notificationv2()

	response, err := notificationCli.Get(notificationv2.GetNotificationRequest{
		Identifier: &notificationv2.Identifier{
			UserID: "4ca72bed-c1a9-40f8-a140-44cc957bf31a",
			RuleID: "15212908-31cd-4d32-9c30-10bb22cb0008",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		notification := response.Notification
		notificationJson, _ := json.Marshal(notification)

		fmt.Println(string(notificationJson))
	}
}
