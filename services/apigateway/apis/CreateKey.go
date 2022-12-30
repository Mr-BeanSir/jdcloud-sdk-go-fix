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

type CreateKeyRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* key 名称 (Optional) */
    KeyName *string `json:"keyName"`

    /* key 描述 (Optional) */
    KeyDesc *string `json:"keyDesc"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateKeyRequest(
    regionId string,
) *CreateKeyRequest {

	return &CreateKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/kms",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param keyName: key 名称 (Optional)
 * param keyDesc: key 描述 (Optional)
 */
func NewCreateKeyRequestWithAllParams(
    regionId string,
    keyName *string,
    keyDesc *string,
) *CreateKeyRequest {

    return &CreateKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/kms",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        KeyName: keyName,
        KeyDesc: keyDesc,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateKeyRequestWithoutParam() *CreateKeyRequest {

    return &CreateKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/kms",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateKeyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param keyName: key 名称(Optional) */
func (r *CreateKeyRequest) SetKeyName(keyName string) {
    r.KeyName = &keyName
}

/* param keyDesc: key 描述(Optional) */
func (r *CreateKeyRequest) SetKeyDesc(keyDesc string) {
    r.KeyDesc = &keyDesc
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateKeyRequest) GetRegionId() string {
    return r.RegionId
}

type CreateKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateKeyResult `json:"result"`
}

type CreateKeyResult struct {
    AccessKey string `json:"accessKey"`
    SecretKey string `json:"secretKey"`
}