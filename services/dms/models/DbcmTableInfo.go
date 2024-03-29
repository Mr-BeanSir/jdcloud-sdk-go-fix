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


type DbcmTableInfo struct {

    /* 表名。 (Optional) */
    Table_name string `json:"table_name"`

    /* 字符编码。 (Optional) */
    Charset string `json:"charset"`

    /* 表注释。 (Optional) */
    Table_comment string `json:"table_comment"`

    /* 表类型，未使用。 (Optional) */
    Table_type string `json:"table_type"`

    /* 切分键，未使用。 (Optional) */
    Split_key string `json:"split_key"`
}
