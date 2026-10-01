---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-eth-trunk-mstp-md-a062f044
title: "kaiyrkhan-huawei-datacom-ensp-blob-head-eth-trunk-mstp-md-a062f044"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-eth-trunk-mstp-md-a062f044.md
source_anchor: ""
source_lines: [1, 37]
sha256: 1ed674ce533e6e7909b041d82d0c1282ee3cef4b8dac534f73442e2435db2c13
---

# kaiyrkhan-huawei-datacom-ensp-blob-head-eth-trunk-mstp-md-a062f044

D1 and D2 Switch
vlan batch 10 20
display vlan
Step1: Eth-Trunk конфигурациялау
interface Eth-Trunk1
 mode lacp-static
 port link-type trunk
 port trunk allow-pass vlan 10 20
Step2: Физикалық порттарды Eth-Trunk-қа қосу
interface g0/0/1
 eth-trunk 1
interface g0/0/2
 eth-trunk 1interface Eth-Trunk1
load-balance src-dst-mac
Step3: MSTP instance mapping
stp mode mstp
stp region-configuration
 region-name HQ
 revision-level 1
 instance 1 vlan 10
 instance 2 vlan 20
 active region-configuration
Root placement (load-balancing)
D1
stp instance 1 root primary
stp instance 2 root secondary
D2
stp instance 2 root primary
stp instance 1 root secondary
Step4: Тексеру (Verification)
display eth-trunk 1
display stp brief
display stp instance 1 brief
display stp instance 2 brief
display stp vlan 10
display stp vlan 20display interface eth-trunk 1
display lacp statistics eth-trunk 1
