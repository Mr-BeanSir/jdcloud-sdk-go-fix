// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    cloudsign "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudsign/models"
)

type SignContractRequest struct {

    core.JDCloudRequest

    /*   */
    ContractSpec *cloudsign.ContractSpec `json:"contractSpec"`
}

/*
 * param contractSpec:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSignContractRequest(
    contractSpec *cloudsign.ContractSpec,
) *SignContractRequest {

	return &SignContractRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/contract",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ContractSpec: contractSpec,
	}
}

/*
 * param contractSpec:  (Required)
 */
func NewSignContractRequestWithAllParams(
    contractSpec *cloudsign.ContractSpec,
) *SignContractRequest {

    return &SignContractRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/contract",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ContractSpec: contractSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSignContractRequestWithoutParam() *SignContractRequest {

    return &SignContractRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/contract",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param contractSpec: (Required) */
func (r *SignContractRequest) SetContractSpec(contractSpec *cloudsign.ContractSpec) {
    r.ContractSpec = contractSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SignContractRequest) GetRegionId() string {
    return ""
}

type SignContractResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SignContractResult `json:"result"`
}

type SignContractResult struct {
    ContractId string `json:"contractId"`
    ContractContent string `json:"contractContent"`
}