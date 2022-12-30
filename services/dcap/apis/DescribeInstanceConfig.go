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

type DescribeInstanceConfigRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例 ID  */
    InstanceId string `json:"instanceId"`

    /*   */
    AppUuid string `json:"appUuid"`

    /*   */
    AppAddr string `json:"appAddr"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 * param appUuid:  (Required)
 * param appAddr:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceConfigRequest(
    regionId string,
    instanceId string,
    appUuid string,
    appAddr string,
) *DescribeInstanceConfigRequest {

	return &DescribeInstanceConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:config",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AppUuid: appUuid,
        AppAddr: appAddr,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 * param appUuid:  (Required)
 * param appAddr:  (Required)
 */
func NewDescribeInstanceConfigRequestWithAllParams(
    regionId string,
    instanceId string,
    appUuid string,
    appAddr string,
) *DescribeInstanceConfigRequest {

    return &DescribeInstanceConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        AppUuid: appUuid,
        AppAddr: appAddr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceConfigRequestWithoutParam() *DescribeInstanceConfigRequest {

    return &DescribeInstanceConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeInstanceConfigRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例 ID(Required) */
func (r *DescribeInstanceConfigRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param appUuid: (Required) */
func (r *DescribeInstanceConfigRequest) SetAppUuid(appUuid string) {
    r.AppUuid = appUuid
}

/* param appAddr: (Required) */
func (r *DescribeInstanceConfigRequest) SetAppAddr(appAddr string) {
    r.AppAddr = appAddr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceConfigRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceConfigResult `json:"result"`
}

type DescribeInstanceConfigResult struct {
    Configuration string `json:"configuration"`
}