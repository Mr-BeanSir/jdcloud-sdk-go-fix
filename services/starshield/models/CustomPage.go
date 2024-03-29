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


type CustomPage struct {

    /* 自定义页面类型的名称 (Optional) */
    Id string `json:"id"`

    /* 创建自定义页面时间 (Optional) */
    Created_on string `json:"created_on"`

    /* 上次修改自定义页面的时间 (Optional) */
    Modified_on string `json:"modified_on"`

    /* 与自定义页面关联的URL。 (Optional) */
    Url string `json:"url"`

    /* 自定义页面状态 (Optional) */
    State string `json:"state"`

    /* 自定义HTML页面中必须存在的字符串标记 (Optional) */
    Required_tokens []string `json:"required_tokens"`

    /* 预览自定义页面时，需要将“target”作为查询字符串的一部分 (Optional) */
    Preview_target string `json:"preview_target"`

    /* 自定义页面的简短描述。 (Optional) */
    Description string `json:"description"`
}
