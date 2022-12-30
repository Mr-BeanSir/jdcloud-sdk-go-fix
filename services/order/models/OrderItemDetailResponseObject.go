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


type OrderItemDetailResponseObject struct {

    /* 订单总金额 (Optional) */
    TotalFee int `json:"totalFee"`

    /* 应付金额（订单总金额-折扣金额） (Optional) */
    ActualFee int `json:"actualFee"`

    /* 余额支付金额 (Optional) */
    BalancePay int `json:"balancePay"`

    /* 计费时长 (Optional) */
    ChargeDuration int `json:"chargeDuration"`

    /* 现金支付金额 (Optional) */
    MoneyPay int `json:"moneyPay"`

    /* 退款金额 (Optional) */
    RefundFee int `json:"refundFee"`

    /* 计费类型(CONFIG-按配置,FLOW-按用量MONTHLY-包年包月,ONCE-按次付费) (Optional) */
    ChargeMode string `json:"chargeMode"`

    /* 订单创建时间（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    CreateTime string `json:"createTime"`

    /* 续费后资源到期时间（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    ExpireDateAfter string `json:"expireDateAfter"`

    /* 续费前资源到期时间（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    ExpireDateBefore string `json:"expireDateBefore"`

    /* 销售属性 (Optional) */
    ExtraInfo []ExtraInfo `json:"extraInfo"`

    /* 续费后资源到期-销售属性 (Optional) */
    ExtraInfoAfter []ExtraInfo `json:"extraInfoAfter"`

    /* 续费前资源到期-销售属性 (Optional) */
    ExtraInfoBefore []ExtraInfo `json:"extraInfoBefore"`

    /* 代金券金额 (Optional) */
    FavorableFee int `json:"favorableFee"`

    /* 配置计费项 (Optional) */
    Formula string `json:"formula"`

    /* 资源id (Optional) */
    ItemId string `json:"itemId"`

    /* 资源名称 (Optional) */
    ItemName string `json:"itemName"`

    /* 订单号 (Optional) */
    OrderNumber string `json:"orderNumber"`

    /* 价格快照 (Optional) */
    PriceSnapshot string `json:"priceSnapshot"`

    /* 数量 (Optional) */
    Quantity int `json:"quantity"`

    /* 备注 (Optional) */
    Remark string `json:"remark"`

    /* 变配明细(UP-升配补差价，DOWN-降配延时,MODIFY_CONFIG-调整配置，RENEW-续费，RENEW_UP-续费升配，RENEW_DOWN-续费降配，MONTHLY-配置转包年包月，RENEW_FREE-补偿续费) (Optional) */
    ResizeFormulaType string `json:"resizeFormulaType"`

    /* 产品名称 (Optional) */
    ServiceName string `json:"serviceName"`

    /* 站点名称（MAIN_SITE-主站，INTERNATIONAL_SITE-国际站，SUQIAN_DEDICATED_CLOUD-宿迁专有云） (Optional) */
    SiteType string `json:"siteType"`

    /* 资源状态（CREATING-创建中,SUCCESS-成功,FAIL-失败） (Optional) */
    Status string `json:"status"`

    /* 计费时长单位（HOUR-小时,DAY-天,MONTH-月,YEAR-年） (Optional) */
    Unit string `json:"unit"`

    /* 子订单 (Optional) */
    OrderItemDetailResponse []OrderItemDetailResponseObject `json:"orderItemDetailResponse"`
}