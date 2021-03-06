package main

import (
	"github.com/opsgenie/opsgenie-go-sdk/_samples/constants"
	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	"fmt"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func main() {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(constants.APIKey)

	alertCli, _ := cli.AlertV2()

	request := alertsv2.GetSavedSearchRequest{
		Name: "list-blue-team-alerts",
	}

	response, err := alertCli.GetSavedSearch(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		savedSearch := response.SavedSearch

		fmt.Println("Name: " + savedSearch.Name)
		fmt.Println("Query: " + savedSearch.Query)
	}
}