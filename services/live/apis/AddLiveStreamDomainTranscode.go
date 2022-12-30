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

type AddLiveStreamDomainTranscodeRequest struct {

    core.JDCloudRequest

    /* 直播的推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 转码模版
- 取值范围: 系统标准转码模板, 用户自定义转码模板
- 系统标准转码模板
  ld (h.264/640*360/15f)
  sd (h.264/960*540/25f)
  hd (h.264/1280*720/25f)
  shd (h.264/1920*1080/30f)
  ld-265 (h.265/640*360/15f)
  sd-265 (h.265/960*540/25f)
  hd-265 (h.265/1280*720/25f)
  shd-265 (h.265/1920*1080/30f)
  */
    Template string `json:"template"`
}

/*
 * param publishDomain: 直播的推流域名 (Required)
 * param template: 转码模版
- 取值范围: 系统标准转码模板, 用户自定义转码模板
- 系统标准转码模板
  ld (h.264/640*360/15f)
  sd (h.264/960*540/25f)
  hd (h.264/1280*720/25f)
  shd (h.264/1920*1080/30f)
  ld-265 (h.265/640*360/15f)
  sd-265 (h.265/960*540/25f)
  hd-265 (h.265/1280*720/25f)
  shd-265 (h.265/1920*1080/30f)
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddLiveStreamDomainTranscodeRequest(
    publishDomain string,
    template string,
) *AddLiveStreamDomainTranscodeRequest {

	return &AddLiveStreamDomainTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodeDomains:config",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        Template: template,
	}
}

/*
 * param publishDomain: 直播的推流域名 (Required)
 * param template: 转码模版
- 取值范围: 系统标准转码模板, 用户自定义转码模板
- 系统标准转码模板
  ld (h.264/640*360/15f)
  sd (h.264/960*540/25f)
  hd (h.264/1280*720/25f)
  shd (h.264/1920*1080/30f)
  ld-265 (h.265/640*360/15f)
  sd-265 (h.265/960*540/25f)
  hd-265 (h.265/1280*720/25f)
  shd-265 (h.265/1920*1080/30f)
 (Required)
 */
func NewAddLiveStreamDomainTranscodeRequestWithAllParams(
    publishDomain string,
    template string,
) *AddLiveStreamDomainTranscodeRequest {

    return &AddLiveStreamDomainTranscodeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeDomains:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddLiveStreamDomainTranscodeRequestWithoutParam() *AddLiveStreamDomainTranscodeRequest {

    return &AddLiveStreamDomainTranscodeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeDomains:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 直播的推流域名(Required) */
func (r *AddLiveStreamDomainTranscodeRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param template: 转码模版
- 取值范围: 系统标准转码模板, 用户自定义转码模板
- 系统标准转码模板
  ld (h.264/640*360/15f)
  sd (h.264/960*540/25f)
  hd (h.264/1280*720/25f)
  shd (h.264/1920*1080/30f)
  ld-265 (h.265/640*360/15f)
  sd-265 (h.265/960*540/25f)
  hd-265 (h.265/1280*720/25f)
  shd-265 (h.265/1920*1080/30f)
(Required) */
func (r *AddLiveStreamDomainTranscodeRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddLiveStreamDomainTranscodeRequest) GetRegionId() string {
    return ""
}

type AddLiveStreamDomainTranscodeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddLiveStreamDomainTranscodeResult `json:"result"`
}

type AddLiveStreamDomainTranscodeResult struct {
}