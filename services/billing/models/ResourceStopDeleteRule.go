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


type ResourceStopDeleteRule struct {

    /* 站点  */
    Site int `json:"site"`

    /* 产品线  */
    AppCode string `json:"appCode"`

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 规则类型 1：试用规则 2、用户产品规则 3：用户规则 4：产品规则 5：通用规则 6：用户等级产品规则  */
    RuleType int `json:"ruleType"`

    /* 产品 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 按配置欠费是否停服  1：欠费需要停服 0：欠费不需要停服  */
    ArrearStop int `json:"arrearStop"`

    /* 按配置欠费停服延后时长(0：立即停服  n：n几小时后停服)  */
    ArrearStopDelayHours int `json:"arrearStopDelayHours"`

    /* 按配置欠费停服是否释放资源  1：需要释放资源 0：不需要释放资源  */
    ArrearDelete int `json:"arrearDelete"`

    /* 按配置欠费停服释放资源延后时长(0：立即释放资源 n：n几小时后释放资源)  */
    ArrearDeleteDelayHours int `json:"arrearDeleteDelayHours"`

    /* 到期是否停服  1：到期需要停服 0：到期不需要停服  */
    ExpireStop int `json:"expireStop"`

    /* 到期停服延后时长(0：立即停服  n：n几小时后停服)  */
    ExpireStopDelayHours int `json:"expireStopDelayHours"`

    /* 到期停服是否释放资源  1：需要释放资源 0：不需要释放资源  */
    ExpireDelete int `json:"expireDelete"`

    /* 到期停服释放资源延后时长(0：立即释放资源  n：n几小时后释放资源)  */
    ExpireDeleteDelayHours int `json:"expireDeleteDelayHours"`

    /* 按配置欠费释放类型 1：释放资源  2：释放数据  */
    ArrearDeleteType int `json:"arrearDeleteType"`

    /* 到期释放类型 1：释放资源  2：释放数据  */
    ExpireDeleteType int `json:"expireDeleteType"`

    /* 按用量欠费是否停服  1：欠费需要停服 0：欠费不需要停服  */
    FlowArrearStop int `json:"flowArrearStop"`

    /* 按用量欠费停服延后时长  */
    FlowArrearStopDelayHours int `json:"flowArrearStopDelayHours"`

    /* 按用量欠费停服是否释放资源  1：需要释放资源 0：不需要释放资源  */
    FlowArrearDelete int `json:"flowArrearDelete"`

    /* 按用量欠费停服释放资源延后时长  */
    FlowArrearDeleteDelayHours int `json:"flowArrearDeleteDelayHours"`

    /* 按用量欠费释放类型 1：释放资源  2：释放数据  */
    FlowArrearDeleteType int `json:"flowArrearDeleteType"`

    /* 客户级别 1-普通客户 2-VIP客户 (Optional) */
    ClientType int `json:"clientType"`
}
