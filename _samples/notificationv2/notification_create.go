package main

import (
	"fmt"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	//"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
)

func main() {

	cli := new(ogcli.OpsGenieClient)
	//cli.SetAPIKey(constants.APIKey)
	cli.SetAPIKey("f17414c6-4c4b-42f3-9e61-be7c3c65685a")

	notificationCli, _ := cli.Notificationv2()

	identifier := &notificationv2.Identifier{
		UserID: "4ca72bed-c1a9-40f8-a140-44cc957bf31a",
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

	steps := []notificationv2.Step{
		{
			SendAfter: notificationv2.SendAfter{
				TimeAmount: 1,
				TimeUnit: notificationv2.Minutes,
			},
			Contact: notificationv2.Contact{
				Method: notificationv2.EmailNotifyMethod,
				To: "vlamug@gmail.com",
			},
			Enabled: true,
		},
	}

	repeat := notificationv2.Repeat{Enabled: false}

	response, err := notificationCli.Create(notificationv2.CreateNotificationRequest{
		Identifier: identifier,
		Name: "Test create-alert",
		ActionType: notificationv2.CreateAlertActionType,
		Criteria: criteria,
		TimeRestriction: timeRestriction,
		Order: 1,
		Steps: steps,
		Repeat: repeat,
		Enabled: true,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("ID: %s, Name: %s, Enabled: %t\n", response.ID, response.Name, response.Enabled)
	}
}
