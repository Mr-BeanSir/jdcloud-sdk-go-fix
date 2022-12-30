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


type UserReqVo struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 用户手机号 (Optional) */
    CscPhone string `json:"cscPhone"`

    /* 用户邮箱 (Optional) */
    CscEmail string `json:"cscEmail"`

    /* 用户名 (Optional) */
    Name string `json:"name"`

    /* 用户类型 (Optional) */
    UserType int `json:"userType"`

    /* 公司名 (Optional) */
    CompanyName string `json:"companyName"`

    /* 起始时间 (Optional) */
    CreateTimeStart string `json:"createTimeStart"`

    /* 结束时间 (Optional) */
    CreateTimeEnd string `json:"createTimeEnd"`

    /* 欠费状态： (Optional) */
    ArrearageStatus int `json:"arrearageStatus"`

    /* 用户分组，多个逗号分隔:1-自然流量，2-内部测试，3-内部重点，4-渠道用户 (Optional) */
    Groups string `json:"groups"`

    /* 用户分组:1-自然流量，2-内部测试，3-内部重点，4-渠道用户 (Optional) */
    Group int `json:"group"`

    /* 计费白名单：1、在白名单  2、不在白名单 (Optional) */
    BillingWhite int `json:"billingWhite"`

    /* 渠道等级;1普通用户2测试用户4VIP用户8其他VIP用户16boss迁移账户 (Optional) */
    Tag int `json:"tag"`

    /* 页大小(必传) (Optional) */
    PageSize int `json:"pageSize"`

    /* 当前页(必传) (Optional) */
    CurrentPage int `json:"currentPage"`
}