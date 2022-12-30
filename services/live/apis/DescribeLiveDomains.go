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
    live "github.com/jdcloud-api/jdcloud-sdk-go/services/live/models"
)

type DescribeLiveDomainsRequest struct {

    core.JDCloudRequest

    /* 页码
- 取值范围[1, 100000]
 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小
- 取值范围[10, 100]
 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 推流域名
- 目前仅支持精确查询
- 为空时,查询用户所有直播域名
 (Optional) */
    PublishDomain *string `json:"publishDomain"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveDomainsRequest(
) *DescribeLiveDomainsRequest {

	return &DescribeLiveDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNum: 页码
- 取值范围[1, 100000]
 (Optional)
 * param pageSize: 分页大小
- 取值范围[10, 100]
 (Optional)
 * param publishDomain: 推流域名
- 目前仅支持精确查询
- 为空时,查询用户所有直播域名
 (Optional)
 */
func NewDescribeLiveDomainsRequestWithAllParams(
    pageNum *int,
    pageSize *int,
    publishDomain *string,
) *DescribeLiveDomainsRequest {

    return &DescribeLiveDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        PublishDomain: publishDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveDomainsRequestWithoutParam() *DescribeLiveDomainsRequest {

    return &DescribeLiveDomainsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 页码
- 取值范围[1, 100000]
(Optional) */
func (r *DescribeLiveDomainsRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 分页大小
- 取值范围[10, 100]
(Optional) */
func (r *DescribeLiveDomainsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param publishDomain: 推流域名
- 目前仅支持精确查询
- 为空时,查询用户所有直播域名
(Optional) */
func (r *DescribeLiveDomainsRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = &publishDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveDomainsRequest) GetRegionId() string {
    return ""
}

type DescribeLiveDomainsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveDomainsResult `json:"result"`
}

type DescribeLiveDomainsResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    DomainDetails []live.DomainDetails `json:"domainDetails"`
}