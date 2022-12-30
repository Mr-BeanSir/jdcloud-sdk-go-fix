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


type VideoTask struct {

    /* 数据Id。需要保证在一次请求中所有的Id不重复 (Optional) */
    DataId string `json:"dataId"`

    /* 待检测视频的URL，最大200M (Optional) */
    Url string `json:"url"`

    /* 视频截帧间隔，单位为秒，取值范围为1~60。默认值为1秒 (Optional) */
    Interval int `json:"interval"`

    /* 本视频截帧的张数上限，取值范围为5~3600，默认为200张，该参数仅在文件检测中生效(live=false) 如果是视频流(live=true)该参数无效。 (Optional) */
    MaxFrames int `json:"maxFrames"`
}