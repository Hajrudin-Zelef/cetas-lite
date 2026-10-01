---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ospfv3-0096-html-695025ce-2
title: "Configure RouterA."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ospfv3-0096-html-695025ce.md
source_anchor: ""
source_lines: [182, 241]
sha256: 3ff411a50cb42b8205d1fbecb02eaca06a04f968c1153dde90b943e6cde0a6af
---

# Configure RouterA.

 ipv6 address 2001:DB8:1::3/64
 ospfv3 1 area 0.0.0.0
#
interface gigabitethernet1/0/1
 ipv6 enable
 ipv6 address 2001:DB8:3::1/64
 ospfv3 1 area 0.0.0.0
#
return
Configuration file of RouterB
#
 sysname RouterB
#
 ipv6
#
 bfd
#
ospfv3 1
 router-id 2.2.2.2
 bfd all-interfaces enable
 bfd all-interfaces min-transmit-interval 100 min-receive-interval 100 detect-multiplier 4
#
interface gigabitethernet1/0/0
 ipv6 enable
 ipv6 address 2001:DB8:1::2/64
 ospfv3 1 area 0.0.0.0
#
interface gigabitethernet1/0/1
 ipv6 enable
 ipv6 address 2001:DB8:2::1/64
 ospfv3 1 area 0.0.0.0
#
interface gigabitethernet1/0/2
 ipv6 enable
 ipv6 address 2001:DB8:4::1/64
 ospfv3 1 area 0.0.0.0
#
return
Configuration file of RouterC
#
 sysname RouterC
#
 ipv6
#
ospfv3 1
 router-id 3.3.3.3
 bfd all-interfaces enable
 bfd all-interfaces min-transmit-interval 100 min-receive-interval 100 detect-multiplier 4
#
interface gigabitethernet1/0/0
 ipv6 enable
 ipv6 address 2001:DB8:2::2/64
 ospfv3 1 area 0.0.0.0
#
interface gigabitethernet1/0/1
 ipv6 enable
 ipv6 address 2001:DB8:3::3/64
 ospfv3 1 area 0.0.0.0
#
return
