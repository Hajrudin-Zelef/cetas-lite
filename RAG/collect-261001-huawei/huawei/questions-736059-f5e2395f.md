---
id: collect-261001-huawei/huawei/questions-736059-f5e2395f
title: "questions-736059-f5e2395f"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-736059-f5e2395f.md
source_anchor: ""
source_lines: [1, 4]
sha256: 93f2321c683759a427e8ec4f53064e5c06ed16dd4e192c350109a4abeed1edc2
---

# questions-736059-f5e2395f

I'm newbie network enginner wondering is there a way to assign same vlan to multiple ports rapidly (not separately)? Somehow like this: "int eth 0/0/1 0/0/9 0/0/15 default vlan 101"
1 Answer 1
I have not used that particular brand of switch before but practically every other vendor has an interface range command. For instance on Cisco I can do int range Gi0/1-5 and will then be in interface configuration mode for all those ports simultaneously where I could then do switchport access vlan xxx.
If you search any manual you might have for that switch for "range" you should hit on it - if supported.
