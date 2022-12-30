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

type QueryStatisticsDataGroupByAreaRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* 查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可 (Optional) */
    SubDomain *string `json:"subDomain"`

    /* 需要查询的字段 (Optional) */
    Fields *string `json:"fields"`

    /* 查询的区域，如beijing,shanghai。多个用逗号分隔 (Optional) */
    Area *string `json:"area"`

    /* 查询的运营商，cmcc,cnc,ct，表示移动、联通、电信。多个用逗号分隔 (Optional) */
    Isp *string `json:"isp"`

    /* 是否查询回源统计信息。取值为true和false，默认为false。注意，如果查询回源信息，Fields的取值当前只支持oribandwidth，oripv，oricodestat三个，其余Fields忽略 (Optional) */
    Origin *bool `json:"origin"`

    /* 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional) */
    Period *string `json:"period"`

    /* 分组依据,可选值：[terminal,sdtfrom],如果为空，则只按area/isp进行group (Optional) */
    GroupBy *string `json:"groupBy"`

    /* 查询协议，可选值:[http,https,all],传空默认返回全部协议汇总后的数据 (Optional) */
    Scheme *string `json:"scheme"`

    /* true 代表查询境外数据，默认false查询境内数据 (Optional) */
    Abroad *bool `json:"abroad"`

    /* 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间 (Optional) */
    CacheType *string `json:"cacheType"`

    /* 查询IP类型，可选值:[,ipv4,ipv6],默认查询all (Optional) */
    IpType *string `json:"ipType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryStatisticsDataGroupByAreaRequest(
) *QueryStatisticsDataGroupByAreaRequest {

	return &QueryStatisticsDataGroupByAreaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/vodStatistics:groupByArea",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 * param subDomain: 查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可 (Optional)
 * param fields: 需要查询的字段 (Optional)
 * param area: 查询的区域，如beijing,shanghai。多个用逗号分隔 (Optional)
 * param isp: 查询的运营商，cmcc,cnc,ct，表示移动、联通、电信。多个用逗号分隔 (Optional)
 * param origin: 是否查询回源统计信息。取值为true和false，默认为false。注意，如果查询回源信息，Fields的取值当前只支持oribandwidth，oripv，oricodestat三个，其余Fields忽略 (Optional)
 * param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional)
 * param groupBy: 分组依据,可选值：[terminal,sdtfrom],如果为空，则只按area/isp进行group (Optional)
 * param scheme: 查询协议，可选值:[http,https,all],传空默认返回全部协议汇总后的数据 (Optional)
 * param abroad: true 代表查询境外数据，默认false查询境内数据 (Optional)
 * param cacheType: 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间 (Optional)
 * param ipType: 查询IP类型，可选值:[,ipv4,ipv6],默认查询all (Optional)
 */
func NewQueryStatisticsDataGroupByAreaRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    subDomain *string,
    fields *string,
    area *string,
    isp *string,
    origin *bool,
    period *string,
    groupBy *string,
    scheme *string,
    abroad *bool,
    cacheType *string,
    ipType *string,
) *QueryStatisticsDataGroupByAreaRequest {

    return &QueryStatisticsDataGroupByAreaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/vodStatistics:groupByArea",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        SubDomain: subDomain,
        Fields: fields,
        Area: area,
        Isp: isp,
        Origin: origin,
        Period: period,
        GroupBy: groupBy,
        Scheme: scheme,
        Abroad: abroad,
        CacheType: cacheType,
        IpType: ipType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryStatisticsDataGroupByAreaRequestWithoutParam() *QueryStatisticsDataGroupByAreaRequest {

    return &QueryStatisticsDataGroupByAreaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/vodStatistics:groupByArea",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}
/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}
/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetDomain(domain string) {
    r.Domain = &domain
}
/* param subDomain: 查询泛域名时，指定的子域名列表，多个用逗号分隔。非泛域名时，传入空即可(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetSubDomain(subDomain string) {
    r.SubDomain = &subDomain
}
/* param fields: 需要查询的字段(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetFields(fields string) {
    r.Fields = &fields
}
/* param area: 查询的区域，如beijing,shanghai。多个用逗号分隔(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetArea(area string) {
    r.Area = &area
}
/* param isp: 查询的运营商，cmcc,cnc,ct，表示移动、联通、电信。多个用逗号分隔(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetIsp(isp string) {
    r.Isp = &isp
}
/* param origin: 是否查询回源统计信息。取值为true和false，默认为false。注意，如果查询回源信息，Fields的取值当前只支持oribandwidth，oripv，oricodestat三个，其余Fields忽略(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetOrigin(origin bool) {
    r.Origin = &origin
}
/* param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetPeriod(period string) {
    r.Period = &period
}
/* param groupBy: 分组依据,可选值：[terminal,sdtfrom],如果为空，则只按area/isp进行group(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetGroupBy(groupBy string) {
    r.GroupBy = &groupBy
}
/* param scheme: 查询协议，可选值:[http,https,all],传空默认返回全部协议汇总后的数据(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetScheme(scheme string) {
    r.Scheme = &scheme
}
/* param abroad: true 代表查询境外数据，默认false查询境内数据(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetAbroad(abroad bool) {
    r.Abroad = &abroad
}
/* param cacheType: 查询节点层级，可选值:[all,edge,mid],默认查询all,edge边缘 mid中间(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetCacheType(cacheType string) {
    r.CacheType = &cacheType
}
/* param ipType: 查询IP类型，可选值:[,ipv4,ipv6],默认查询all(Optional) */
func (r *QueryStatisticsDataGroupByAreaRequest) SetIpType(ipType string) {
    r.IpType = &ipType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryStatisticsDataGroupByAreaRequest) GetRegionId() string {
    return ""
}

type QueryStatisticsDataGroupByAreaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryStatisticsDataGroupByAreaResult `json:"result"`
}

type QueryStatisticsDataGroupByAreaResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Domain string `json:"domain"`
    Statistics []cdn.StatisticsWithAreaGroupDetail `json:"statistics"`
}