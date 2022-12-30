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

type GetDomainListRequest struct {

    core.JDCloudRequest

    /* 根据关键字进行模糊匹配 (Optional) */
    KeyWord *string `json:"keyWord"`

    /* pageNumber,默认值1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* pageSize,最大值50,默认值20 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 根据域名状态查询, 可选值[offline, online, configuring, auditing, audit_reject] (Optional) */
    Status *string `json:"status"`

    /* 域名类型，(web:静态小文件，download:大文件加速，vod:视频加速，live:直播加速),不传查所有 (Optional) */
    Type *string `json:"type"`

    /* 加速区域，(mainLand:中国大陆，nonMainLand:海外加港澳台，all:全球),不传为全球 (Optional) */
    AccelerateRegion *string `json:"accelerateRegion"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDomainListRequest(
) *GetDomainListRequest {

	return &GetDomainListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param keyWord: 根据关键字进行模糊匹配 (Optional)
 * param pageNumber: pageNumber,默认值1 (Optional)
 * param pageSize: pageSize,最大值50,默认值20 (Optional)
 * param status: 根据域名状态查询, 可选值[offline, online, configuring, auditing, audit_reject] (Optional)
 * param type_: 域名类型，(web:静态小文件，download:大文件加速，vod:视频加速，live:直播加速),不传查所有 (Optional)
 * param accelerateRegion: 加速区域，(mainLand:中国大陆，nonMainLand:海外加港澳台，all:全球),不传为全球 (Optional)
 */
func NewGetDomainListRequestWithAllParams(
    keyWord *string,
    pageNumber *int,
    pageSize *int,
    status *string,
    type_ *string,
    accelerateRegion *string,
) *GetDomainListRequest {

    return &GetDomainListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        KeyWord: keyWord,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Status: status,
        Type: type_,
        AccelerateRegion: accelerateRegion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDomainListRequestWithoutParam() *GetDomainListRequest {

    return &GetDomainListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param keyWord: 根据关键字进行模糊匹配(Optional) */
func (r *GetDomainListRequest) SetKeyWord(keyWord string) {
    r.KeyWord = &keyWord
}
/* param pageNumber: pageNumber,默认值1(Optional) */
func (r *GetDomainListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: pageSize,最大值50,默认值20(Optional) */
func (r *GetDomainListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param status: 根据域名状态查询, 可选值[offline, online, configuring, auditing, audit_reject](Optional) */
func (r *GetDomainListRequest) SetStatus(status string) {
    r.Status = &status
}
/* param type_: 域名类型，(web:静态小文件，download:大文件加速，vod:视频加速，live:直播加速),不传查所有(Optional) */
func (r *GetDomainListRequest) SetType(type_ string) {
    r.Type = &type_
}
/* param accelerateRegion: 加速区域，(mainLand:中国大陆，nonMainLand:海外加港澳台，all:全球),不传为全球(Optional) */
func (r *GetDomainListRequest) SetAccelerateRegion(accelerateRegion string) {
    r.AccelerateRegion = &accelerateRegion
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDomainListRequest) GetRegionId() string {
    return ""
}

type GetDomainListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDomainListResult `json:"result"`
}

type GetDomainListResult struct {
    TotalCount int `json:"totalCount"`
    PageSize int `json:"pageSize"`
    PageNumber int `json:"pageNumber"`
    Domains []cdn.ListDomainItem `json:"domains"`
}