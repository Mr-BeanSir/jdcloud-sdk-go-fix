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

type QueryPurgeTaskRequest struct {

    core.JDCloudRequest

    /* url (Optional) */
    Url *string `json:"url"`

    /* 查询状态 1:进行中 2:已完成 (Optional) */
    Status *int `json:"status"`

    /* 同url,系统内部url对应id（url和file_id同时存在时以url为准） (Optional) */
    FileId *string `json:"fileId"`

    /* 页码数,最小为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页大小,默认10 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 域名 (Optional) */
    Domain *string `json:"domain"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryPurgeTaskRequest(
) *QueryPurgeTaskRequest {

	return &QueryPurgeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/purgeTask:query",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param url: url (Optional)
 * param status: 查询状态 1:进行中 2:已完成 (Optional)
 * param fileId: 同url,系统内部url对应id（url和file_id同时存在时以url为准） (Optional)
 * param pageNumber: 页码数,最小为1 (Optional)
 * param pageSize: 每页大小,默认10 (Optional)
 * param domain: 域名 (Optional)
 */
func NewQueryPurgeTaskRequestWithAllParams(
    url *string,
    status *int,
    fileId *string,
    pageNumber *int,
    pageSize *int,
    domain *string,
) *QueryPurgeTaskRequest {

    return &QueryPurgeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/purgeTask:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Url: url,
        Status: status,
        FileId: fileId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Domain: domain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryPurgeTaskRequestWithoutParam() *QueryPurgeTaskRequest {

    return &QueryPurgeTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/purgeTask:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param url: url(Optional) */
func (r *QueryPurgeTaskRequest) SetUrl(url string) {
    r.Url = &url
}
/* param status: 查询状态 1:进行中 2:已完成(Optional) */
func (r *QueryPurgeTaskRequest) SetStatus(status int) {
    r.Status = &status
}
/* param fileId: 同url,系统内部url对应id（url和file_id同时存在时以url为准）(Optional) */
func (r *QueryPurgeTaskRequest) SetFileId(fileId string) {
    r.FileId = &fileId
}
/* param pageNumber: 页码数,最小为1(Optional) */
func (r *QueryPurgeTaskRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 每页大小,默认10(Optional) */
func (r *QueryPurgeTaskRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param domain: 域名(Optional) */
func (r *QueryPurgeTaskRequest) SetDomain(domain string) {
    r.Domain = &domain
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryPurgeTaskRequest) GetRegionId() string {
    return ""
}

type QueryPurgeTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryPurgeTaskResult `json:"result"`
}

type QueryPurgeTaskResult struct {
    TotalNumber int `json:"totalNumber"`
    TotalPage int `json:"totalPage"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TaskList []cdn.PurgeTaskInfo `json:"taskList"`
}