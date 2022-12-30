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


type PodCondition struct {

    /* 最后一次探测状态的时间 (Optional) */
    LastProbeTime string `json:"lastProbeTime"`

    /* 最后一次改变状态的时间 (Optional) */
    LastTransitionTime string `json:"lastTransitionTime"`

    /* 最后一次状态改变的简要原因 (Optional) */
    Reason string `json:"reason"`

    /* Status is the status of the condition. Can be True, False, Unknown. (Optional) */
    Status string `json:"status"`

    /* 最后一次状态改变的信息 (Optional) */
    Message string `json:"message"`

    /* 状态的条件。目前仅限 Ready. (Optional) */
    ConditionType string `json:"conditionType"`
}