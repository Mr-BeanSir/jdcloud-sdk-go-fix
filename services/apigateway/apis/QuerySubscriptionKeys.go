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
    apigateway "github.com/jdcloud-api/jdcloud-sdk-go/services/apigateway/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type QuerySubscriptionKeysRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* description - 名称，模糊匹配
subscriptionKeyId - subscriptionKeyId，精确匹配
orderBy - 排序类型 desc asc
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQuerySubscriptionKeysRequest(
    regionId string,
) *QuerySubscriptionKeysRequest {

	return &QuerySubscriptionKeysRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subscriptionKeys",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: description - 名称，模糊匹配
subscriptionKeyId - subscriptionKeyId，精确匹配
orderBy - 排序类型 desc asc
 (Optional)
 */
func NewQuerySubscriptionKeysRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *QuerySubscriptionKeysRequest {

    return &QuerySubscriptionKeysRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subscriptionKeys",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQuerySubscriptionKeysRequestWithoutParam() *QuerySubscriptionKeysRequest {

    return &QuerySubscriptionKeysRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subscriptionKeys",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QuerySubscriptionKeysRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *QuerySubscriptionKeysRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *QuerySubscriptionKeysRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: description - 名称，模糊匹配
subscriptionKeyId - subscriptionKeyId，精确匹配
orderBy - 排序类型 desc asc
(Optional) */
func (r *QuerySubscriptionKeysRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QuerySubscriptionKeysRequest) GetRegionId() string {
    return r.RegionId
}

type QuerySubscriptionKeysResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QuerySubscriptionKeysResult `json:"result"`
}

type QuerySubscriptionKeysResult struct {
    SubscriptionKeys []apigateway.SubscriptionKey `json:"subscriptionKeys"`
    TotalCount int `json:"totalCount"`
}