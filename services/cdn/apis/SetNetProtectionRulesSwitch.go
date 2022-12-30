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

type SetNetProtectionRulesSwitchRequest struct {

    core.JDCloudRequest

    /* on,off (Optional) */
    SwitchStatus *string `json:"switchStatus"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetNetProtectionRulesSwitchRequest(
) *SetNetProtectionRulesSwitchRequest {

	return &SetNetProtectionRulesSwitchRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/netProtectionSwitch",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param switchStatus: on,off (Optional)
 */
func NewSetNetProtectionRulesSwitchRequestWithAllParams(
    switchStatus *string,
) *SetNetProtectionRulesSwitchRequest {

    return &SetNetProtectionRulesSwitchRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/netProtectionSwitch",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        SwitchStatus: switchStatus,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetNetProtectionRulesSwitchRequestWithoutParam() *SetNetProtectionRulesSwitchRequest {

    return &SetNetProtectionRulesSwitchRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/netProtectionSwitch",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param switchStatus: on,off(Optional) */
func (r *SetNetProtectionRulesSwitchRequest) SetSwitchStatus(switchStatus string) {
    r.SwitchStatus = &switchStatus
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetNetProtectionRulesSwitchRequest) GetRegionId() string {
    return ""
}

type SetNetProtectionRulesSwitchResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetNetProtectionRulesSwitchResult `json:"result"`
}

type SetNetProtectionRulesSwitchResult struct {
}