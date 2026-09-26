---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-32-4
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-32.md
source_anchor: ""
source_lines: [371, 390]
sha256: 0eefe2d96949bd7532cf772f6706d5bb5e66b60c6abc995f76d87606e0f2a5fb
---

# Overview

Sub-menu shows label bindings for routes received from other routers. Static mapping can be configured if there is no intention to use LDP dynamically. This table is used to build Forwarding Table

**Properties**

| Property | Description | 
|---|---|
| **comments** (*string* ; Default: ) | Short description of the entry | 
| **disabled** (*yes \| no* ; Default:**no** ) |  | 
| **dst-address** (*IP/Mask* ; Default: ) | Destination prefix the label is assigned to. | 
| **label** (*integer[0..1048576] \| alert \| expl-null \| expl-null6 \| impl-null \| none* ; Default: ) | Label number assigned to destination. | 
| **nexthop (***IP* ; Default:**)** |  | 
| **vrf (***name* ; Default: main**)** | Name of the VRF table this mapping belongs to. | 

**Read-only Properties**

| Property | Description | 
|---|---|
| **inactive** (*yes \| no* ) | Whether binding is active and can be selected as a candidate for forwarding. | 
| **dynamic** (*yes \| no* ) | Whether entry was dynamically added | 
| **path** (*string* ) |  |
