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


type InstanceTemplateDiskSpec struct {

    /* 云硬盘类型。各类型介绍请参见[云硬盘类型](https://docs.jdcloud.com/cn/cloud-disk-service/instance-type)。
可选值：
`ssd.gp1`：通用型SSD
`ssd.io1`：性能型SSD
`hdd.std1`：容量型HDD
 (Optional) */
    DiskType *string `json:"diskType"`

    /* 云硬盘容量，单位为 GiB，步长10GiB。
取值范围：
系统盘：`[40,500]`GiB，且不能小于镜像系统盘容量
数据盘：`[20,16000]`GiB，如指定`snapshotId`创建云硬盘则不能小于快照容量。
 (Optional) */
    DiskSizeGB *int `json:"diskSizeGB"`

    /* 创建云硬盘的快照ID。 (Optional) */
    SnapshotId *string `json:"snapshotId"`

    /* 云硬盘自动快照策略ID。 (Optional) */
    PolicyId *string `json:"policyId"`

    /* 云硬盘是否加密。
可选值：
`true`：加密
`false`（默认值）：不加密
 (Optional) */
    Encrypt *bool `json:"encrypt"`

    /* 云硬盘的最大iops。 (Optional) */
    Iops *int `json:"iops"`
}