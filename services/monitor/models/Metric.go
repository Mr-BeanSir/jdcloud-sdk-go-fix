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


type Metric struct {

    /* 自增id (Optional) */
    Id int `json:"id"`

    /* serviceCode (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* metricIndex (Optional) */
    MetricIndex int `json:"metricIndex"`

    /* metricID (Optional) */
    MetricID string `json:"metricID"`

    /* metricNameCH (Optional) */
    MetricNameCH string `json:"metricNameCH"`

    /* metricNameEN (Optional) */
    MetricNameEN string `json:"metricNameEN"`

    /* valueType (Optional) */
    ValueType string `json:"valueType"`

    /* downsampleAgg (Optional) */
    DownsampleAgg string `json:"downsampleAgg"`

    /* groupAgg (Optional) */
    GroupAgg string `json:"groupAgg"`

    /* isRate (Optional) */
    IsRate int `json:"isRate"`

    /* isSumRate (Optional) */
    IsSumRate int `json:"isSumRate"`

    /* defaultDownsample (Optional) */
    DefaultDownsample string `json:"defaultDownsample"`

    /* upUnit (Optional) */
    UpUnit string `json:"upUnit"`

    /* monitorUnitCH (Optional) */
    MonitorUnitCH string `json:"monitorUnitCH"`

    /* monitorUnitEN (Optional) */
    MonitorUnitEN string `json:"monitorUnitEN"`

    /* alarmUnitCH (Optional) */
    AlarmUnitCH string `json:"alarmUnitCH"`

    /* alarmUnitEN (Optional) */
    AlarmUnitEN string `json:"alarmUnitEN"`

    /* unitConvertFrom (Optional) */
    UnitConvertFrom int `json:"unitConvertFrom"`

    /* unitConvertTo (Optional) */
    UnitConvertTo int `json:"unitConvertTo"`

    /* isShow (Optional) */
    IsShow int `json:"isShow"`

    /* isEnable (Optional) */
    IsEnable int `json:"isEnable"`

    /* defaultTagName (Optional) */
    DefaultTagName string `json:"defaultTagName"`

    /* defaultTagValue (Optional) */
    DefaultTagValue string `json:"defaultTagValue"`

    /* tags (Optional) */
    Tags string `json:"tags"`

    /* isAlarm (Optional) */
    IsAlarm int `json:"isAlarm"`

    /* newNet (Optional) */
    NewNet int `json:"newNet"`

    /* timeInterval (Optional) */
    TimeInterval int `json:"timeInterval"`

    /* 业务线代码 (Optional) */
    ServiceCodes []string `json:"serviceCodes"`

    /* 分组id (Optional) */
    GroupIds []int `json:"groupIds"`

    /* 分组信息 (Optional) */
    GroupCodes string `json:"groupCodes"`

    /* 排序字段 (Optional) */
    Column string `json:"column"`

    /* 排序方式 (Optional) */
    Dir string `json:"dir"`

    /* 页码 (Optional) */
    Number int `json:"number"`

    /* 查询条数 (Optional) */
    Size int `json:"size"`
}
