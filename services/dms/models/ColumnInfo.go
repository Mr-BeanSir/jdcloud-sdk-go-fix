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


type ColumnInfo struct {

    /* 列名。 (Optional) */
    ColumnName string `json:"columnName"`

    /* 新列名，修改表结构时使用。 (Optional) */
    NewColumnName string `json:"newColumnName"`

    /* 列类型。 (Optional) */
    DataType string `json:"dataType"`

    /* 列类型, 返回int(3), varchar(64)等。 (Optional) */
    ColumnType string `json:"columnType"`

    /* 列长度。 (Optional) */
    ColumnLength int `json:"columnLength"`

    /* 浮点数小数点后位数。 (Optional) */
    ColumnScale int `json:"columnScale"`

    /* 是否为空。 (Optional) */
    IsNull bool `json:"isNull"`

    /* 默认值。 (Optional) */
    DefaultValue string `json:"defaultValue"`

    /* 列注释。 (Optional) */
    ColumnComment string `json:"columnComment"`

    /* 是否自增。 (Optional) */
    IsAutoIncrease bool `json:"isAutoIncrease"`

    /* 是否为主键。 (Optional) */
    IsPrimaryKey bool `json:"isPrimaryKey"`
}
