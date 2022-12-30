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

import vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"

type InstanceNetworkInterface struct {

    /* 弹性网卡ID。 (Optional) */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 弹性网卡MAC地址。 (Optional) */
    MacAddress string `json:"macAddress"`

    /* 弹性网卡所属VPC的ID。 (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网ID。 (Optional) */
    SubnetId string `json:"subnetId"`

    /*  (Optional) */
    SecurityGroups []SecurityGroupSimple `json:"securityGroups"`

    /* PortSecurity，源和目标IP地址校验，取值为0或者1。 (Optional) */
    SanityCheck int `json:"sanityCheck"`

    /* 网卡主IP配置。 (Optional) */
    PrimaryIp vpc.NetworkInterfacePrivateIp `json:"primaryIp"`

    /* 网卡辅IP地址列表。 (Optional) */
    SecondaryIps []vpc.NetworkInterfacePrivateIp `json:"secondaryIps"`
}