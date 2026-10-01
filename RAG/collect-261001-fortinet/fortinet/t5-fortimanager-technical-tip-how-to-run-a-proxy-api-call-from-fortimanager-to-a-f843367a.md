---
id: collect-261001-fortinet/fortinet/t5-fortimanager-technical-tip-how-to-run-a-proxy-api-call-from-fortimanager-to-a-f843367a
title: "t5-fortimanager-technical-tip-how-to-run-a-proxy-api-call-from-fortimanager-to-a-f843367a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/t5-fortimanager-technical-tip-how-to-run-a-proxy-api-call-from-fortimanager-to-a-f843367a.md
source_anchor: ""
source_lines: [1, 58]
sha256: e591c2711c3bc73f29b66defcf0e84e0870f01440847bee28848af604a2a6796
---

# t5-fortimanager-technical-tip-how-to-run-a-proxy-api-call-from-fortimanager-to-a-f843367a

Technical Tip: How to run a proxy API call from FortiManager to a managed FortiGate
Description
This article describes how to run a proxy API call from FortiManager to a managed FortiGate to collect data.
Scope
FortiManager and FortiGate.
Solution
- FortiManager will send proxy API calls via url:"sys/proxy/json" to managed FortiGate API "resource": "<FortiGate API Call>" via the FGFM tunnel that established between the FortiManager and the managed FortiGate.
Note:
For a complete list of FortiGate API calls, refer to Fortinet Development Network (FNDN):
{
     "id" : "1",
     "method": "exec",
     "params": [ 
         {
             "url": "sys/proxy/json",
             "data": {
                   "target": [ "device/<device-name>" ],
                   "action": "get",
                   "resource": "<FortiGate API Call>"
             } 
         } 
     ],
     "session" : "<session-id>"
}
- Below are two sample usages of proxy calls from FortiManager to a managed FortiGate:
- Get FortiGate Firewall Address:
{
     "id" : "1",
     "method": "exec",
     "params": [ 
         {
             "url": "sys/proxy/json",
             "data": {
                   "target": [ "device/lab-fgt2" ],
                   "action": "get",
                   "resource": "/api/v2/cmdb/firewall/address"
             } 
         } 
     ],
     "session" : "{{session-id}}"
}
- Get FortiGate License Info:
{
     "id" : "1",
     "method": "exec",
     "params": [ 
         {
             "url": "sys/proxy/json",
             "data": {
                   "target": [ "device/lab-fgt2" ],
                   "action": "get",
                   "resource": "/api/v2/monitor/license/status"
             } 
         } 
     ],
     "session" : "{{session-id}}"
}
Related articles:
