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

type CreateDataSourceTableRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 数据源 ID  */
    DataSourceId string `json:"dataSourceId"`

    /* 表名  */
    TableName string `json:"tableName"`
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param tableName: 表名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateDataSourceTableRequest(
    regionId string,
    dataSourceId string,
    tableName string,
) *CreateDataSourceTableRequest {

	return &CreateDataSourceTableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DataSourceId: dataSourceId,
        TableName: tableName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param tableName: 表名 (Required)
 */
func NewCreateDataSourceTableRequestWithAllParams(
    regionId string,
    dataSourceId string,
    tableName string,
) *CreateDataSourceTableRequest {

    return &CreateDataSourceTableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        TableName: tableName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateDataSourceTableRequestWithoutParam() *CreateDataSourceTableRequest {

    return &CreateDataSourceTableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateDataSourceTableRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源 ID(Required) */
func (r *CreateDataSourceTableRequest) SetDataSourceId(dataSourceId string) {
    r.DataSourceId = dataSourceId
}

/* param tableName: 表名(Required) */
func (r *CreateDataSourceTableRequest) SetTableName(tableName string) {
    r.TableName = tableName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDataSourceTableRequest) GetRegionId() string {
    return r.RegionId
}

type CreateDataSourceTableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateDataSourceTableResult `json:"result"`
}

type CreateDataSourceTableResult struct {
}