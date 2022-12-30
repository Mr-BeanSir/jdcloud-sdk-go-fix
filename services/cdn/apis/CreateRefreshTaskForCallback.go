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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type CreateRefreshTaskForCallbackRequest struct {

    core.JDCloudRequest

    /* 刷新预热类型,(url:url刷新,dir:目录刷新,prefetch:预热) (Optional) */
    TaskType *string `json:"taskType"`

    /*  (Optional) */
    UrlItems []cdn.UrlItem `json:"urlItems"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRefreshTaskForCallbackRequest(
) *CreateRefreshTaskForCallbackRequest {

	return &CreateRefreshTaskForCallbackRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/task:createForCallback",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param taskType: 刷新预热类型,(url:url刷新,dir:目录刷新,prefetch:预热) (Optional)
 * param urlItems:  (Optional)
 */
func NewCreateRefreshTaskForCallbackRequestWithAllParams(
    taskType *string,
    urlItems []cdn.UrlItem,
) *CreateRefreshTaskForCallbackRequest {

    return &CreateRefreshTaskForCallbackRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:createForCallback",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        TaskType: taskType,
        UrlItems: urlItems,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRefreshTaskForCallbackRequestWithoutParam() *CreateRefreshTaskForCallbackRequest {

    return &CreateRefreshTaskForCallbackRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:createForCallback",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param taskType: 刷新预热类型,(url:url刷新,dir:目录刷新,prefetch:预热)(Optional) */
func (r *CreateRefreshTaskForCallbackRequest) SetTaskType(taskType string) {
    r.TaskType = &taskType
}
/* param urlItems: (Optional) */
func (r *CreateRefreshTaskForCallbackRequest) SetUrlItems(urlItems []cdn.UrlItem) {
    r.UrlItems = urlItems
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRefreshTaskForCallbackRequest) GetRegionId() string {
    return ""
}

type CreateRefreshTaskForCallbackResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRefreshTaskForCallbackResult `json:"result"`
}

type CreateRefreshTaskForCallbackResult struct {
    ErrorCount int `json:"errorCount"`
    TaskId string `json:"taskId"`
}