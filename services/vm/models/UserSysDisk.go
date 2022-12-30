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


type UserSysDisk struct {

    /* 用户pin。 (Optional) */
    Pin string `json:"pin"`

    /* 地域。 (Optional) */
    Region string `json:"region"`

    /* 系统类型。支持范围：`linux、windows`。 (Optional) */
    SystemType string `json:"systemType"`

    /* 默认本地盘系统盘大小，单位GB。 (Optional) */
    SystemDiskSize int `json:"systemDiskSize"`

    /* 逗号分隔的规格列表，`*` 代表所有。 (Optional) */
    Flavors string `json:"flavors"`
}