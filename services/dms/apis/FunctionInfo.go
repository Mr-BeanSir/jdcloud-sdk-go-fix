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
    dms "github.com/jdcloud-api/jdcloud-sdk-go/services/dms/models"
)

type FunctionInfoRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 数据库名称。 (Optional) */
    DbName *string `json:"dbName"`

    /* 函数名称。 (Optional) */
    FunctionName *string `json:"functionName"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewFunctionInfoRequest(
    regionId string,
) *FunctionInfoRequest {

	return &FunctionInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/function:info",
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
 * param functionName: 函数名称。 (Optional)
 */
func NewFunctionInfoRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    dbName *string,
    functionName *string,
) *FunctionInfoRequest {

    return &FunctionInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function:info",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        DbName: dbName,
        FunctionName: functionName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewFunctionInfoRequestWithoutParam() *FunctionInfoRequest {

    return &FunctionInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function:info",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *FunctionInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param dataSourceId: 数据源id(Optional) */
func (r *FunctionInfoRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}
/* param dbName: 数据库名称。(Optional) */
func (r *FunctionInfoRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}
/* param functionName: 函数名称。(Optional) */
func (r *FunctionInfoRequest) SetFunctionName(functionName string) {
    r.FunctionName = &functionName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r FunctionInfoRequest) GetRegionId() string {
    return r.RegionId
}

type FunctionInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result FunctionInfoResult `json:"result"`
}

type FunctionInfoResult struct {
    DbName string `json:"dbName"`
    FunctionName string `json:"functionName"`
    FunctionSecurity string `json:"functionSecurity"`
    DataAccess string `json:"dataAccess"`
    Deterministic bool `json:"deterministic"`
    DefinitionSql string `json:"definitionSql"`
    CompleteSql string `json:"completeSql"`
    Parameters []dms.Parameter `json:"parameters"`
    Definer string `json:"definer"`
    Comment string `json:"comment"`
    ReturnType string `json:"returnType"`
    ReturnLength int `json:"returnLength"`
    ReturnDecimal int `json:"returnDecimal"`
}