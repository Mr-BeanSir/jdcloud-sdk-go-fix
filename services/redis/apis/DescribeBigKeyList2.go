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

type DescribeBigKeyList2Request struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBigKeyList2Request(
    regionId string,
    cacheInstanceId string,
) *DescribeBigKeyList2Request {

	return &DescribeBigKeyList2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKeyAnalysisResultList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 */
func NewDescribeBigKeyList2RequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    pageNumber *int,
    pageSize *int,
) *DescribeBigKeyList2Request {

    return &DescribeBigKeyList2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKeyAnalysisResultList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBigKeyList2RequestWithoutParam() *DescribeBigKeyList2Request {

    return &DescribeBigKeyList2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/bigKeyAnalysisResultList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *DescribeBigKeyList2Request) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *DescribeBigKeyList2Request) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}
/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeBigKeyList2Request) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeBigKeyList2Request) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBigKeyList2Request) GetRegionId() string {
    return r.RegionId
}

type DescribeBigKeyList2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBigKeyList2Result `json:"result"`
}

type DescribeBigKeyList2Result struct {
    CacheAnalysisList []redis.ClusterCacheAnalysis `json:"cacheAnalysisList"`
    TotalCount int `json:"totalCount"`
}