---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c-2
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c.md
source_anchor: ""
source_lines: [159, 210]
sha256: 9d25057dcba6a2161f1ab5a73bf645a60ad46e76cb9e01ef2acf0e3f4d9c30be
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c

Configuring VLANs and SVI
                           vlan 10
vn-segment 10010
vlan 101
vn-segment 10101
interface Vlan101
no shutdown
mtu 9216
vrf member vxlan-10101
no ip redirects
ip forward
ipv6 address use-link-local-only
no ipv6 redirects
interface vlan10
no shutdown
mtu 9216
vrf member vxlan-10101
no ip redirects
ip address 192.0.2.102/24
ipv6 address 2001:DB8:0:1::1/64
no ipv6 redirects
fabric forwarding mode anycast-gateway
                           
                           
                           
                           
                        
                           Configuring Virtual Port Channel
                           
                           interface Ethernet1/3
switchport
switchport mode trunk
channel-group 100
no shutdown
exit
interface Ethernet1/39
switchport
switchport mode trunk
channel-group 101
no shutdown
interface Ethernet1/46
switchport
switchport mode trunk
channel-group 102
no shutdown
interface port-channel100 
vpc 100
interface port-channel101
vpc 101 
interface port-channel102
vpc 102
exit
