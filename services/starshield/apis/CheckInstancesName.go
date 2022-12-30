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
)

type CheckInstancesNameRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceName: 实例名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckInstancesNameRequest(
    regionId string,
    instanceName string,
) *CheckInstancesNameRequest {

	return &CheckInstancesNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:checkName",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceName: instanceName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceName: 实例名称 (Required)
 */
func NewCheckInstancesNameRequestWithAllParams(
    regionId string,
    instanceName string,
) *CheckInstancesNameRequest {

    return &CheckInstancesNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:checkName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceName: instanceName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckInstancesNameRequestWithoutParam() *CheckInstancesNameRequest {

    return &CheckInstancesNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:checkName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CheckInstancesNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceName: 实例名称(Required) */
func (r *CheckInstancesNameRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckInstancesNameRequest) GetRegionId() string {
    return r.RegionId
}

type CheckInstancesNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckInstancesNameResult `json:"result"`
}

type CheckInstancesNameResult struct {
}