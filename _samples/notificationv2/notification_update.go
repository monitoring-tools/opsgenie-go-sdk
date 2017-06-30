package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	notificationCli, _ := cli.Notificationv2()

	identifier := &notificationv2.Identifier{
		UserID: "4ca72bed-c1a9-40f8-a140-44cc957bf31a",
		RuleID: "231550b5-e6fc-452b-b786-43fe8a3cd092",
	}

	criteria := notificationv2.Criteria{
		Type: "match-all-conditions",
		Conditions: []notificationv2.Condition{
			{
				Field: "extra-properties",
				Key: "system",
				Not: true,
				Operation: "equals",
				ExpectedValue: "mysql",
			},
		},
	}

	timeRestriction := notificationv2.TimeRestriction{
		Type: notificationv2.WeekendAndTimeOfDayTimeRestriction,
		Restrictions: []notificationv2.Restriction{
			{
				StartDay:  notificationv2.Monday,
				EndDay:    notificationv2.Friday,
				StartHour: 3,
				EndHour:   15,
				StartMin:  5,
				EndMin:    30,
			},
		},
	}

	response, err := notificationCli.Update(notificationv2.UpdateNotificationRequest{
		Identifier: identifier,
		Name: "Test create-alert(changed)",
		Criteria: criteria,
		TimeRestriction: timeRestriction,
		Enabled: true,
		Order: 2,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf(
			"ID: %s, Name: %s, Enabled: %t\n",
			response.Notification.ID,
			response.Notification.Name,
			response.Notification.Enabled,
		)
	}
}
