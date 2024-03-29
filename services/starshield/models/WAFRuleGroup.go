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


type WAFRuleGroup struct {

    /* WAF组标识符标签 (Optional) */
    Id string `json:"id"`

    /* 防火墙规则组的名称 (Optional) */
    Name string `json:"name"`

    /* WAF规则集的功能摘要 (Optional) */
    Description string `json:"description"`

    /* 此组中包含多少条规则 (Optional) */
    Rules_count int `json:"rules_count"`

    /* 组中有多少规则已从默认规则修改 (Optional) */
    Modified_rules_count int `json:"modified_rules_count"`

    /* WAF包标识符标签 (Optional) */
    Package_id string `json:"package_id"`

    /* 此组中包含的规则是否可配置/可用 (Optional) */
    Mode string `json:"mode"`

    /* 组可以具有的可用状态。这将影响此组中规则的状态。 (Optional) */
    Allowed_modes []string `json:"allowed_modes"`
}
