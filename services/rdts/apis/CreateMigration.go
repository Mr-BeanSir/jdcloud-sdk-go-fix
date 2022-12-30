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
    rdts "github.com/jdcloud-api/jdcloud-sdk-go/services/rdts/models"
)

type CreateMigrationRequest struct {

    core.JDCloudRequest

    /* 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1)  */
    RegionId string `json:"regionId"`

    /* 迁移配置，创建成功后可修改  */
    MigrateConfig *rdts.MigrateConfig `json:"migrateConfig"`

    /* 可用区与网络配置，创建成功后不能修改 (Optional) */
    AzAndNetConfig *rdts.AzAndNetConfig `json:"azAndNetConfig"`
}

/*
 * param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1) (Required)
 * param migrateConfig: 迁移配置，创建成功后可修改 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateMigrationRequest(
    regionId string,
    migrateConfig *rdts.MigrateConfig,
) *CreateMigrationRequest {

	return &CreateMigrationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        MigrateConfig: migrateConfig,
	}
}

/*
 * param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1) (Required)
 * param migrateConfig: 迁移配置，创建成功后可修改 (Required)
 * param azAndNetConfig: 可用区与网络配置，创建成功后不能修改 (Optional)
 */
func NewCreateMigrationRequestWithAllParams(
    regionId string,
    migrateConfig *rdts.MigrateConfig,
    azAndNetConfig *rdts.AzAndNetConfig,
) *CreateMigrationRequest {

    return &CreateMigrationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        MigrateConfig: migrateConfig,
        AzAndNetConfig: azAndNetConfig,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateMigrationRequestWithoutParam() *CreateMigrationRequest {

    return &CreateMigrationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 迁移任务所在区域的Region ID。华北-北京(cn-north-1)，华东-上海(cn-east-2)，华南-广州(cn-south-1)(Required) */
func (r *CreateMigrationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param migrateConfig: 迁移配置，创建成功后可修改(Required) */
func (r *CreateMigrationRequest) SetMigrateConfig(migrateConfig *rdts.MigrateConfig) {
    r.MigrateConfig = migrateConfig
}

/* param azAndNetConfig: 可用区与网络配置，创建成功后不能修改(Optional) */
func (r *CreateMigrationRequest) SetAzAndNetConfig(azAndNetConfig *rdts.AzAndNetConfig) {
    r.AzAndNetConfig = azAndNetConfig
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateMigrationRequest) GetRegionId() string {
    return r.RegionId
}

type CreateMigrationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateMigrationResult `json:"result"`
}

type CreateMigrationResult struct {
    InstanceId string `json:"instanceId"`
}