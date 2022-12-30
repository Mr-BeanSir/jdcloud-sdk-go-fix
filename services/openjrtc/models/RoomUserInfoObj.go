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


type RoomUserInfoObj struct {

    /* appId (Optional) */
    AppId string `json:"appId"`

    /* 用户定义的房间号 (Optional) */
    UserRoomId string `json:"userRoomId"`

    /* 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Optional) */
    UserId string `json:"userId"`

    /* 用户房间内昵称 (Optional) */
    NickName string `json:"nickName"`

    /* 用户socketIo长连接id (Optional) */
    ConnectId string `json:"connectId"`

    /* 状态 1-在线 2-离线 (Optional) */
    Status int `json:"status"`

    /* 创建时间 (Optional) */
    JoinTime string `json:"joinTime"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}