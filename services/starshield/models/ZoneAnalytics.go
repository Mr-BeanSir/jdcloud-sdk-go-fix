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


type ZoneAnalytics struct {

    /*  (Optional) */
    Totals Totals `json:"totals"`

    /* Time deltas containing metadata about each bucket of time. The number of buckets (resolution) is determined by the amount of time between the since and until parameters. (Optional) */
    Timeseries []Totals `json:"timeseries"`
}
