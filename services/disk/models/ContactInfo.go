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


type ContactInfo struct {

    /* 是否发送短信。0:不发送 1:发送 (Optional) */
    Sms *int `json:"sms"`

    /* 是否发送短信。0:不发送 1:发送 (Optional) */
    Email *int `json:"email"`

    /* 联系人id,默认为空 (Optional) */
    PersonIds []int `json:"personIds"`

    /* 联系组id,默认为空 (Optional) */
    GroupIds []int `json:"groupIds"`
}