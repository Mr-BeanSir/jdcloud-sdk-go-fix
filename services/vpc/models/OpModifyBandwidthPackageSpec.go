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


type OpModifyBandwidthPackageSpec struct {

    /* 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，限制取值必须为5的整倍数，且不能低于共享带宽包内公网IP带宽上限 (Optional) */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符 (Optional) */
    Name string `json:"name"`

    /* 描述，长度不超过256个字符 (Optional) */
    Description string `json:"description"`

    /* 资源所属的用户pin (Optional) */
    UserPin string `json:"userPin"`
}
