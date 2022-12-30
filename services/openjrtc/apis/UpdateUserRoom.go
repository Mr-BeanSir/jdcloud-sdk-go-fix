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

type UpdateUserRoomRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 用户房间号 (Optional) */
    UserRoomId *string `json:"userRoomId"`

    /* 房间名称 (Optional) */
    RoomName *string `json:"roomName"`

    /* 房间类型 1-小房间；2-大房间 (Optional) */
    RoomType *int `json:"roomType"`
}

/*
 * param appId: 应用ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateUserRoomRequest(
    appId string,
) *UpdateUserRoomRequest {

	return &UpdateUserRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/updateUserRoom/{appId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
	}
}

/*
 * param appId: 应用ID (Required)
 * param userRoomId: 用户房间号 (Optional)
 * param roomName: 房间名称 (Optional)
 * param roomType: 房间类型 1-小房间；2-大房间 (Optional)
 */
func NewUpdateUserRoomRequestWithAllParams(
    appId string,
    userRoomId *string,
    roomName *string,
    roomType *int,
) *UpdateUserRoomRequest {

    return &UpdateUserRoomRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/updateUserRoom/{appId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        UserRoomId: userRoomId,
        RoomName: roomName,
        RoomType: roomType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateUserRoomRequestWithoutParam() *UpdateUserRoomRequest {

    return &UpdateUserRoomRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/updateUserRoom/{appId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *UpdateUserRoomRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param userRoomId: 用户房间号(Optional) */
func (r *UpdateUserRoomRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = &userRoomId
}

/* param roomName: 房间名称(Optional) */
func (r *UpdateUserRoomRequest) SetRoomName(roomName string) {
    r.RoomName = &roomName
}

/* param roomType: 房间类型 1-小房间；2-大房间(Optional) */
func (r *UpdateUserRoomRequest) SetRoomType(roomType int) {
    r.RoomType = &roomType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateUserRoomRequest) GetRegionId() string {
    return ""
}

type UpdateUserRoomResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateUserRoomResult `json:"result"`
}

type UpdateUserRoomResult struct {
    RoomId int64 `json:"roomId"`
    UserRoomId string `json:"userRoomId"`
    RoomName string `json:"roomName"`
    RoomType int `json:"roomType"`
    AppId string `json:"appId"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}