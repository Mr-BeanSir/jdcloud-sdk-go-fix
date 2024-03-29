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


type UsrBotRules struct {

    /* 规则id (Optional) */
    Id int `json:"id"`

    /* 规则名 (Optional) */
    RuleName string `json:"ruleName"`

    /* 次数阈值 (Optional) */
    DetectThrsd int `json:"detectThrsd"`

    /* 检测时长，秒 (Optional) */
    DetectPeriod int `json:"detectPeriod"`

    /* 匹配条件集,总长度不能超过4096 (Optional) */
    MatchItems []BotMatchItem `json:"matchItems"`

    /* 动作配置 (Optional) */
    Action DenyActionCfg `json:"action"`

    /* 0-使用中 1-禁用 (Optional) */
    Disable int `json:"disable"`

    /* 更新时间 (Optional) */
    UpdateTime int `json:"updateTime"`

    /* 响应状态码 (Optional) */
    Status int `json:"status"`

    /* 状态码数量阀值 (Optional) */
    Ststhrst int `json:"ststhrst"`

    /* 状态码比例阀值 (Optional) */
    StsthrstRatio int `json:"ststhrstRatio"`

    /* 响应码功能是否启用 (Optional) */
    StatusDisable int `json:"statusDisable"`

    /* 规则生效时间是否启用 (Optional) */
    DateDisable int `json:"dateDisable"`

    /* 统计维度 (Optional) */
    Unit string `json:"unit"`

    /* 持续时间, 单位分钟，范围[1-24*60] (Optional) */
    BlockTime int `json:"blockTime"`
}
