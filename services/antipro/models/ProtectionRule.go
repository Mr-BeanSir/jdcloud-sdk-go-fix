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


type ProtectionRule struct {

    /* 防护规则类型. <br>- 0: 默认防护包规则<br>- 1: IP 自定义规则 (Optional) */
    Type int `json:"type"`

    /* 清洗触发值 bps (Optional) */
    CleanThresholdBps int64 `json:"cleanThresholdBps"`

    /* 清洗触发值 pps (Optional) */
    CleanThresholdPps int64 `json:"cleanThresholdPps"`

    /* 虚假源开启 (Optional) */
    SpoofIpEnable int `json:"spoofIpEnable"`

    /* 源新建连接限速开启 (Optional) */
    SrcNewConnLimitEnable int `json:"srcNewConnLimitEnable"`

    /* 源新建连接速率 (Optional) */
    SrcNewConnLimitValue int64 `json:"srcNewConnLimitValue"`

    /* 目的新建连接开启 (Optional) */
    DstNewConnLimitEnable int `json:"dstNewConnLimitEnable"`

    /* 目的新建连接速率 (Optional) */
    DstNewConnLimitValue int64 `json:"dstNewConnLimitValue"`

    /* 报文最小长度 (Optional) */
    DatagramRangeMin int64 `json:"datagramRangeMin"`

    /* 报文最大长度 (Optional) */
    DatagramRangeMax int64 `json:"datagramRangeMax"`

    /* geo 拦截地域列表 (Optional) */
    GeoBlackList []GeoBlack `json:"geoBlackList"`

    /* IP 黑名单列表 (Optional) */
    IpBlackList []string `json:"ipBlackList"`

    /* IP 白名单列表 (Optional) */
    IpWhiteList []string `json:"ipWhiteList"`
}
