---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9-5
title: "c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9.md
source_anchor: ""
source_lines: [227, 233]
sha256: 960304a9c0aa67d0894a23bb2c9474be371e3dfc303e78526fc21312baf2a7a1
---

# c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9

| Object groups | 7.0(1) | Object groups simplify ACL creation and maintenance. We introduced or modified the following commands: object-group protocol , object-group network , object-group service , object-group icmp_type . | 
| Regular expressions and policy maps | 7.2(1) | Regular expressions and policy maps were introduced to be used under inspection policy maps. The following commands were introduced: class-map type regex , regex , match regex . | 
| Objects | 8.3(1) | Object support was introduced. We introduced or modified the following commands: object-network , object-service , object-group network , object-group service , network object , access-list extended , access-list webtype , access-list remark . | 
| User Object Groups for Identity Firewall | 8.4(2) | User object groups for identity firewall were introduced. We introduced the following commands: object-network user , user . | 
| Mixed IPv4 and IPv6 network object groups | 9.0(1) | Previously, network object groups could only contain all IPv4 addresses or all IPv6 addresses. Now network object groups can support a mix of both IPv4 and IPv6 addresses. Note You cannot use a mixed object group for NAT. We modified the following commands: object-group network . | 
| Security Group Object Groups for Cisco TrustSec | 8.4(2) | Security group object groups for TrustSec were introduced. We introduced the following commands: object-network security , security . | 
| Extended ACLand object enhancement to filter ICMP traffic by ICMP code | 9.0(1) | ICMP traffic can now be permitted/denied based on ICMP code. We introduced or modified the following commands: access-list extended , service-object , service . |
