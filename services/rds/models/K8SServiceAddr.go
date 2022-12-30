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


type K8SServiceAddr struct {

    /* 连接信息的类型，目前支持以下类型：<br>- database 通常数据访问，读写等 <br>- pd 数据迁移时连接PD节点 <br>- monitor 查看监控数据 <br>- dms 使用DMS客户端，访问数据库 <br>-其他需要的类型等，各产品可视需要添加<br>参数大小敏感 (Optional) */
    AddrType string `json:"addrType"`

    /* 从K8S集群外部访问实例的方式，目前支持以下两种类型 - NodePort - LoadBalancer  参数大小敏感 (Optional) */
    AccessType string `json:"accessType"`

    /* 从K8S集群外部访问实例的地址,如域名或IP (Optional) */
    Addr string `json:"addr"`

    /* 端口 (Optional) */
    Port string `json:"port"`
}