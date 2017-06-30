package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/notificationv2"
)

type OpsGenieNotificationV2Client struct {
	RestClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieTeamClient.
func (cli *OpsGenieNotificationV2Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieNotificationV2Client) Get(req notificationv2.GetNotificationRequest) (
	*notificationv2.DetailedNotificationResponse,
	error,
) {
	var response notificationv2.DetailedNotificationResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieNotificationV2Client) List(req notificationv2.ListNotificationRequest) (
	*notificationv2.ListNotificationResponse,
	error,
) {
	var response notificationv2.ListNotificationResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieNotificationV2Client) Create(req notificationv2.CreateNotificationRequest) (
	*notificationv2.CreateNotificationResponse,
	error,
) {
	var response notificationv2.CreateNotificationResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieNotificationV2Client) Delete(req notificationv2.DeleteNotificationRequest) (
	*notificationv2.DeleteNotificationResponse,
	error,
) {
	var response notificationv2.DeleteNotificationResponse
	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieNotificationV2Client) Enable(req notificationv2.EnableNotificationRequest) (
	*notificationv2.EnableNotificationResponse,
	error,
) {
	req.Identifier.StatusAction = notificationv2.EnableStatusAction

	var response notificationv2.EnableNotificationResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieNotificationV2Client) Disable(req notificationv2.DisableNotificationRequest) (
	*notificationv2.DisableNotificationResponse,
	error,
) {
	req.Identifier.StatusAction = notificationv2.DisableStatusAction

	var response notificationv2.DisableNotificationResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
