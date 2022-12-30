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
    dbs "github.com/jdcloud-api/jdcloud-sdk-go/services/dbs/models"
)

type RestorePhysicalBackupRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划ID  */
    BackupPlanId string `json:"backupPlanId"`

    /* 用于恢复的备份Id，仅限于本备份计划生成的备份  */
    BackupId string `json:"backupId"`

    /* 备份实例的数据源信息  */
    SourceEndpoint *dbs.RestoreSourceEndpoint `json:"sourceEndpoint"`

    /* 是否是新建数据源；true：是；false：不是  */
    CreateNewEndpoint bool `json:"createNewEndpoint"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 * param backupId: 用于恢复的备份Id，仅限于本备份计划生成的备份 (Required)
 * param sourceEndpoint: 备份实例的数据源信息 (Required)
 * param createNewEndpoint: 是否是新建数据源；true：是；false：不是 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRestorePhysicalBackupRequest(
    regionId string,
    backupPlanId string,
    backupId string,
    sourceEndpoint *dbs.RestoreSourceEndpoint,
    createNewEndpoint bool,
) *RestorePhysicalBackupRequest {

	return &RestorePhysicalBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:restorePhysicalBackup",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        BackupPlanId: backupPlanId,
        BackupId: backupId,
        SourceEndpoint: sourceEndpoint,
        CreateNewEndpoint: createNewEndpoint,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 * param backupId: 用于恢复的备份Id，仅限于本备份计划生成的备份 (Required)
 * param sourceEndpoint: 备份实例的数据源信息 (Required)
 * param createNewEndpoint: 是否是新建数据源；true：是；false：不是 (Required)
 */
func NewRestorePhysicalBackupRequestWithAllParams(
    regionId string,
    backupPlanId string,
    backupId string,
    sourceEndpoint *dbs.RestoreSourceEndpoint,
    createNewEndpoint bool,
) *RestorePhysicalBackupRequest {

    return &RestorePhysicalBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:restorePhysicalBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BackupPlanId: backupPlanId,
        BackupId: backupId,
        SourceEndpoint: sourceEndpoint,
        CreateNewEndpoint: createNewEndpoint,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRestorePhysicalBackupRequestWithoutParam() *RestorePhysicalBackupRequest {

    return &RestorePhysicalBackupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}:restorePhysicalBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *RestorePhysicalBackupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupPlanId: 备份计划ID(Required) */
func (r *RestorePhysicalBackupRequest) SetBackupPlanId(backupPlanId string) {
    r.BackupPlanId = backupPlanId
}

/* param backupId: 用于恢复的备份Id，仅限于本备份计划生成的备份(Required) */
func (r *RestorePhysicalBackupRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}

/* param sourceEndpoint: 备份实例的数据源信息(Required) */
func (r *RestorePhysicalBackupRequest) SetSourceEndpoint(sourceEndpoint *dbs.RestoreSourceEndpoint) {
    r.SourceEndpoint = sourceEndpoint
}

/* param createNewEndpoint: 是否是新建数据源；true：是；false：不是(Required) */
func (r *RestorePhysicalBackupRequest) SetCreateNewEndpoint(createNewEndpoint bool) {
    r.CreateNewEndpoint = createNewEndpoint
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestorePhysicalBackupRequest) GetRegionId() string {
    return r.RegionId
}

type RestorePhysicalBackupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestorePhysicalBackupResult `json:"result"`
}

type RestorePhysicalBackupResult struct {
}