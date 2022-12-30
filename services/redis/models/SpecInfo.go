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


type SpecInfo struct {

    /* 内存大小（GB） (Optional) */
    MemoryGB int `json:"memoryGB"`

    /* 实例规格，标准版不为空，4.0 自定义分片集群版规格为空，具体规格参考单分片规格 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* 实例CPU核数，0表示自定义分片集群版规格，CPU核数由分片数变化 (Optional) */
    Cpu int `json:"cpu"`

    /* 实例磁盘大小（GB)，0表示自定义分片集群版规格，磁盘大小由分片数变化 (Optional) */
    DiskGB int `json:"diskGB"`

    /* 最大连接数，0表示自定义分片集群版规格，最大连接数由分片数变化 (Optional) */
    MaxConnection int `json:"maxConnection"`

    /* 带宽（Mbps)，0表示自定义分片集群版规格，带宽由分片数变化 (Optional) */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* 需要的IP数，0表示自定义分片集群版规格，IP数由分片数变化 (Optional) */
    IpNumber int `json:"ipNumber"`

    /* 实例的分片列表信息，redis 2.8标准版、集群版以及redis 4.0标准版没有分片列表信息 (Optional) */
    Shard ShardInfo `json:"shard"`

    /* az列表 (Optional) */
    Azs []string `json:"azs"`
}