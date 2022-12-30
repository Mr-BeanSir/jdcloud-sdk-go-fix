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


type RuleCount struct {

    /* 处于报警状态的规则个数  status:2 (Optional) */
    AlarmRuleCount int64 `json:"alarmRuleCount"`

    /* 规则类型,resourceMonitor-资源监控 customMetric-自定义监控 (Optional) */
    RuleType string `json:"ruleType"`

    /* 处于数据不足状态的规则 个数  status:4 (Optional) */
    UnknownRuleCount int64 `json:"unknownRuleCount"`
}