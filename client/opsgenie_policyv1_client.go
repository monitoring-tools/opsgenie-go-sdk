package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/policyv1"
)

type OpsGeniePolicyV1Client struct {
	RestClient
}

func (cli *OpsGeniePolicyV1Client) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGeniePolicyV1Client) Create(req policyv1.CreatePolicyRequest) (*policyv1.CreatePolicyResponse, error) {
	var response policyv1.CreatePolicyResponse
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Get(req policyv1.GetPolicyRequest) (*policyv1.GetPolicyResponse, error) {
	var response policyv1.GetPolicyResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Update(req policyv1.UpdatePolicyRequest) (*policyv1.UpdatePolicyResponse, error) {
	var response policyv1.UpdatePolicyResponse
	err := cli.sendPutRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Delete(req policyv1.DeletePolicyRequest) (*policyv1.DeletePolicyResponse, error) {
	var response policyv1.DeletePolicyResponse
	err := cli.sendDeleteRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Disable(req policyv1.DisablePolicyRequest) (*policyv1.DisablePolicyResponse, error) {
	var response policyv1.DisablePolicyResponse
	req.StatusAction = policyv1.DisableAction
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) Enable(req policyv1.EnablePolicyRequest) (*policyv1.EnablePolicyResponse, error) {
	var response policyv1.EnablePolicyResponse
	req.StatusAction = policyv1.EnableAction
	err := cli.sendPostRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cli *OpsGeniePolicyV1Client) List(req policyv1.ListPolicyRequest) (*policyv1.ListPolicyResponse, error) {
	var response policyv1.ListPolicyResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}