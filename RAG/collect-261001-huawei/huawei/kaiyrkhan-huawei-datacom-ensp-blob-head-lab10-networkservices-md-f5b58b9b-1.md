---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-networkservices-md-f5b58b9b-1
title: "Create VLANs"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab10-networkservices-md-f5b58b9b.md
source_anchor: ""
source_lines: [1, 321]
sha256: b21411e573005552f1edef0ec559bfa896ebce26e76ebf86bf73e4859b9c9518
---

# Create VLANs

Download Link for eNSP Topology File
- Configure VLAN (Create VLANs and Access Port, Trunk Port)
 LACP Link Aggregation. Eth-Trunk
MSTP (Multiple Spanning Tree Protocol)
- VRRP (Virtual Router Redundancy Protocol)
- Single-Area OSPF
- DHCP
- NAT (Easy IP)
- Remote Access (SSH, Telnet)
- DNS and HTTP
- FTP
- TFTP
- NTP
undo terminal monitor
system-view
sysname A1
Create VLANs
vlan batch 111 112 50
vlan 111
 description Service VLAN
 quit
vlan 112
 description Service VLAN
 quit
vlan 50
 description MGMT VLAN
 quit
display vlan
Configure Access Port
interface Ethernet0/0/1
 port link-type access
 port default vlan 111
 quit
interface Ethernet0/0/3
 port link-type access
 port default vlan 111
 quit
interface Ethernet0/0/2
 port link-type access
 port default vlan 112
 quit
display port vlan
Configure Trunk Port and Allowed VLANs
interface g0/0/1
 port link-type trunk
 port trunk allow-pass vlan 111 112 50
 quit
interface g0/0/2
 port link-type trunk
 port trunk allow-pass vlan 111 112 50
 quit
display port vlanundo terminal monitor
system-view
sysname D1
Create VLANs
vlan batch 111 112 50
vlan 111
 description Service VLAN
 quit
vlan 112
 description Service VLAN
 quit
vlan 50
 description MGMT VLAN
 quit
display vlan
Configure Trunk Port and Allowed VLANs
interface g0/0/1
 port link-type trunk
 port trunk allow-pass vlan 111 112 50
 quit
interface g0/0/2
 port link-type trunk
 port trunk allow-pass vlan 111 112 50
 quit
display port vlan
Configure LACP Link Aggregation
interface Eth-Trunk 1                                          // Create Eth-Trunk
 port link-type trunk                                          // Trunk Port
 port trunk allow-pass vlan 111 112 50                         // Allowed VLANs         
 mode lacp-static                                              // Link Aggregation Mode
 quitinterface Eth-Trunk 1
 port link-type trunk
 port trunk allow-pass vlan 111 112 50        
 mode lacp-static
 quit
display port vlan
display eth-trunk 1
Add a Port to the Eth-Trunk
interface g0/0/3
 eth-trunk 1
 quit
interface g0/0/4
 eth-trunk 1
 quit
display int brief# Verify Configuration
display eth-trunk 1
Configure MSTP
display stpstp region-configuration
 region-name HQ1
 revision-level 1
 instance 1 vlan 111
 instance 2 vlan 112
 instance 3 vlan 50
 active region-configuration
 check region-configuration
 quitdisplay cu | begin stp# D1 Switch
stp instance 1 root primary
stp instance 3 root primary
stp instance 2 root secondary# D2 Switch
stp instance 2 root primary
stp instance 1 root secondary
stp instance 3 root secondary
Configure MSTP
stp region-configuration
 region-name HQ1
 revision-level 1
 instance 1 vlan 111
 instance 2 vlan 112
 instance 3 vlan 50
 active region-configuration
 check region-configuration
 quitdisplay cu | begin stp
Verify Configuration
display stp vlan 111
display stp vlan 112
display stp vlan 50
немесе
display stp instance 1 brief
display stp instance 2 brief
display stp instance 3 brief
D1 Switch
interface vlanif 111
 ip address 172.16.111.1 24
 vrrp vrid 111 virtual-ip 172.16.111.254
 vrrp vrid 111 priority 105
 quit
interface vlanif 112
 ip address 172.16.112.1 24
 vrrp vrid 112 virtual-ip 172.16.112.254
 quit
interface vlanif 50
 ip address 10.1.50.1 24
 vrrp vrid 50 virtual-ip 10.1.50.254
 vrrp vrid 50 priority 105
 quit
display ip int brief
display vrrp brief
D2 Switch
interface vlanif 111
 ip address 172.16.111.2 24
 vrrp vrid 111 virtual-ip 172.16.111.254
 quit
interface vlanif 112
 ip address 172.16.112.2 24
 vrrp vrid 112 virtual-ip 172.16.112.254
 vrrp vrid 112 priority 105
 quit
interface vlanif 50
 ip address 10.1.50.2 24
 vrrp vrid 50 virtual-ip 10.1.50.254
 quit
display ip int brief
display vrrp brief
D1 Switch
# Create VLANs
vlan 4
 quit
display vlan
# Configure Access Port
interface GigabitEthernet0/0/5
 port link-type access
 port default vlan 4
display port vlan
# Create VLANIF interface
interface vlanif 4
 ip address 10.1.1.106 30
 quit# Create Loopback interface
interface Loopback 50
 ip address 50.3.3.3 32
 quitdisplay ip int briefospf 1 router-id 50.3.3.3
 area 0
 network 10.1.1.104 0.0.0.3
 network 172.16.111.0 0.0.0.255
 network 172.16.112.0 0.0.0.255
 network 10.1.50.0 0.0.0.255
 network 50.3.3.3 0.0.0.0
 quit
 quit
display cu | begin ospf
D2 Switch
# Create VLANs
vlan 8
 quit
display vlan
# Configure Access Port
interface GigabitEthernet0/0/5
 port link-type access
 port default vlan 8
display port vlan
# Create VLANIF interface
interface vlanif 8
 ip address 10.1.1.110 30
 quit# Create Loopback interface
interface Loopback 50
 ip address 50.4.4.4 32
 quitdisplay ip int briefospf 1 router-id 50.4.4.4
 area 0
 network 10.1.1.108 0.0.0.3
 network 172.16.111.0 0.0.0.255
 network 172.16.112.0 0.0.0.255
 network 10.1.50.0 0.0.0.255
 network 50.4.4.4 0.0.0.0
 quit
 quit
display ospf peer brief
C1 Switch
undo terminal monitor
system-view
sysname C1interface g0/0/0
 ip address 10.1.1.102 30
 quit
interface g0/0/1
 ip address 10.1.1.105 30
 quit
interface g0/0/2
 ip address 10.1.1.109 30
 quit
interface Loopback 50
 ip address 50.2.2.2 32
 quit
display ip int briefdisplay ip int brief
ospf 1 router-id 50.2.2.2
 area 0
 network 10.1.1.100 0.0.0.3
 network 10.1.1.104 0.0.0.3
 network 10.1.1.108 0.0.0.3
 network 50.2.2.2 0.0.0.0
 quit
 quit
display ospf peer brief
EdgeR1 Router
undo terminal monitor
system-view
sysname EdgeR1interface g0/0/0
 ip address 10.1.1.101 30
 quit
interface g0/0/2
 ip address 172.16.128.1 24
 quit
interface g0/0/1
 ip address 192.168.137.254 24
 quit
interface Loopback 50
 ip address 50.1.1.1 32
 quit
display ip int briefping 192.168.137.1
 Request time out
Windows+R ➜ Turn off Windows Defender Firewall
ping 192.168.137.1
 Reply from 192.168.137.1: bytes=56 Sequence=2 ttl=128 time=10 msdisplay ip int brief
ospf 1 router-id 50.1.1.1
 area 0
 network 10.1.1.100 0.0.0.3
 network 172.16.128.0 0.0.0.255
 network 50.1.1.1 0.0.0.0
 quit
 quit
display ospf peer brief
DHCP Router
undo terminal monitor
system-view
sysname DHCPinterface g0/0/0
 ip address 172.16.128.67 24
 quit
interface Loopback 50
 ip address 50.5.5.5 32
 quit
display ip int briefdisplay ip int brief
ospf 1 router-id 50.5.5.5
 area 0
 network 172.16.128.0 0.0.0.255
 network 50.5.5.5 0.0.0.0
 quit
 quit
display ospf peer briefdhcp enable
ip pool VLAN111
 network 172.16.111.0 mask 24
 gateway-list 172.16.111.254
 dns-list 8.8.8.8
 excluded-ip-address 172.16.111.1 172.16.111.10
 excluded-ip-address 172.16.111.251 172.16.111.253
 lease day 5
 quit
ip pool VLAN112
 network 172.16.112.0 mask 24
 gateway-list 172.16.112.254
 dns-list 172.16.128.53
 excluded-ip-address 172.16.112.1 172.16.112.10
 excluded-ip-address 172.16.112.251 172.16.112.253
 lease day 5
 quit
interface g0/0/0
 dhcp select global
 quit
Verify Configuration
display ip pool
display ip pool name VLAN111
display dhcp server statistics
DHCP Relay Agent (D1 and D2 Switch)
dhcp enable
interface vlanif 111
 dhcp select relay
 dhcp relay server-ip 172.16.128.67
 quit
interface vlanif 112
 dhcp select relay
 dhcp relay server-ip 172.16.128.67
 quit<DHCP> display dhcp server statistics
DHCP Server Statistics:
 
