---
id: collect-261001-automatisation-infra/automatisation-infra/github-ramobis-nornir-l3-vpn-service-deployment-at-the-ost-eastern-switzerland-university--2
title: "github-ramobis-nornir-l3-vpn-service-deployment-at-the-ost-eastern-switzerland-university-of-applied"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: ["2024-04"]
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/github-ramobis-nornir-l3-vpn-service-deployment-at-the-ost-eastern-switzerland-university-of-applied.md
source_anchor: ""
source_lines: [159, 191]
sha256: 42f845ae4b32db5cb3db5fb065cd28a589ae2e088f0dfba610f5a4da5e9384f2
---

# github-ramobis-nornir-l3-vpn-service-deployment-at-the-ost-eastern-switzerland-university-of-applied

```
{
    "CustA": {
        "id": 10,
        "description": "Customer A",
        "route_import": ["10:0"],
        "route_export": ["10:0"],
    },
    "CustB": {
        "id": 20,
        "description": "Customer B",
        "route_import": ["20:0"],
        "route_export": ["20:0"],
    },
}
```
The datastructure below maps customer VRF names to dynamic/host specific service information. Currently it is just the required interface configuration. The example below is the configuration for routers located at site-A for both services CustA and CustB.

```
{
    "CustA": {
        "interfaces": [
            {"name": "GigabitEthernet4.10", "ip": "192.168.1.1 255.255.255.0"}
        ]
    },
    "CustB": {
        "interfaces": [
            {"name": "GigabitEthernet4.20", "ip": "192.168.21.2 255.255.255.0"}
        ]
    },
}
```
There will be no further development work on this project after the submission deadline on 19th of April 2024.
