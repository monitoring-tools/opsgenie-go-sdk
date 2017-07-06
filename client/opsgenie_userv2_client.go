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
func (cli *OpsGenieUserV2Client) List(req userv2.ListUsersRequest) (*userv2.ListUsersResponse, error) {
	var response userv2.ListUsersResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieUserV2Client) Create(req userv2.CreateUserRequest) (*userv2.CreateUserResponse, error) {
	var response userv2.CreateUserResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieUserV2Client) Get(req userv2.GetUserRequest) (*userv2.GetUserResponse, error) {
	var response userv2.GetUserResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieUserV2Client) Update(req userv2.UpdateUserRequest) (*userv2.UpdateUserResponse, error) {
	var response userv2.UpdateUserResponse
	err := cli.sendPatchRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGenieUserV2Client) Delete(req userv2.DeleteUserRequest) (*userv2.DeleteUserResponse, error) {
	var response userv2.DeleteUserResponse
	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
