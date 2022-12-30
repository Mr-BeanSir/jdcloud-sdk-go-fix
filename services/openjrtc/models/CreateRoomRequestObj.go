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


type CreateRoomRequestObj struct {

    /* 房间名称 (Optional) */
    RoomName string `json:"roomName"`

    /* 应用ID (Optional) */
    AppId string `json:"appId"`

    /* 房间类型 1-小房间(音频单流订阅) 2-大房间(音频固定订阅) (Optional) */
    RoomType int `json:"roomType"`

    /* 会议类型 0-即时会议 1-预约会议 (Optional) */
    MeetingType int `json:"meetingType"`

    /* 用户ID(创建者ID) (Optional) */
    PeerId int64 `json:"peerId"`
}