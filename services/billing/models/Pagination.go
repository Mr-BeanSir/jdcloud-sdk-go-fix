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


type Pagination struct {

    /* 当前页(默认:1) (Optional) */
    CurrPageNo int `json:"currPageNo"`

    /* 总页数 (Optional) */
    NumberPages int `json:"numberPages"`

    /* 总记录数 (Optional) */
    NumberRecords int `json:"numberRecords"`

    /* 每页记录数(默认每页:10) (Optional) */
    PageSize int `json:"pageSize"`

    /* 起始页 (Optional) */
    StartIndex int `json:"startIndex"`
}
