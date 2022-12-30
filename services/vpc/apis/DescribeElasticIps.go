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
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeElasticIpsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* elasticIpIds - elasticip id数组条件，支持多个
elasticIpAddress - eip的IP地址，支持单个
chargeStatus	- eip的费用支付状态,normal(正常状态) or overdue(预付费已到期) or arrear(欠费状态)，支持单个
ipType - eip类型，取值：all(所有类型)、standard(标准弹性IP)、edge(边缘弹性IP)，默认standard，支持单个
azs - eip可用区，支持多个
bandwidthPackageId - 共享带宽包ID，支持单个
status - IP是否被绑定，取值：ASSOCIATED（被绑定）、NOT_ASSOCIATED（未被绑定）、ALL（全部）。支持单个
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* Tag筛选条件 (Optional) */
    Tags []vpc.TagFilter `json:"tags"`

    /* 资源组筛选条件 (Optional) */
    ResourceGroupIds []string `json:"resourceGroupIds"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeElasticIpsRequest(
    regionId string,
) *DescribeElasticIpsRequest {

	return &DescribeElasticIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: elasticIpIds - elasticip id数组条件，支持多个
elasticIpAddress - eip的IP地址，支持单个
chargeStatus	- eip的费用支付状态,normal(正常状态) or overdue(预付费已到期) or arrear(欠费状态)，支持单个
ipType - eip类型，取值：all(所有类型)、standard(标准弹性IP)、edge(边缘弹性IP)，默认standard，支持单个
azs - eip可用区，支持多个
bandwidthPackageId - 共享带宽包ID，支持单个
status - IP是否被绑定，取值：ASSOCIATED（被绑定）、NOT_ASSOCIATED（未被绑定）、ALL（全部）。支持单个
 (Optional)
 * param tags: Tag筛选条件 (Optional)
 * param resourceGroupIds: 资源组筛选条件 (Optional)
 */
func NewDescribeElasticIpsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
    tags []vpc.TagFilter,
    resourceGroupIds []string,
) *DescribeElasticIpsRequest {

    return &DescribeElasticIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
        Tags: tags,
        ResourceGroupIds: resourceGroupIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeElasticIpsRequestWithoutParam() *DescribeElasticIpsRequest {

    return &DescribeElasticIpsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeElasticIpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeElasticIpsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeElasticIpsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: elasticIpIds - elasticip id数组条件，支持多个
elasticIpAddress - eip的IP地址，支持单个
chargeStatus	- eip的费用支付状态,normal(正常状态) or overdue(预付费已到期) or arrear(欠费状态)，支持单个
ipType - eip类型，取值：all(所有类型)、standard(标准弹性IP)、edge(边缘弹性IP)，默认standard，支持单个
azs - eip可用区，支持多个
bandwidthPackageId - 共享带宽包ID，支持单个
status - IP是否被绑定，取值：ASSOCIATED（被绑定）、NOT_ASSOCIATED（未被绑定）、ALL（全部）。支持单个
(Optional) */
func (r *DescribeElasticIpsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

/* param tags: Tag筛选条件(Optional) */
func (r *DescribeElasticIpsRequest) SetTags(tags []vpc.TagFilter) {
    r.Tags = tags
}

/* param resourceGroupIds: 资源组筛选条件(Optional) */
func (r *DescribeElasticIpsRequest) SetResourceGroupIds(resourceGroupIds []string) {
    r.ResourceGroupIds = resourceGroupIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeElasticIpsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeElasticIpsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeElasticIpsResult `json:"result"`
}

type DescribeElasticIpsResult struct {
    ElasticIps []vpc.ElasticIp `json:"elasticIps"`
    TotalCount int `json:"totalCount"`
}