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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeInstanceMonitorInfoRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 云物理服务器ID  */
    InstanceId string `json:"instanceId"`

    /* 开始时间的时间戳，格式：1562915166551 (Optional) */
    StartTime *int `json:"startTime"`

    /* 结束时间的时间戳，格式：1562915166551 (Optional) */
    EndTime *int `json:"endTime"`

    /* metric - 监控指标，精确匹配，支持多个，具体如下<br/>
cps.cpu.util - CPU使用率<br/>
cps.memory.util - 内存使用率<br/>
cps.memory.used - 内存使用量<br/>
cps.disk.used - 磁盘使用量<br/>
cps.disk.util - 磁盘使用率<br/>
cps.disk.bytes.read - 磁盘读流量<br/>
cps.disk.bytes.write - 磁盘写流量<br/>
cps.disk.counts.read - 磁盘读IOPS<br/>
cps.disk.counts.write - 磁盘写IOPS<br/>
cps.network.bytes.ingress - 网卡进流量<br/>
cps.network.bytes.egress - 网卡出流量<br/>
cps.network.packets.ingress - 网络进包量<br/>
cps.network.packets.egress - 网络出包量<br/>
cps.avg.load1 - CPU平均负载1min<br/>
cps.avg.load5 - CPU平均负载5min<br/>
cps.avg.load15 - CPU平均负载15min<br/>
cps.tcp.connect.total - TCP总连接数<br/>
cps.tcp.connect.established - TCP正常连接数<br/>
cps.process.total - 总进程数
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: 云物理服务器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceMonitorInfoRequest(
    regionId string,
    instanceId string,
) *DescribeInstanceMonitorInfoRequest {

	return &DescribeInstanceMonitorInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/monitor",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: 云物理服务器ID (Required)
 * param startTime: 开始时间的时间戳，格式：1562915166551 (Optional)
 * param endTime: 结束时间的时间戳，格式：1562915166551 (Optional)
 * param filters: metric - 监控指标，精确匹配，支持多个，具体如下<br/>
cps.cpu.util - CPU使用率<br/>
cps.memory.util - 内存使用率<br/>
cps.memory.used - 内存使用量<br/>
cps.disk.used - 磁盘使用量<br/>
cps.disk.util - 磁盘使用率<br/>
cps.disk.bytes.read - 磁盘读流量<br/>
cps.disk.bytes.write - 磁盘写流量<br/>
cps.disk.counts.read - 磁盘读IOPS<br/>
cps.disk.counts.write - 磁盘写IOPS<br/>
cps.network.bytes.ingress - 网卡进流量<br/>
cps.network.bytes.egress - 网卡出流量<br/>
cps.network.packets.ingress - 网络进包量<br/>
cps.network.packets.egress - 网络出包量<br/>
cps.avg.load1 - CPU平均负载1min<br/>
cps.avg.load5 - CPU平均负载5min<br/>
cps.avg.load15 - CPU平均负载15min<br/>
cps.tcp.connect.total - TCP总连接数<br/>
cps.tcp.connect.established - TCP正常连接数<br/>
cps.process.total - 总进程数
 (Optional)
 */
func NewDescribeInstanceMonitorInfoRequestWithAllParams(
    regionId string,
    instanceId string,
    startTime *int,
    endTime *int,
    filters []common.Filter,
) *DescribeInstanceMonitorInfoRequest {

    return &DescribeInstanceMonitorInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/monitor",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        StartTime: startTime,
        EndTime: endTime,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceMonitorInfoRequestWithoutParam() *DescribeInstanceMonitorInfoRequest {

    return &DescribeInstanceMonitorInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/monitor",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeInstanceMonitorInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云物理服务器ID(Required) */
func (r *DescribeInstanceMonitorInfoRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param startTime: 开始时间的时间戳，格式：1562915166551(Optional) */
func (r *DescribeInstanceMonitorInfoRequest) SetStartTime(startTime int) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间的时间戳，格式：1562915166551(Optional) */
func (r *DescribeInstanceMonitorInfoRequest) SetEndTime(endTime int) {
    r.EndTime = &endTime
}

/* param filters: metric - 监控指标，精确匹配，支持多个，具体如下<br/>
cps.cpu.util - CPU使用率<br/>
cps.memory.util - 内存使用率<br/>
cps.memory.used - 内存使用量<br/>
cps.disk.used - 磁盘使用量<br/>
cps.disk.util - 磁盘使用率<br/>
cps.disk.bytes.read - 磁盘读流量<br/>
cps.disk.bytes.write - 磁盘写流量<br/>
cps.disk.counts.read - 磁盘读IOPS<br/>
cps.disk.counts.write - 磁盘写IOPS<br/>
cps.network.bytes.ingress - 网卡进流量<br/>
cps.network.bytes.egress - 网卡出流量<br/>
cps.network.packets.ingress - 网络进包量<br/>
cps.network.packets.egress - 网络出包量<br/>
cps.avg.load1 - CPU平均负载1min<br/>
cps.avg.load5 - CPU平均负载5min<br/>
cps.avg.load15 - CPU平均负载15min<br/>
cps.tcp.connect.total - TCP总连接数<br/>
cps.tcp.connect.established - TCP正常连接数<br/>
cps.process.total - 总进程数
(Optional) */
func (r *DescribeInstanceMonitorInfoRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceMonitorInfoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceMonitorInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceMonitorInfoResult `json:"result"`
}

type DescribeInstanceMonitorInfoResult struct {
    MetricDatas []cps.MetricData `json:"metricDatas"`
}