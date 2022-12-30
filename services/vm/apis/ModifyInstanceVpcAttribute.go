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

type ModifyInstanceVpcAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 云主机ID。  */
    InstanceId string `json:"instanceId"`

    /* 子网Id。  */
    SubnetId string `json:"subnetId"`

    /* `true`: 分配IPV6地址。
`false`: 不分配IPV6地址。
 (Optional) */
    AssignIpv6 *bool `json:"assignIpv6"`

    /* Ipv4地址。
不变更 `vpc` 及子网时必须指定Ipv4地址
 (Optional) */
    PrivateIpAddress *string `json:"privateIpAddress"`

    /* 安全组列表。
更换 `vpc` 时必须指定新的安全组。
不更换 `vpc` 时不可以指定安全组。
 (Optional) */
    SecurityGroups []string `json:"securityGroups"`

}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 * param subnetId: 子网Id。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceVpcAttributeRequest(
    regionId string,
    instanceId string,
    subnetId string,
) *ModifyInstanceVpcAttributeRequest {

	return &ModifyInstanceVpcAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceVpcAttribute",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 * param subnetId: 子网Id。 (Required)
 * param assignIpv6: `true`: 分配IPV6地址。
`false`: 不分配IPV6地址。
 (Optional)
 * param privateIpAddress: Ipv4地址。
不变更 `vpc` 及子网时必须指定Ipv4地址
 (Optional)
 * param securityGroups: 安全组列表。
更换 `vpc` 时必须指定新的安全组。
不更换 `vpc` 时不可以指定安全组。
 (Optional)
 */
func NewModifyInstanceVpcAttributeRequestWithAllParams(
    regionId string,
    instanceId string,
    subnetId string,
    assignIpv6 *bool,
    privateIpAddress *string,
    securityGroups []string,
) *ModifyInstanceVpcAttributeRequest {

    return &ModifyInstanceVpcAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceVpcAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        SubnetId: subnetId,
        AssignIpv6: assignIpv6,
        PrivateIpAddress: privateIpAddress,
        SecurityGroups: securityGroups,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceVpcAttributeRequestWithoutParam() *ModifyInstanceVpcAttributeRequest {

    return &ModifyInstanceVpcAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceVpcAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *ModifyInstanceVpcAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云主机ID。(Required) */
func (r *ModifyInstanceVpcAttributeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param subnetId: 子网Id。(Required) */
func (r *ModifyInstanceVpcAttributeRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

/* param assignIpv6: `true`: 分配IPV6地址。
`false`: 不分配IPV6地址。
(Optional) */
func (r *ModifyInstanceVpcAttributeRequest) SetAssignIpv6(assignIpv6 bool) {
    r.AssignIpv6 = &assignIpv6
}

/* param privateIpAddress: Ipv4地址。
不变更 `vpc` 及子网时必须指定Ipv4地址
(Optional) */
func (r *ModifyInstanceVpcAttributeRequest) SetPrivateIpAddress(privateIpAddress string) {
    r.PrivateIpAddress = &privateIpAddress
}

/* param securityGroups: 安全组列表。
更换 `vpc` 时必须指定新的安全组。
不更换 `vpc` 时不可以指定安全组。
(Optional) */
func (r *ModifyInstanceVpcAttributeRequest) SetSecurityGroups(securityGroups []string) {
    r.SecurityGroups = securityGroups
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceVpcAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceVpcAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceVpcAttributeResult `json:"result"`
}

type ModifyInstanceVpcAttributeResult struct {
}