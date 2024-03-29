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


type InternalAttackLog struct {

    /* 公网 IP 地址 (Optional) */
    Ip string `json:"ip"`

    /* 攻击记录 ID (Optional) */
    AttackLogId string `json:"attackLogId"`

    /* 攻击开始时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    StartTime string `json:"startTime"`

    /* 攻击结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    EndTime string `json:"endTime"`

    /* normal: 正常, unregister: 未备案, illegalmail: 非法邮件, clean: 超阈值, blackhole: 黑洞 (Optional) */
    AttackStatus []string `json:"attackStatus"`
}
