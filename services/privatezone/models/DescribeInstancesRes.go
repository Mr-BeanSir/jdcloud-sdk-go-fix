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

package models


type DescribeInstancesRes struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 套餐类型 (Optional) */
    PackType string `json:"packType"`

    /* zone数量 (Optional) */
    ZoneNum int `json:"zoneNum"`

    /* 绑定vpc数量 (Optional) */
    BindVpcNum int `json:"bindVpcNum"`

    /* zone级别 (Optional) */
    ZoneLevel int `json:"zoneLevel"`

    /* 解析记录数量 (Optional) */
    RrNum int `json:"rrNum"`

    /* 域名等级 (Optional) */
    DomainLevel int `json:"domainLevel"`

    /* 导出解析记录权限 (Optional) */
    RrAuthorExport bool `json:"rrAuthorExport"`

    /* 购买时长 (Optional) */
    Duration int `json:"duration"`

    /* 时长单位（YEAR->年） (Optional) */
    DurationUnit string `json:"durationUnit"`

    /* 到期时间, UTC时间格式，例如2017-11-10T23:00:00Z (Optional) */
    ExpireTime string `json:"expireTime"`

    /* 计费状态（NORMAL->正常 EXPIRE->到期 DELETED->已删除） (Optional) */
    ChargeStutas string `json:"chargeStutas"`

    /* 已使用zone数量 (Optional) */
    UsedZoneNum int `json:"usedZoneNum"`
}
