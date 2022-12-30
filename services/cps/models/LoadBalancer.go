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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type LoadBalancer struct {

    /* 负载均衡实例ID (Optional) */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 地域，如cn-east-1 (Optional) */
    Region string `json:"region"`

    /* IP版本，取值ipv4 (Optional) */
    IpAddressType string `json:"ipAddressType"`

    /* 网络类型，取值public (Optional) */
    NetType string `json:"netType"`

    /* 私有网络ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 弹性公网IPID (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 公网IP (Optional) */
    PublicIp string `json:"publicIp"`

    /* 带宽 (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 状态，取值active|inactive (Optional) */
    Status string `json:"status"`

    /* 名称 (Optional) */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`
}