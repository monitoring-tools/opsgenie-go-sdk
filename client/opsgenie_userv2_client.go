package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/userv2"
)

type OpsGenieUserV2Client struct {
	RestClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieAlertV2Client.
func (cli *OpsGenieUserV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

// List method retrieves the list of users from OpsGenie
func (cli *OpsGenieUserV2Client) List(req userv2.ListUserRequest) (*userv2.ListUserResponse, error) {
	var response userv2.ListUserResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
