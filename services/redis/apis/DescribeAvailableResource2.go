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
    redis "github.com/jdcloud-api/jdcloud-sdk-go/services/redis/models"
)

type DescribeAvailableResource2Request struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAvailableResource2Request(
    regionId string,
) *DescribeAvailableResource2Request {

	return &DescribeAvailableResource2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/availableResource2",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 */
func NewDescribeAvailableResource2RequestWithAllParams(
    regionId string,
) *DescribeAvailableResource2Request {

    return &DescribeAvailableResource2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availableResource2",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAvailableResource2RequestWithoutParam() *DescribeAvailableResource2Request {

    return &DescribeAvailableResource2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availableResource2",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *DescribeAvailableResource2Request) SetRegionId(regionId string) {
    r.RegionId = regionId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAvailableResource2Request) GetRegionId() string {
    return r.RegionId
}

type DescribeAvailableResource2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAvailableResource2Result `json:"result"`
}

type DescribeAvailableResource2Result struct {
    AvailableResources []redis.AvailableResource `json:"availableResources"`
}