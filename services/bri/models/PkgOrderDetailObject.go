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


type PkgOrderDetailObject struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 产品名称 (Optional) */
    PackageType string `json:"packageType"`

    /* 规格 (Optional) */
    Specs string `json:"specs"`

    /* 付费类型 (Optional) */
    PayType string `json:"payType"`

    /* 下单时间, yyyy-mm-dd hh:mm:ss格式 (Optional) */
    PayTime string `json:"payTime"`

    /* 订单状态：PAID - 支付,PRE-PAID - 预支付 (Optional) */
    Status string `json:"status"`

    /* 订单金额 (Optional) */
    TotalFee int `json:"totalFee"`

    /* 代金券 (Optional) */
    FavorableFee int `json:"favorableFee"`

    /* 余额 (Optional) */
    BalanceFee int `json:"balanceFee"`

    /* 实付总金额 (Optional) */
    ActualFee int `json:"actualFee"`
}