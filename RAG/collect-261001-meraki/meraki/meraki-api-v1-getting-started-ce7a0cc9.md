---
id: collect-261001-meraki/meraki/meraki-api-v1-getting-started-ce7a0cc9
title: "meraki-api-v1-getting-started-ce7a0cc9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-getting-started-ce7a0cc9.md
source_anchor: ""
source_lines: [1, 54]
sha256: 7192d0e3a6f776d3364bffcccfb032f8e82fd6b0c7b7ab4df9daff4555531f69
---

# meraki-api-v1-getting-started-ce7a0cc9

CopySuccessful HTTP Status: 200
[
 {
  "mac": "00:11:22:33:44:55",
  "name": "My Switch 1",
  "network": {
   "id": "L_24329156"
  },
  "productType": "switch",
  "serial": "{serial}",
  "tags": [
   "example",
   "switch"
  ],
  "uplinks": [
   {
    "interface": "man1",
    "addresses": [
     {
      "protocol": "ipv4",
      "address": "10.0.1.2",
      "gateway": "10.0.1.1",
      "assignmentMode": "dynamic",
      "nameservers": {
       "addresses": [
        "208.67.222.222",
        "208.67.220.220"
       ]
      },
      "public": {
       "address": "78.11.19.49"
      }
     },
     {
      "protocol": "ipv6",
      "address": "2600:1700:ae0::c8ff:fe1e:12d2",
      "gateway": "fe80::fe1b:202a",
      "assignmentMode": "dynamic",
      "nameservers": {
       "addresses": [
        "::",
        "::"
       ]
      },
      "public": {
       "address": None
      }
     }
    ]
   }
  ]
 }
]
Copy[{'mac': '00:11:22:33:44:55', 'name': 'My Switch 1', 'network': {'id': 'L_24329156'}, 'productType': 'switch', 'serial': '{serial}', 'tags': ['example', 'switch'], 'uplinks': [{'interface': 'man1', 'addresses': [{'protocol': 'ipv4', 'address': '10.0.1.2', 'gateway': '10.0.1.1', 'assignmentMode': 'dynamic', 'nameservers': {'addresses': ['208.67.222.222', '208.67.220.220']}, 'public': {'address': '78.11.19.49'}}, {'protocol': 'ipv6', 'address': '2600:1700:ae0::c8ff:fe1e:12d2', 'gateway': 'fe80::fe1b:202a', 'assignmentMode': 'dynamic', 'nameservers': {'addresses': ['::', '::']}, 'public': {'address': None}}]}]}]
