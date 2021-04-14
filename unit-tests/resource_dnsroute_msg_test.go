// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package test

import (
	"terraform-provider-vmp/unit-tests/helper"
	"terraform-provider-vmp/unit-tests/model"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUT_MasterServerGroup_basic(t *testing.T) {
	t.Parallel()

	// // Test cases for storage account name conversion logic
	tc, err := getTestCases()
	if err != nil {
		t.Errorf("Reading credential_ucc.json file error:%s", err)
	}

	for _, input := range *tc {
		// Specify the test case folder and "-var" options
		tfOptions := &terraform.Options{
			TerraformDir: "../examples/resources/dns_route",
			Vars: map[string]interface{}{
				"credentials": map[string]interface{}{
					"api_token":         input.Credential.ApiToken,
					"ids_client_secret": input.Credential.IdsClientSecret,
					"ids_client_id":     input.Credential.IdsClientID,
					"ids_scope":         input.Credential.IdsScope,
					"api_address":       input.Credential.ApiAddress,
					"ids_address":       input.Credential.IdsAddress,
				},
			},
		}

		// Construct the terraform options with default retryable errors to handle the most common
		// retryable errors in terraform testing.
		terraformOptions := terraform.WithDefaultRetryableErrors(t, tfOptions)
		// At the end of the test, run `terraform destroy` to clean up any resources that were created.
		defer terraform.Destroy(t, terraformOptions)

		// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
		terraform.InitAndApply(t, terraformOptions)

	}
}

func getTestCases() (*map[string]model.MasterServerGroupRequest, error) {
	tc := make(map[string]model.MasterServerGroupRequest)
	credential := model.Credentials{}
	err := helper.ReadCredentialJsonfile("credential_ucc.json", &credential)
	if err != nil {
		return nil, err
	}

	request := model.MasterServerGroupRequest{
		Credential: credential,
	}
	tc["terratest.testing.vmp.msg1"] = request

	return &tc, nil
}