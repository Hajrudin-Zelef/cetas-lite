---
id: collect-261001-huawei/huawei/walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255-3
title: "walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255.md
source_anchor: ""
source_lines: [447, 447]
sha256: 910b6bec1f3bd58bf73a2e03c3a637cc4fda592586946665c45ab3791cfc2926
---

# walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255

2.IPSec(Internet Protocol Security,互联网安全协议)属于第三层协议，通过重新封装IP头部字段并实施加密以实现数据包在公网中传输的安全性。但MPLS属于2.5层隧道，路由器不会拆封MPLS报文的IP头部字段并查询IP路由表转发(MPLS报文IP头部目的地址为私网IP，即使拆封也无法在公网中投递),而是采用类似二层交换机的方式对MPLS报文标签查询转发,因此不能通过IPSec协议保证其在公网投递的安全性。
