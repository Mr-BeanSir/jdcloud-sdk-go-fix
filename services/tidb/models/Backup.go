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


type Backup struct {

    /* 备份ID (Optional) */
    BackupId string `json:"backupId"`

    /* 备份名称，最长支持64个英文字符或等长的中文字符 (Optional) */
    BackupName string `json:"backupName"`

    /* 备份所属实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 备份状态 (Optional) */
    BackupStatus string `json:"backupStatus"`

    /* 备份开始时间 (Optional) */
    BackupStartTime string `json:"backupStartTime"`

    /* 备份结束时间 (Optional) */
    BackupEndTime string `json:"backupEndTime"`

    /* 备份类型 (Optional) */
    BackupType string `json:"backupType"`

    /* 备份模式 (Optional) */
    BackupMode string `json:"backupMode"`

    /* 整个备份集大小，单位：Byte (Optional) */
    BackupSizeByte int64 `json:"backupSizeByte"`
}
