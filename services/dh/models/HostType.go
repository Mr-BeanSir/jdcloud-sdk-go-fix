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


type HostType struct {

    /* 专有宿主机机型 (Optional) */
    DedicatedHostType string `json:"dedicatedHostType"`

    /* 专有宿主机机型售卖状态 (Optional) */
    State []HostTypeState `json:"state"`

    /* CPU总数 (Optional) */
    TotalVCPUs int `json:"totalVCPUs"`

    /* 内存总大小，单位MB (Optional) */
    TotalMemoryMB int `json:"totalMemoryMB"`

    /* 本地磁盘总大小，单位GB (Optional) */
    TotalDiskGB int `json:"totalDiskGB"`

    /* GPU总个数 (Optional) */
    TotalGPUs int `json:"totalGPUs"`

    /* 专有宿主机支持的云主机实例规格 (Optional) */
    SupportedInstanceType []string `json:"supportedInstanceType"`
}
