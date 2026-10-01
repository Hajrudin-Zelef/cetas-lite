---
id: collect-261001-cisco/cisco/ansible-policybook-examples-policy-network-cisco-readme-md-at-1f98566ae2b5b888aede4455031b
title: "ansible-policybook-examples-policy-network-cisco-readme-md-at-1f98566ae2b5b888aede4455031b29315aac00"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/ansible-policybook-examples-policy-network-cisco-readme-md-at-1f98566ae2b5b888aede4455031b29315aac00.md
source_anchor: ""
source_lines: [1, 23]
sha256: 3f0634b01368b82d67de4c22bfe4e73d2519e71f157d33a950e773d560f5c380
---

# ansible-policybook-examples-policy-network-cisco-readme-md-at-1f98566ae2b5b888aede4455031b29315aac00

Here is a sample playbook, let's call it "cisco_ios.yml", which can be used to demonstrate policy checks:

```
---
- name: Configure Cisco IOS
  cisco.ios.ios_config:
    lines:
      - hostname foobar
      - line vty 0
      - password password123
      - no login
      - set ntp server 127.0.0.1
```
```
$ ansible-policy -p examples/check_project/cisco_ios.yml --policy-dir ansible-policybook-examples/policy/network/cisco/
TASK [Configure Cisco IOS] examples/check_project/cisco_ios.yml L2-10 ************************************************************
... Check_on_IOS_configuration Not Validated
    WARNING: Cisco IOS settings violation
--------------------------------------------------------------------------------------------------------------------------------------------
SUMMARY
... Total files: 1, Validated: 0, Not Validated: 1
Violations are detected! in 1 task
```
