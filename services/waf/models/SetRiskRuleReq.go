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


type SetRiskRuleReq struct {

    /* 规则id,新增时传0 (Optional) */
    Id int `json:"id"`

    /* WAF实例id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 规则名称  */
    Name string `json:"name"`

    /* uri 以/开头  */
    Uri string `json:"uri"`

    /* 匹配动作, 拦截:forbidden,redirect 人机识别:verify@jscookie,verify@captcha,verify@rdtcookie 观察:notice 旗舰版全部支持,高级版不支持人机识别  */
    Action string `json:"action"`

    /* 请求方法 支持 POST:1,GET:1,PUT:1  */
    Methods string `json:"methods"`

    /* 场景 支持 account_login / account_register / data_risk_control  */
    SceneRef string `json:"sceneRef"`

    /* 事件 支持 passwd:BODY.passwd,username:ARGS.username  */
    Event string `json:"event"`

    /* 跳转地址，action为redirect时必须为当前实例下的域名的url，forbidden时为自定义页面名称  */
    Redirection string `json:"redirection"`
}
