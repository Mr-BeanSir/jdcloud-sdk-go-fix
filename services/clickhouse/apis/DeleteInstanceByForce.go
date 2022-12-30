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

type DeleteInstanceByForceRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 发出请求的运营后台的ERP账号  */
    RequestUser string `json:"requestUser"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param requestUser: 发出请求的运营后台的ERP账号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteInstanceByForceRequest(
    regionId string,
    instanceId string,
    requestUser string,
) *DeleteInstanceByForceRequest {

	return &DeleteInstanceByForceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:deleteInstanceByForce",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        RequestUser: requestUser,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param requestUser: 发出请求的运营后台的ERP账号 (Required)
 */
func NewDeleteInstanceByForceRequestWithAllParams(
    regionId string,
    instanceId string,
    requestUser string,
) *DeleteInstanceByForceRequest {

    return &DeleteInstanceByForceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:deleteInstanceByForce",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        RequestUser: requestUser,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteInstanceByForceRequestWithoutParam() *DeleteInstanceByForceRequest {

    return &DeleteInstanceByForceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:deleteInstanceByForce",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DeleteInstanceByForceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *DeleteInstanceByForceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param requestUser: 发出请求的运营后台的ERP账号(Required) */
func (r *DeleteInstanceByForceRequest) SetRequestUser(requestUser string) {
    r.RequestUser = requestUser
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteInstanceByForceRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteInstanceByForceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteInstanceByForceResult `json:"result"`
}

type DeleteInstanceByForceResult struct {
}