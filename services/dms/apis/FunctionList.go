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

type FunctionListRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 数据库名称。 (Optional) */
    DbName *string `json:"dbName"`

    /* 过滤条件。 (Optional) */
    Filter *string `json:"filter"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewFunctionListRequest(
    regionId string,
) *FunctionListRequest {

	return &FunctionListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/function:list",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param dataSourceId: 数据源id (Optional)
 * param dbName: 数据库名称。 (Optional)
 * param filter: 过滤条件。 (Optional)
 */
func NewFunctionListRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    dbName *string,
    filter *string,
) *FunctionListRequest {

    return &FunctionListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        DbName: dbName,
        Filter: filter,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewFunctionListRequestWithoutParam() *FunctionListRequest {

    return &FunctionListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *FunctionListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param dataSourceId: 数据源id(Optional) */
func (r *FunctionListRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}
/* param dbName: 数据库名称。(Optional) */
func (r *FunctionListRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}
/* param filter: 过滤条件。(Optional) */
func (r *FunctionListRequest) SetFilter(filter string) {
    r.Filter = &filter
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r FunctionListRequest) GetRegionId() string {
    return r.RegionId
}

type FunctionListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result FunctionListResult `json:"result"`
}

type FunctionListResult struct {
    FunctionNames []string `json:"functionNames"`
}