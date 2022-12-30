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


type RestoreTask struct {

    /* 恢复任务ID (Optional) */
    RestoreTaskId string `json:"restoreTaskId"`

    /* 恢复类型,逻辑恢复 LogicalRestore,  物理恢复 PhysicalRestore, 时间点恢复 TimepointRestore (Optional) */
    RestoreType string `json:"restoreType"`

    /* 恢复后的数据时间点, UTC时间格式：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    DataTimestamp string `json:"dataTimestamp"`

    /* 恢复开始时间，UTC时间格式：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    RestoreStartTime string `json:"restoreStartTime"`

    /* 恢复结束时间，UTC时间格式：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    RestoreEndTime string `json:"restoreEndTime"`

    /* 备份大小，单位：Byte (Optional) */
    SizeByte int64 `json:"sizeByte"`

    /* 恢复状态 成功-COMPLETED  运行中-RESTORING  失败-ERROR (Optional) */
    Status string `json:"status"`

    /* 错误信息，仅状态为 ERROR 时返回；其他状态返回空字符串 (Optional) */
    ErrorMessages string `json:"errorMessages"`

    /* 要恢复到的主机的名称 (Optional) */
    HostName string `json:"hostName"`

    /* 要恢复到的主机的地址，可以为IP或者域名 (Optional) */
    Ip string `json:"ip"`

    /* 数据库的端口 (Optional) */
    Port int `json:"port"`
}