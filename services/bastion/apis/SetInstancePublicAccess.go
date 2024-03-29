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

type SetInstancePublicAccessRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 堡垒机id  */
    InstanceId string `json:"instanceId"`

    /*   */
    Status bool `json:"status"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 堡垒机id (Required)
 * param status:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetInstancePublicAccessRequest(
    regionId string,
    instanceId string,
    status bool,
) *SetInstancePublicAccessRequest {

	return &SetInstancePublicAccessRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceId}/publicAccess",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Status: status,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 堡垒机id (Required)
 * param status:  (Required)
 */
func NewSetInstancePublicAccessRequestWithAllParams(
    regionId string,
    instanceId string,
    status bool,
) *SetInstancePublicAccessRequest {

    return &SetInstancePublicAccessRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceId}/publicAccess",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetInstancePublicAccessRequestWithoutParam() *SetInstancePublicAccessRequest {

    return &SetInstancePublicAccessRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceId}/publicAccess",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *SetInstancePublicAccessRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 堡垒机id(Required) */
func (r *SetInstancePublicAccessRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param status: (Required) */
func (r *SetInstancePublicAccessRequest) SetStatus(status bool) {
    r.Status = status
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetInstancePublicAccessRequest) GetRegionId() string {
    return r.RegionId
}

type SetInstancePublicAccessResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetInstancePublicAccessResult `json:"result"`
}

type SetInstancePublicAccessResult struct {
    Result bool `json:"result"`
}