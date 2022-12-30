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


type CustomReq struct {

    /* 检测类型，api/oss/website,默认api (Optional) */
    CensorType string `json:"censorType"`

    /* 站点检查实例Id，多个以 , 分割；当censorType为website时，该参数必填 (Optional) */
    WebsiteInstanceId string `json:"websiteInstanceId"`

    /* 文件类型，text-文本，image-图片，audio-音频，video-视频  */
    ResourceType string `json:"resourceType"`

    /* 匹配方式，exact:精确匹配，fuzzy:模糊匹配；仅限文本类型,默认exact (Optional) */
    MatchType string `json:"matchType"`

    /* 敏感库名  */
    Name string `json:"name"`

    /* 文本/语音支持 antispam-反垃圾，视频/图片支持 porn-涉黄，terrorism-涉政暴恐，ad-图文广告  */
    Scenes string `json:"scenes"`

    /* pass 白名单，block 黑名单，review 疑似名单  */
    Suggestion string `json:"suggestion"`

    /* 状态 1启用，0禁用,默认 1启用 (Optional) */
    Status int `json:"status"`

    /* 敏感库id，更新时该参数必填 (Optional) */
    LibId string `json:"libId"`

    /* 敏感库来源：custom自定义，feedback系统库，更新时该参数必填 (Optional) */
    Source string `json:"source"`

    /* 机审策略，可以不填，为空时前端显示空即可 (Optional) */
    BizType string `json:"bizType"`
}