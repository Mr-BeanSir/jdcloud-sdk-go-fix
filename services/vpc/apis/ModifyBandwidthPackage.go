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
)

type ModifyBandwidthPackageRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 共享带宽包ID  */
    BandwidthPackageId string `json:"bandwidthPackageId"`

    /* 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，且不能低于共享带宽包内公网IP带宽上限 (Optional) */
    BandwidthMbps *int `json:"bandwidthMbps"`

    /* 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符 (Optional) */
    Name *string `json:"name"`

    /* 描述，长度不超过256个字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param bandwidthPackageId: 共享带宽包ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyBandwidthPackageRequest(
    regionId string,
    bandwidthPackageId string,
) *ModifyBandwidthPackageRequest {

	return &ModifyBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param bandwidthPackageId: 共享带宽包ID (Required)
 * param bandwidthMbps: 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，且不能低于共享带宽包内公网IP带宽上限 (Optional)
 * param name: 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符 (Optional)
 * param description: 描述，长度不超过256个字符 (Optional)
 */
func NewModifyBandwidthPackageRequestWithAllParams(
    regionId string,
    bandwidthPackageId string,
    bandwidthMbps *int,
    name *string,
    description *string,
) *ModifyBandwidthPackageRequest {

    return &ModifyBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
        BandwidthMbps: bandwidthMbps,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyBandwidthPackageRequestWithoutParam() *ModifyBandwidthPackageRequest {

    return &ModifyBandwidthPackageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyBandwidthPackageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bandwidthPackageId: 共享带宽包ID(Required) */
func (r *ModifyBandwidthPackageRequest) SetBandwidthPackageId(bandwidthPackageId string) {
    r.BandwidthPackageId = bandwidthPackageId
}

/* param bandwidthMbps: 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，且不能低于共享带宽包内公网IP带宽上限(Optional) */
func (r *ModifyBandwidthPackageRequest) SetBandwidthMbps(bandwidthMbps int) {
    r.BandwidthMbps = &bandwidthMbps
}

/* param name: 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符(Optional) */
func (r *ModifyBandwidthPackageRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 描述，长度不超过256个字符(Optional) */
func (r *ModifyBandwidthPackageRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyBandwidthPackageRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyBandwidthPackageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyBandwidthPackageResult `json:"result"`
}

type ModifyBandwidthPackageResult struct {
}