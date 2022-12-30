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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    kubernetes "github.com/jdcloud-api/jdcloud-sdk-go/services/kubernetes/models"
)

type CreateNodeGroupRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 名称（同一用户的 cluster 内部唯一）  */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description *string `json:"description"`

    /* 工作节点所属的集群  */
    ClusterId string `json:"clusterId"`

    /* 工作节点配置信息  */
    NodeConfig *kubernetes.NodeConfigSpec `json:"nodeConfig"`

    /* 工作节点组的 az，必须为集群az的子集，默认为集群az (Optional) */
    Azs []string `json:"azs"`

    /* 工作节点组初始化大小  */
    InitialNodeCount int `json:"initialNodeCount"`

    /* 工作节点组初始化大小运行的VPC  */
    VpcId string `json:"vpcId"`

    /* 工作节点组的cidr (Optional) */
    NodeCidr *string `json:"nodeCidr"`

    /* 是否开启工作节点组的自动修复，默认关闭 (Optional) */
    AutoRepair *bool `json:"autoRepair"`

    /* 自动伸缩配置 (Optional) */
    CaConfig *kubernetes.CAConfigSpec `json:"caConfig"`

    /* 节点组的网络配置，如果集群的类型customized类型，则必须指定该参数，如果是auto，则不是必须 (Optional) */
    NodeGroupNetwork *kubernetes.NodeGroupNetworkSpec `json:"nodeGroupNetwork"`
}

/*
 * param regionId: 地域 ID (Required)
 * param name: 名称（同一用户的 cluster 内部唯一） (Required)
 * param clusterId: 工作节点所属的集群 (Required)
 * param nodeConfig: 工作节点配置信息 (Required)
 * param initialNodeCount: 工作节点组初始化大小 (Required)
 * param vpcId: 工作节点组初始化大小运行的VPC (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateNodeGroupRequest(
    regionId string,
    name string,
    clusterId string,
    nodeConfig *kubernetes.NodeConfigSpec,
    initialNodeCount int,
    vpcId string,
) *CreateNodeGroupRequest {

	return &CreateNodeGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/nodeGroups",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        ClusterId: clusterId,
        NodeConfig: nodeConfig,
        InitialNodeCount: initialNodeCount,
        VpcId: vpcId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param name: 名称（同一用户的 cluster 内部唯一） (Required)
 * param description: 描述 (Optional)
 * param clusterId: 工作节点所属的集群 (Required)
 * param nodeConfig: 工作节点配置信息 (Required)
 * param azs: 工作节点组的 az，必须为集群az的子集，默认为集群az (Optional)
 * param initialNodeCount: 工作节点组初始化大小 (Required)
 * param vpcId: 工作节点组初始化大小运行的VPC (Required)
 * param nodeCidr: 工作节点组的cidr (Optional)
 * param autoRepair: 是否开启工作节点组的自动修复，默认关闭 (Optional)
 * param caConfig: 自动伸缩配置 (Optional)
 * param nodeGroupNetwork: 节点组的网络配置，如果集群的类型customized类型，则必须指定该参数，如果是auto，则不是必须 (Optional)
 */
func NewCreateNodeGroupRequestWithAllParams(
    regionId string,
    name string,
    description *string,
    clusterId string,
    nodeConfig *kubernetes.NodeConfigSpec,
    azs []string,
    initialNodeCount int,
    vpcId string,
    nodeCidr *string,
    autoRepair *bool,
    caConfig *kubernetes.CAConfigSpec,
    nodeGroupNetwork *kubernetes.NodeGroupNetworkSpec,
) *CreateNodeGroupRequest {

    return &CreateNodeGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Description: description,
        ClusterId: clusterId,
        NodeConfig: nodeConfig,
        Azs: azs,
        InitialNodeCount: initialNodeCount,
        VpcId: vpcId,
        NodeCidr: nodeCidr,
        AutoRepair: autoRepair,
        CaConfig: caConfig,
        NodeGroupNetwork: nodeGroupNetwork,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateNodeGroupRequestWithoutParam() *CreateNodeGroupRequest {

    return &CreateNodeGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *CreateNodeGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: 名称（同一用户的 cluster 内部唯一）(Required) */
func (r *CreateNodeGroupRequest) SetName(name string) {
    r.Name = name
}

/* param description: 描述(Optional) */
func (r *CreateNodeGroupRequest) SetDescription(description string) {
    r.Description = &description
}

/* param clusterId: 工作节点所属的集群(Required) */
func (r *CreateNodeGroupRequest) SetClusterId(clusterId string) {
    r.ClusterId = clusterId
}

/* param nodeConfig: 工作节点配置信息(Required) */
func (r *CreateNodeGroupRequest) SetNodeConfig(nodeConfig *kubernetes.NodeConfigSpec) {
    r.NodeConfig = nodeConfig
}

/* param azs: 工作节点组的 az，必须为集群az的子集，默认为集群az(Optional) */
func (r *CreateNodeGroupRequest) SetAzs(azs []string) {
    r.Azs = azs
}

/* param initialNodeCount: 工作节点组初始化大小(Required) */
func (r *CreateNodeGroupRequest) SetInitialNodeCount(initialNodeCount int) {
    r.InitialNodeCount = initialNodeCount
}

/* param vpcId: 工作节点组初始化大小运行的VPC(Required) */
func (r *CreateNodeGroupRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

/* param nodeCidr: 工作节点组的cidr(Optional) */
func (r *CreateNodeGroupRequest) SetNodeCidr(nodeCidr string) {
    r.NodeCidr = &nodeCidr
}

/* param autoRepair: 是否开启工作节点组的自动修复，默认关闭(Optional) */
func (r *CreateNodeGroupRequest) SetAutoRepair(autoRepair bool) {
    r.AutoRepair = &autoRepair
}

/* param caConfig: 自动伸缩配置(Optional) */
func (r *CreateNodeGroupRequest) SetCaConfig(caConfig *kubernetes.CAConfigSpec) {
    r.CaConfig = caConfig
}

/* param nodeGroupNetwork: 节点组的网络配置，如果集群的类型customized类型，则必须指定该参数，如果是auto，则不是必须(Optional) */
func (r *CreateNodeGroupRequest) SetNodeGroupNetwork(nodeGroupNetwork *kubernetes.NodeGroupNetworkSpec) {
    r.NodeGroupNetwork = nodeGroupNetwork
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateNodeGroupRequest) GetRegionId() string {
    return r.RegionId
}

type CreateNodeGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateNodeGroupResult `json:"result"`
}

type CreateNodeGroupResult struct {
    NodeGroupId string `json:"nodeGroupId"`
}