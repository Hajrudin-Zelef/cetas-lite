---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-netstream-000421-26e1fc78
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-netstream-000421--26e1fc78"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-netstream-000421--26e1fc78.md
source_anchor: ""
source_lines: [1, 8]
sha256: b93ab7c5dcc76a3f22f232ec6b51f6e9e0a1b2aaecc084351942cc3e2994846b
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-netstream-000421--26e1fc78

You can configure the function of restoring logs in NetStream traffic statistics to save the information about all aging flow entries. You can also search for the saved flow entries based on the seven attributes (including the source IP address, destination IP address, source port number, destination port number, protocol number, start time, and end time).
Only the AR6140-16G4XG, AR6140H-S, AR6280K, AR6280-S, AR6300-S, AR6280, AR6300K, and AR6300 support this function.
The system view is displayed.
The path for storing logs in NetStream traffic statistics is configured.
By default, the path for storing logs is not configured.
The function of recording logs in NetStream traffic statistics is enabled globally.
By default, the function of recording logs is disabled.
Run the display ip netstream cache classify [ source-ip source-ip | destination-ip destination-ip | source-port source-port | destination-port destination-port | protocol protocol-id | start-time start-time | end-time end-time ] * | [ verbose ] command to displays flow information in the NetStream cache.
