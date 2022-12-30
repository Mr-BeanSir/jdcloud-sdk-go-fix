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


type ContractInfo struct {

    /* 合同ID (Optional) */
    ContractId string `json:"contractId"`

    /* 合同标题 (Optional) */
    ContractTitle string `json:"contractTitle"`

    /* 印章名称(可能有多个印章) (Optional) */
    StampNames []string `json:"stampNames"`

    /* 合同文件（base64） (Optional) */
    ContractContent string `json:"contractContent"`

    /* 合同文件摘要 (Optional) */
    ContractDigest string `json:"contractDigest"`

    /* 合同签章时间 (Optional) */
    CreateTime string `json:"createTime"`
}