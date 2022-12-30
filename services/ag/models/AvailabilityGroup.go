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


type AvailabilityGroup struct {

    /* 高可用组ID (Optional) */
    Id string `json:"id"`

    /* 高可用组名称 (Optional) */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 实例模板的ID (Optional) */
    InstanceTemplateId string `json:"instanceTemplateId"`

    /* 可用区域 (Optional) */
    Azs []string `json:"azs"`

    /* 高可用组资源类型 (Optional) */
    AgType string `json:"agType"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 高可用组中实例的数量 (Optional) */
    Count int `json:"count"`

    /* 是否开启自动伸缩 (Optional) */
    AutoScaling bool `json:"autoScaling"`

    /* 高可用组配置类型 (Optional) */
    ConfigurationType string `json:"configurationType"`

    /* 高可用组放置类型 (Optional) */
    PlacementType string `json:"placementType"`

    /* 高可用组中实例数量的限制。 (Optional) */
    InstancesQuotas []InstancesQuota `json:"instancesQuotas"`
}