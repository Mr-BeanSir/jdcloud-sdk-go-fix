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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type ModifyWhiteListRuleOfForwardRuleRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 转发规则 Id  */
    ForwardRuleId string `json:"forwardRuleId"`

    /* 修改转发规则的黑名单规则请求参数  */
    ModifySpec *ipanti.ModifyWhiteListRuleOfForwardRuleSpec `json:"modifySpec"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 * param modifySpec: 修改转发规则的黑名单规则请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyWhiteListRuleOfForwardRuleRequest(
    regionId string,
    instanceId string,
    forwardRuleId string,
    modifySpec *ipanti.ModifyWhiteListRuleOfForwardRuleSpec,
) *ModifyWhiteListRuleOfForwardRuleRequest {

	return &ModifyWhiteListRuleOfForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}/forwardWhiteListRule",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
        ModifySpec: modifySpec,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 * param modifySpec: 修改转发规则的黑名单规则请求参数 (Required)
 */
func NewModifyWhiteListRuleOfForwardRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    forwardRuleId string,
    modifySpec *ipanti.ModifyWhiteListRuleOfForwardRuleSpec,
) *ModifyWhiteListRuleOfForwardRuleRequest {

    return &ModifyWhiteListRuleOfForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}/forwardWhiteListRule",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
        ModifySpec: modifySpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyWhiteListRuleOfForwardRuleRequestWithoutParam() *ModifyWhiteListRuleOfForwardRuleRequest {

    return &ModifyWhiteListRuleOfForwardRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}/forwardWhiteListRule",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *ModifyWhiteListRuleOfForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *ModifyWhiteListRuleOfForwardRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param forwardRuleId: 转发规则 Id(Required) */
func (r *ModifyWhiteListRuleOfForwardRuleRequest) SetForwardRuleId(forwardRuleId string) {
    r.ForwardRuleId = forwardRuleId
}

/* param modifySpec: 修改转发规则的黑名单规则请求参数(Required) */
func (r *ModifyWhiteListRuleOfForwardRuleRequest) SetModifySpec(modifySpec *ipanti.ModifyWhiteListRuleOfForwardRuleSpec) {
    r.ModifySpec = modifySpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyWhiteListRuleOfForwardRuleRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyWhiteListRuleOfForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyWhiteListRuleOfForwardRuleResult `json:"result"`
}

type ModifyWhiteListRuleOfForwardRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}