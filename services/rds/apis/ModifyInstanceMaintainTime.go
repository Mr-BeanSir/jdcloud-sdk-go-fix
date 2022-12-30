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

type ModifyInstanceMaintainTimeRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 实例的可维护时间段。格式：HH:mm-HH:mm，取值为一个小时整点 (Optional) */
    MaintainTime *string `json:"maintainTime"`

    /* 选择维护周期，可选择一周中的某一天或多天 (Optional) */
    MaintainPeriod []string `json:"maintainPeriod"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceMaintainTimeRequest(
    regionId string,
    instanceId string,
) *ModifyInstanceMaintainTimeRequest {

	return &ModifyInstanceMaintainTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceMaintainTime",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param maintainTime: 实例的可维护时间段。格式：HH:mm-HH:mm，取值为一个小时整点 (Optional)
 * param maintainPeriod: 选择维护周期，可选择一周中的某一天或多天 (Optional)
 */
func NewModifyInstanceMaintainTimeRequestWithAllParams(
    regionId string,
    instanceId string,
    maintainTime *string,
    maintainPeriod []string,
) *ModifyInstanceMaintainTimeRequest {

    return &ModifyInstanceMaintainTimeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceMaintainTime",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        MaintainTime: maintainTime,
        MaintainPeriod: maintainPeriod,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceMaintainTimeRequestWithoutParam() *ModifyInstanceMaintainTimeRequest {

    return &ModifyInstanceMaintainTimeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceMaintainTime",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ModifyInstanceMaintainTimeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *ModifyInstanceMaintainTimeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param maintainTime: 实例的可维护时间段。格式：HH:mm-HH:mm，取值为一个小时整点(Optional) */
func (r *ModifyInstanceMaintainTimeRequest) SetMaintainTime(maintainTime string) {
    r.MaintainTime = &maintainTime
}

/* param maintainPeriod: 选择维护周期，可选择一周中的某一天或多天(Optional) */
func (r *ModifyInstanceMaintainTimeRequest) SetMaintainPeriod(maintainPeriod []string) {
    r.MaintainPeriod = maintainPeriod
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceMaintainTimeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceMaintainTimeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceMaintainTimeResult `json:"result"`
}

type ModifyInstanceMaintainTimeResult struct {
}