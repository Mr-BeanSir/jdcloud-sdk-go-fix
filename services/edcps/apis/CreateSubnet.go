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
    edcps "github.com/jdcloud-api/jdcloud-sdk-go/services/edcps/models"
)

type CreateSubnetRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 子网配置  */
    SubnetSpec *edcps.SubnetSpec `json:"subnetSpec"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param subnetSpec: 子网配置 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSubnetRequest(
    regionId string,
    subnetSpec *edcps.SubnetSpec,
) *CreateSubnetRequest {

	return &CreateSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnets",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SubnetSpec: subnetSpec,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param subnetSpec: 子网配置 (Required)
 */
func NewCreateSubnetRequestWithAllParams(
    regionId string,
    clientToken *string,
    subnetSpec *edcps.SubnetSpec,
) *CreateSubnetRequest {

    return &CreateSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientToken: clientToken,
        SubnetSpec: subnetSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSubnetRequestWithoutParam() *CreateSubnetRequest {

    return &CreateSubnetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *CreateSubnetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *CreateSubnetRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param subnetSpec: 子网配置(Required) */
func (r *CreateSubnetRequest) SetSubnetSpec(subnetSpec *edcps.SubnetSpec) {
    r.SubnetSpec = subnetSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSubnetRequest) GetRegionId() string {
    return r.RegionId
}

type CreateSubnetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSubnetResult `json:"result"`
}

type CreateSubnetResult struct {
    SubnetId string `json:"subnetId"`
}