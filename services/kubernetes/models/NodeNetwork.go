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


type NodeNetwork struct {

    /* pod子网的id (Optional) */
    PodSubnetId string `json:"podSubnetId"`

    /* node子网的id (Optional) */
    NodeSubnetId string `json:"nodeSubnetId"`

    /* service子网的id (Optional) */
    ServiceSubnetId string `json:"serviceSubnetId"`

    /* service关联LB的具有公网访问能力的子网id (Optional) */
    ServicePublicSubnetId string `json:"servicePublicSubnetId"`

    /* node的cidr (Optional) */
    NodeNetworkCidr string `json:"nodeNetworkCidr"`

    /* vpc id (Optional) */
    VpcId string `json:"vpcId"`
}