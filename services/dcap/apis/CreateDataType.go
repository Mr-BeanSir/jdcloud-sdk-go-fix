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
    dcap "github.com/jdcloud-api/jdcloud-sdk-go/services/dcap/models"
)

type CreateDataTypeRequest struct {

    core.JDCloudRequest

    /* 分类 ID  */
    CategoryId int `json:"categoryId"`

    /* 新增敏感数据类型描述  */
    DataType *dcap.DataTypeSpec `json:"dataType"`
}

/*
 * param categoryId: 分类 ID (Required)
 * param dataType: 新增敏感数据类型描述 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateDataTypeRequest(
    categoryId int,
    dataType *dcap.DataTypeSpec,
) *CreateDataTypeRequest {

	return &CreateDataTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/category/{categoryId}/dataType",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CategoryId: categoryId,
        DataType: dataType,
	}
}

/*
 * param categoryId: 分类 ID (Required)
 * param dataType: 新增敏感数据类型描述 (Required)
 */
func NewCreateDataTypeRequestWithAllParams(
    categoryId int,
    dataType *dcap.DataTypeSpec,
) *CreateDataTypeRequest {

    return &CreateDataTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/category/{categoryId}/dataType",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CategoryId: categoryId,
        DataType: dataType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateDataTypeRequestWithoutParam() *CreateDataTypeRequest {

    return &CreateDataTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/category/{categoryId}/dataType",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param categoryId: 分类 ID(Required) */
func (r *CreateDataTypeRequest) SetCategoryId(categoryId int) {
    r.CategoryId = categoryId
}

/* param dataType: 新增敏感数据类型描述(Required) */
func (r *CreateDataTypeRequest) SetDataType(dataType *dcap.DataTypeSpec) {
    r.DataType = dataType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDataTypeRequest) GetRegionId() string {
    return ""
}

type CreateDataTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateDataTypeResult `json:"result"`
}

type CreateDataTypeResult struct {
}