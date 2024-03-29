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


type SetWafRuleReq struct {

    /* 规则id,新增时传0  */
    Id int `json:"id"`

    /* WAF实例id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 规则名称  */
    RuleName string `json:"ruleName"`

    /* 匹配动作, 拦截:forbidden,redirect 人机识别:verify@jscookie,verify@captcha,verify@rdtcookie 观察:notice  */
    MatchAction string `json:"matchAction"`

    /* 跳转地址，matchAction为redirect时必须为当前实例下的域名的url，forbidden时为自定义页面名称  */
    Redirection string `json:"redirection"`

    /* 条件集  */
    Conditions []ConditionIdSet `json:"conditions"`
}
