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

package models


type RoomUserRecord struct {

    /* 应用ID (Optional) */
    AppId string `json:"appId"`

    /* 业务接入方定义的且在JRTC系统内注册过的用户id (Optional) */
    UserId string `json:"userId"`

    /* 昵称 (Optional) */
    NickName string `json:"nickName"`

    /* 设备名称 (Optional) */
    DeviceName string `json:"deviceName"`

    /* 设备类型 (Optional) */
    DeviceType string `json:"deviceType"`

    /* 设备型号 (Optional) */
    DeviceMode string `json:"deviceMode"`

    /* 系统版本 (Optional) */
    OsVersion string `json:"osVersion"`

    /* 持续时长 (Optional) */
    Duration int64 `json:"duration"`

    /* 加入时间 (Optional) */
    JoinTime string `json:"joinTime"`

    /* 离开时间 (Optional) */
    LeaveTime string `json:"leaveTime"`
}
