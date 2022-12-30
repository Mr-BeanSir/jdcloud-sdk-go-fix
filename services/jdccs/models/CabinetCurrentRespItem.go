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


type CabinetCurrentRespItem struct {

    /* 资源ID (Optional) */
    ResourceId string `json:"resourceId"`

    /* UNIX时间戳 (Optional) */
    Timestamp int `json:"timestamp"`

    /* A路电流 (Optional) */
    AValue float64 `json:"aValue"`

    /* B路电流 (Optional) */
    BValue float64 `json:"bValue"`

    /* 总电流 (Optional) */
    Value float64 `json:"value"`
}