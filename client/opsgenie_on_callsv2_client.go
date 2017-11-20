package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/oncallsv2"
)

// OpsGenieOnCallsV2Client is a client for working with on calls users
type OpsGenieOnCallsV2Client struct {
	RestClient
}

// OpsGenieOnCallsV2Client sets the embedded OpsGenieClient type of the OpsGenieOnCallsV2Client
func (cli *OpsGenieOnCallsV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

// Get returns on calls information
func (cli *OpsGenieOnCallsV2Client) Get(req oncallsv2.GetOnCallsRequest) (*oncallsv2.GetOnCallsResponse, error) {
	var response oncallsv2.GetOnCallsResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
