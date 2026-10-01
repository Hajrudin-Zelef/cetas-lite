---
id: collect-261001-cisco/cisco/enterprise-en-multi-level-m-lag-vrrp-configuration-thread-667222965948923904-667-26b2aa6d
title: "enterprise-en-multi-level-m-lag-vrrp-configuration-thread-667222965948923904-667-26b2aa6d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-multi-level-m-lag-vrrp-configuration-thread-667222965948923904-667-26b2aa6d.md
source_anchor: ""
source_lines: [1, 44]
sha256: 83088ee6065fd05395d6b3b62c797c9d19ca9d7b907ba6dc6011bb2ff6d87fd9
---

# enterprise-en-multi-level-m-lag-vrrp-configuration-thread-667222965948923904-667-26b2aa6d

Multi-level M-LAG
After M-LAG is deployed between SW1 and SW2, M-LAG is deployed between SW3 and SW4. The two M-LAGs are connected. This deployment simplifies networking and allows more servers to be connected to the network in dual-homing mode. Before deploying multi-level M-LAG, configure Virtual Spanning Tree Protocol (V-STP).
The virtual ip will be work as the gateway of the server.
Multi-level M-LAG + VRRP configuration（CE device as example）
SW1:
interface eth-trunk x //connect to SW3 and SW4
mode lacp-static
port link-type trunk
port trunk allow-pass vlan x
trunkport 10ge x/0/a to x/0/b
dfs-group 1 m-lag 2
//no vrrp configuration others is same with Single-level M-LAG
SW2:
SW3:
stp mode rstp
stp root primary
stp bridge-address xxx-xxx-xxx // the MAC address of the DFS master device
stp v-stp enable
#
ip vpn-instance management
ipv4-family
route-distinguisher 100:1
vpn-target 111:1 both
interface Meth 0/0/0
ip binding vpn-instance management
ip address x.x.x.x x
dfs-group 1
source ip x.x.x.x //ip of meth0/0/0
priority 150
interface eth-trunk x //to SW4
trunkport 10ge 1/0/a to 1/0/b
peer-link 1
vlan x //service vlan
interface eth-trunk x //connect to SW1 and SW2
dfs-group 1 m-lag 1
interface vlanif x
ip address x.x.x.2 x
vrrp vrid 1 virtual-ip x.x.x.1
commit
SW4:
priority 120
interface eth-trunk x //to SW3
interface eth-trunk x //to SW1 and SW2
ip address x.x.x.3 x
