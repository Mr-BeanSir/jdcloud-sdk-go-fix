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

type QueryDomainLogRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔，取值(hour，day，fiveMin)，不传默认小时。 (Optional) */
    Interval *string `json:"interval"`

    /* 日志类型，取值(log，zip,gz)，不传默认gz。 (Optional) */
    LogType *string `json:"logType"`

    /* 页面大小，默认值10 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 分页页数，默认值1 (Optional) */
    PageNumber *int `json:"pageNumber"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDomainLogRequest(
    domain string,
) *QueryDomainLogRequest {

	return &QueryDomainLogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/log",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param interval: 时间间隔，取值(hour，day，fiveMin)，不传默认小时。 (Optional)
 * param logType: 日志类型，取值(log，zip,gz)，不传默认gz。 (Optional)
 * param pageSize: 页面大小，默认值10 (Optional)
 * param pageNumber: 分页页数，默认值1 (Optional)
 */
func NewQueryDomainLogRequestWithAllParams(
    domain string,
    startTime *string,
    endTime *string,
    interval *string,
    logType *string,
    pageSize *int,
    pageNumber *int,
) *QueryDomainLogRequest {

    return &QueryDomainLogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/log",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        StartTime: startTime,
        EndTime: endTime,
        Interval: interval,
        LogType: logType,
        PageSize: pageSize,
        PageNumber: pageNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDomainLogRequestWithoutParam() *QueryDomainLogRequest {

    return &QueryDomainLogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/log",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryDomainLogRequest) SetDomain(domain string) {
    r.Domain = domain
}
/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryDomainLogRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}
/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryDomainLogRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}
/* param interval: 时间间隔，取值(hour，day，fiveMin)，不传默认小时。(Optional) */
func (r *QueryDomainLogRequest) SetInterval(interval string) {
    r.Interval = &interval
}
/* param logType: 日志类型，取值(log，zip,gz)，不传默认gz。(Optional) */
func (r *QueryDomainLogRequest) SetLogType(logType string) {
    r.LogType = &logType
}
/* param pageSize: 页面大小，默认值10(Optional) */
func (r *QueryDomainLogRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param pageNumber: 分页页数，默认值1(Optional) */
func (r *QueryDomainLogRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDomainLogRequest) GetRegionId() string {
    return ""
}

type QueryDomainLogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDomainLogResult `json:"result"`
}

type QueryDomainLogResult struct {
    Total int `json:"total"`
    PageSize int `json:"pageSize"`
    PageNumber int `json:"pageNumber"`
    Urls []cdn.DomainLog `json:"urls"`
}