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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    jdccs "github.com/jdcloud-api/jdcloud-sdk-go/services/jdccs/models"
)

type CreateGeneralServicesTicketRequest struct {

    core.JDCloudRequest

    /* 提单人手机号 (Optional) */
    Phone *string `json:"phone"`

    /* 提单人邮箱 (Optional) */
    Email *string `json:"email"`

    /* idc机房实例id (Optional) */
    Idc *string `json:"idc"`

    /* 是否是商业化外部机房 (Optional) */
    IsExternalIdc *bool `json:"isExternalIdc"`

    /* 数量 (Optional) */
    Count *int `json:"count"`

    /* 外部机房地址 (Optional) */
    ExternalIdcAddress *string `json:"externalIdcAddress"`

    /* 外部机房联系人 (Optional) */
    ExternalIdcContactPerson *string `json:"externalIdcContactPerson"`

    /* 外部机房联系电话 (Optional) */
    ExternalIdcContactPhone *string `json:"externalIdcContactPhone"`

    /* 描述 (Optional) */
    Remarks *string `json:"remarks"`

    /* 附件 (Optional) */
    Attach []jdccs.Attach `json:"attach"`

}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateGeneralServicesTicketRequest(
) *CreateGeneralServicesTicketRequest {

	return &CreateGeneralServicesTicketRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/generalServicesTicket",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param phone: 提单人手机号 (Optional)
 * param email: 提单人邮箱 (Optional)
 * param idc: idc机房实例id (Optional)
 * param isExternalIdc: 是否是商业化外部机房 (Optional)
 * param count: 数量 (Optional)
 * param externalIdcAddress: 外部机房地址 (Optional)
 * param externalIdcContactPerson: 外部机房联系人 (Optional)
 * param externalIdcContactPhone: 外部机房联系电话 (Optional)
 * param remarks: 描述 (Optional)
 * param attach: 附件 (Optional)
 */
func NewCreateGeneralServicesTicketRequestWithAllParams(
    phone *string,
    email *string,
    idc *string,
    isExternalIdc *bool,
    count *int,
    externalIdcAddress *string,
    externalIdcContactPerson *string,
    externalIdcContactPhone *string,
    remarks *string,
    attach []jdccs.Attach,
) *CreateGeneralServicesTicketRequest {

    return &CreateGeneralServicesTicketRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/generalServicesTicket",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Phone: phone,
        Email: email,
        Idc: idc,
        IsExternalIdc: isExternalIdc,
        Count: count,
        ExternalIdcAddress: externalIdcAddress,
        ExternalIdcContactPerson: externalIdcContactPerson,
        ExternalIdcContactPhone: externalIdcContactPhone,
        Remarks: remarks,
        Attach: attach,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateGeneralServicesTicketRequestWithoutParam() *CreateGeneralServicesTicketRequest {

    return &CreateGeneralServicesTicketRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/generalServicesTicket",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param phone: 提单人手机号(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetPhone(phone string) {
    r.Phone = &phone
}

/* param email: 提单人邮箱(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetEmail(email string) {
    r.Email = &email
}

/* param idc: idc机房实例id(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetIdc(idc string) {
    r.Idc = &idc
}

/* param isExternalIdc: 是否是商业化外部机房(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetIsExternalIdc(isExternalIdc bool) {
    r.IsExternalIdc = &isExternalIdc
}

/* param count: 数量(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetCount(count int) {
    r.Count = &count
}

/* param externalIdcAddress: 外部机房地址(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetExternalIdcAddress(externalIdcAddress string) {
    r.ExternalIdcAddress = &externalIdcAddress
}

/* param externalIdcContactPerson: 外部机房联系人(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetExternalIdcContactPerson(externalIdcContactPerson string) {
    r.ExternalIdcContactPerson = &externalIdcContactPerson
}

/* param externalIdcContactPhone: 外部机房联系电话(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetExternalIdcContactPhone(externalIdcContactPhone string) {
    r.ExternalIdcContactPhone = &externalIdcContactPhone
}

/* param remarks: 描述(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetRemarks(remarks string) {
    r.Remarks = &remarks
}

/* param attach: 附件(Optional) */
func (r *CreateGeneralServicesTicketRequest) SetAttach(attach []jdccs.Attach) {
    r.Attach = attach
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateGeneralServicesTicketRequest) GetRegionId() string {
    return ""
}

type CreateGeneralServicesTicketResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateGeneralServicesTicketResult `json:"result"`
}

type CreateGeneralServicesTicketResult struct {
    TicketNo string `json:"ticketNo"`
}