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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryDomainGroupListRequest struct {

    core.JDCloudRequest

    /* 根据是否共享内存筛选 (Optional) */
    ShareCache *string `json:"shareCache"`

    /* pageNumber (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* pageSize (Optional) */
    PageSize *int `json:"pageSize"`

    /* 根据主域名模糊查询 (Optional) */
    PrimaryDomain *string `json:"primaryDomain"`

    /* 根据域名组模糊查询 (Optional) */
    DomainGroupName *string `json:"domainGroupName"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDomainGroupListRequest(
) *QueryDomainGroupListRequest {

	return &QueryDomainGroupListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domainGroup",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param shareCache: 根据是否共享内存筛选 (Optional)
 * param pageNumber: pageNumber (Optional)
 * param pageSize: pageSize (Optional)
 * param primaryDomain: 根据主域名模糊查询 (Optional)
 * param domainGroupName: 根据域名组模糊查询 (Optional)
 */
func NewQueryDomainGroupListRequestWithAllParams(
    shareCache *string,
    pageNumber *int,
    pageSize *int,
    primaryDomain *string,
    domainGroupName *string,
) *QueryDomainGroupListRequest {

    return &QueryDomainGroupListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ShareCache: shareCache,
        PageNumber: pageNumber,
        PageSize: pageSize,
        PrimaryDomain: primaryDomain,
        DomainGroupName: domainGroupName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDomainGroupListRequestWithoutParam() *QueryDomainGroupListRequest {

    return &QueryDomainGroupListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param shareCache: 根据是否共享内存筛选(Optional) */
func (r *QueryDomainGroupListRequest) SetShareCache(shareCache string) {
    r.ShareCache = &shareCache
}
/* param pageNumber: pageNumber(Optional) */
func (r *QueryDomainGroupListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: pageSize(Optional) */
func (r *QueryDomainGroupListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param primaryDomain: 根据主域名模糊查询(Optional) */
func (r *QueryDomainGroupListRequest) SetPrimaryDomain(primaryDomain string) {
    r.PrimaryDomain = &primaryDomain
}
/* param domainGroupName: 根据域名组模糊查询(Optional) */
func (r *QueryDomainGroupListRequest) SetDomainGroupName(domainGroupName string) {
    r.DomainGroupName = &domainGroupName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDomainGroupListRequest) GetRegionId() string {
    return ""
}

type QueryDomainGroupListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDomainGroupListResult `json:"result"`
}

type QueryDomainGroupListResult struct {
    TotalCount int `json:"totalCount"`
    PageSize int `json:"pageSize"`
    PageNumber int `json:"pageNumber"`
    DomainGroups []cdn.DomainGroupItem `json:"domainGroups"`
}