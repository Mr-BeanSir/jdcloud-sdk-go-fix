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


type IpResource struct {

    /* 公网 IP 所在区域 (Optional) */
    Region string `json:"region"`

    /* 公网 IP 类型或绑定资源类型. 
<br>- 0: 未知类型
<br>- 1: 弹性公网 IP(IP 为弹性公网 IP, 绑定资源类型未知)
<br>- 10: 弹性公网 IP(IP 为弹性公网 IP, 但未绑定资源)
<br>- 11: 云主机
<br>- 12: 负载均衡
<br>- 13: 原生容器实例
<br>- 14: 原生容器 Pod
<br>- 2: 云物理服务器公网 IP
<br>- 3: Web应用防火墙 IP
<br>- 4: 托管区公网 IP
 (Optional) */
    ResourceType int `json:"resourceType"`

    /* 公网 IP 地址 (Optional) */
    Ip string `json:"ip"`

    /* 带宽上限, 单位 Mbps (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 请求流量清洗触发值, 单位 bps/s (Optional) */
    CleanThresholdBps int64 `json:"cleanThresholdBps"`

    /* 报文请求清洗触发值, 单位 pps/s (Optional) */
    CleanThresholdPps int64 `json:"cleanThresholdPps"`

    /* 黑洞阈值 (Optional) */
    BalckHoleThreshold int64 `json:"balckHoleThreshold"`

    /* 绑定防护包 ID, 为空字符时表示未绑定防护包 (Optional) */
    InstanceId string `json:"instanceId"`

    /* 绑定防护包名称, 为空字符串时表示未绑定防护包 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 安全状态. <br>- 0: 安全 <br>- 1: 清洗 <br>- 2: 黑洞 (Optional) */
    SafeStatus int `json:"safeStatus"`
}
