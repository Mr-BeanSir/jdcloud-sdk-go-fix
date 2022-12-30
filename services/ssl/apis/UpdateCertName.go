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
)

type UpdateCertNameRequest struct {

    core.JDCloudRequest

    /* 证书Id  */
    CertId string `json:"certId"`

    /* 证书名称  */
    CertName string `json:"certName"`
}

/*
 * param certId: 证书Id (Required)
 * param certName: 证书名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateCertNameRequest(
    certId string,
    certName string,
) *UpdateCertNameRequest {

	return &UpdateCertNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/sslCert:updateName",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CertId: certId,
        CertName: certName,
	}
}

/*
 * param certId: 证书Id (Required)
 * param certName: 证书名称 (Required)
 */
func NewUpdateCertNameRequestWithAllParams(
    certId string,
    certName string,
) *UpdateCertNameRequest {

    return &UpdateCertNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert:updateName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CertId: certId,
        CertName: certName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateCertNameRequestWithoutParam() *UpdateCertNameRequest {

    return &UpdateCertNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/sslCert:updateName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param certId: 证书Id(Required) */
func (r *UpdateCertNameRequest) SetCertId(certId string) {
    r.CertId = certId
}

/* param certName: 证书名称(Required) */
func (r *UpdateCertNameRequest) SetCertName(certName string) {
    r.CertName = certName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateCertNameRequest) GetRegionId() string {
    return ""
}

type UpdateCertNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateCertNameResult `json:"result"`
}

type UpdateCertNameResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}