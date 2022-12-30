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


type CCProtectionConfig struct {

    /* 网站规则的 CC 防护状态, 0: 关闭, 1: 开启 (Optional) */
    Enable int `json:"enable"`

    /* 观察者模式, 0: 关闭, 1: 开启 (Optional) */
    ObserverMode int `json:"observerMode"`

    /* 防护等级, 0: 正常, 1: 宽松, 2: 紧急, 3: 自定义 (Optional) */
    Level int `json:"level"`

    /* HTTP 请求数阈值 (Optional) */
    CcThreshold int64 `json:"ccThreshold"`

    /* Host 的防护阈值 (Optional) */
    HostQps int64 `json:"hostQps"`

    /* Host + Url 的防护阈值 (Optional) */
    HostUrlQps int64 `json:"hostUrlQps"`

    /* 每个源 IP 对 Host 的防护阈值 (Optional) */
    IpHostQps int64 `json:"ipHostQps"`

    /* 每个源 IP 对 Host + Url 的防护阈值 (Optional) */
    IpHostUrlQps int64 `json:"ipHostUrlQps"`

    /* 开启的 CC 防护规则数量 (Optional) */
    CcProtectionRuleEnableCount int64 `json:"ccProtectionRuleEnableCount"`
}