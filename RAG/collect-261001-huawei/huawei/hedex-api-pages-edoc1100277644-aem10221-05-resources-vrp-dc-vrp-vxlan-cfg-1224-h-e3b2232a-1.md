---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100277644-aem10221-05-resources-vrp-dc-vrp-vxlan-cfg-1224-h-e3b2232a-1
title: "hedex-api-pages-edoc1100277644-aem10221-05-resources-vrp-dc-vrp-vxlan-cfg-1224-h-e3b2232a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100277644-aem10221-05-resources-vrp-dc-vrp-vxlan-cfg-1224-h-e3b2232a.md
source_anchor: ""
source_lines: [1, 242]
sha256: 7d5b7b9183adbd326c055902853bcce597db926bb784c12b219f306a8aa3604f
---

# hedex-api-pages-edoc1100277644-aem10221-05-resources-vrp-dc-vrp-vxlan-cfg-1224-h-e3b2232a

This section provides an example for configuring three-segment VXLAN tunnels to enable Layer 2 communication between VMs that belong to the different DCs.
On the network shown in Figure 1, BGP EVPN is configured within DC A and DC B to establish VXLAN tunnels. BGP EVPN is also configured on Leaf2 and Leaf3 to establish a VXLAN tunnel between them. To enable communication between VM1 and VM2, implement Layer 2 interworking between DC A and DC B. In this example, the VXLAN tunnel in DC A uses VNI 10, and that in DC B uses VNI 20. VNI conversion must be performed before a VXLAN tunnel is established between Transit Leaf1 and Transit Leaf2.
In this example, interface1 and interface2 represent GE0/1/0 and GE0/1/8, respectively.
| Table 1 Interface IP addresses |  |  |  |  |  | 
|---|---|---|---|---|---|
| Device | Interface | IP Address | Device | Interface | IP Address | 
|---|---|---|---|---|---|
| Spine1 | GE0/1/0 | 192.168.10.1/24 | Spine2 | GE0/1/0 | 192.168.30.1/24 | 
|  | GE0/1/8 | 192.168.20.1/24 |  | GE0/1/8 | 192.168.40.1/24 | 
| Leaf1 | GE0/1/0 | 192.168.10.2/24 | Leaf4 | GE0/1/0 | 192.168.40.2/24 | 
|  | GE0/1/8 | - |  | GE0/1/8 | - | 
|  | LoopBack1 | 1.1.1.1/32 |  | LoopBack1 | 4.4.4.4/32 | 
| Leaf2 | GE0/1/0 | 192.168.20.2/24 | Leaf3 | GE0/1/0 | 192.168.30.2/24 | 
|  | GE0/1/8 | 192.168.50.1/24 |  | GE0/1/8 | 192.168.50.2/24 | 
|  | LoopBack1 | 2.2.2.2/32 |  | LoopBack1 | 3.3.3.3/32 | 
The configuration roadmap is as follows:
Assign an IP address to each interface.
Configure an IGP to allow devices to communicate with each other.
Configure static routes to achieve interworking between DCs.
Configure BGP EVPN within DC A and DC B to establish VXLAN tunnels.
Configure BGP EVPN on Leaf2 and Leaf3 to establish a VXLAN tunnel between them.
Configure Leaf2 and Leaf3 to advertise routes that are re-originated by the EVPN address family to BGP EVPN peers.
To complete the configuration, you need the following data:
VLAN IDs of the VMs
BD IDs
VNI IDs associated with BDs within DC A and DC B
Number of the AS to which DC A and DC B belong
Name of the SHG to which Leaf2 and Leaf3 belong
For details about how to configure interface IP addresses and masks, see Configuration Files.
For detailed configurations, see Configuration Files.
Run the display vxlan tunnel command on each leaf node to view information about the VXLAN tunnels. The following example uses the command output on Leaf2. The command output shows that the VXLAN tunnels are up.
[~Leaf2] display vxlan tunnel
Number of vxlan tunnel : 2
Tunnel ID   Source                Destination           State  Type     Uptime
-----------------------------------------------------------------------------------
4026531924  2.2.2.2               1.1.1.1               up     dynamic  00:39:19  
4026531925  2.2.2.2               3.3.3.3               up     dynamic  00:39:09 
Run the display vxlan peer command on Leaf2 to check information about VXLAN peers.
[~Leaf2] display vxlan peer
Number of peers : 2
Vni ID    Source                  Destination            Type      Out Vni ID
-------------------------------------------------------------------------------
10        2.2.2.2                 1.1.1.1                dynamic   10         
10        2.2.2.2                 3.3.3.3                dynamic   20
After the preceding configurations are complete, Layer 2 communication can be implemented between VM 1 and VM 2.
Spine 1 configuration file
#
sysname Spine1
#
interface GE0/1/0
 undo shutdown  
 ip address 192.168.10.1 255.255.255.0
#               
interface GE0/1/8
 undo shutdown  
 ip address 192.168.20.1 255.255.255.0
#               
ospf 1          
 area 0.0.0.0   
  network 192.168.10.0 0.0.0.255
  network 192.168.20.0 0.0.0.255
#               
return 
Leaf1 configuration file
#
sysname Leaf1
#
evpn vpn-instance evrf1 bd-mode
 route-distinguisher 10:1
 vpn-target 11:1 export-extcommunity
 vpn-target 11:1 import-extcommunity
#
bridge-domain 10
 vxlan vni 10 split-horizon-mode
 evpn binding vpn-instance evrf1
#               
interface GE0/1/0
 undo shutdown  
 ip address 192.168.10.2 255.255.255.0
#               
interface GE0/1/8
 undo shutdown       
#               
interface GE0/1/8.1 mode l2
 encapsulation dot1q vid 10
 rewrite pop single
 bridge-domain 10
#               
interface LoopBack1
 ip address 1.1.1.1 255.255.255.255
#               
interface Nve1  
 source 1.1.1.1 
 vni 10 head-end peer-list protocol bgp
#               
bgp 100
 peer 2.2.2.2 as-number 100
 peer 2.2.2.2 connect-interface LoopBack1
 #
 ipv4-family unicast
  peer 2.2.2.2 enable
 #              
 l2vpn-family evpn
  undo policy vpn-target
  peer 2.2.2.2 enable
  peer 2.2.2.2 advertise encap-type vxlan
#               
ospf 1          
 area 0.0.0.0   
  network 1.1.1.1 0.0.0.0
  network 192.168.10.0 0.0.0.255
#               
return
Leaf2 configuration file
#
sysname Leaf2
#
evpn vpn-instance evrf1 bd-mode
 route-distinguisher 10:1
 vpn-target 11:1 export-extcommunity
 vpn-target 11:1 import-extcommunity
#
bridge-domain 10
 vxlan vni 10 split-horizon-mode
 evpn binding vpn-instance evrf1
#                         
interface GE0/1/0
 undo shutdown  
 ip address 192.168.20.2 255.255.255.0
#               
interface GE0/1/8
 undo shutdown  
 ip address 192.168.50.1 255.255.255.0
#               
interface LoopBack1
 ip address 2.2.2.2 255.255.255.255
#               
interface Nve1  
 source 2.2.2.2 
 vni 10 head-end peer-list protocol bgp
#               
bgp 100
 peer 1.1.1.1 as-number 100
 peer 1.1.1.1 connect-interface LoopBack1
 peer 3.3.3.3 as-number 200
 peer 3.3.3.3 ebgp-max-hop 255
 peer 3.3.3.3 connect-interface LoopBack1
 #
 ipv4-family unicast
  network 2.2.2.2 255.255.255.255
  peer 1.1.1.1 enable
  peer 3.3.3.3 enable
 #              
 l2vpn-family evpn
  undo policy vpn-target
  peer 1.1.1.1 enable
  peer 1.1.1.1 advertise encap-type vxlan
  peer 1.1.1.1 import reoriginate
  peer 1.1.1.1 advertise route-reoriginated evpn mac
  peer 3.3.3.3 enable
  peer 3.3.3.3 advertise encap-type vxlan
  peer 3.3.3.3 import reoriginate
  peer 3.3.3.3 advertise route-reoriginated evpn mac
  peer 3.3.3.3 split-group sg1
#
ospf 1          
 area 0.0.0.0   
  network 2.2.2.2 0.0.0.0
  network 192.168.20.0 0.0.0.255
#
ip route-static 3.3.3.3 255.255.255.255 192.168.50.2
#               
return  
Spine 2 configuration file
#
sysname Spine2
#
interface GE0/1/0
 undo shutdown  
 ip address 192.168.30.1 255.255.255.0
#               
interface GE0/1/8
 undo shutdown  
 ip address 192.168.40.1 255.255.255.0
#               
ospf 1          
 area 0.0.0.0   
  network 192.168.30.0 0.0.0.255
  network 192.168.40.0 0.0.0.255
#               
return
Leaf3 configuration file
#
sysname Leaf3
#
evpn vpn-instance evrf1 bd-mode
 route-distinguisher 10:1
 vpn-target 11:1 export-extcommunity
 vpn-target 11:1 import-extcommunity
#
bridge-domain 10
 vxlan vni 20 split-horizon-mode
 evpn binding vpn-instance evrf1
#               
interface GE0/1/0
 undo shutdown  
 ip address 192.168.30.2 255.255.255.0
#               
interface GE0/1/8
 undo shutdown  
 ip address 192.168.50.2 255.255.255.0
#               
interface LoopBack1
 ip address 3.3.3.3 255.255.255.255
#               
interface Nve1  
 source 3.3.3.3 
 vni 20 head-end peer-list protocol bgp
#               
bgp 200
 peer 2.2.2.2 as-number 100
 peer 2.2.2.2 ebgp-max-hop 255
 peer 2.2.2.2 connect-interface LoopBack1
 peer 4.4.4.4 as-number 200
 peer 4.4.4.4 connect-interface LoopBack1
 #
 ipv4-family unicast
  network 3.3.3.3 255.255.255.255
  peer 2.2.2.2 enable
  peer 4.4.4.4 enable
 #              
 l2vpn-family evpn
  undo policy vpn-target
  peer 2.2.2.2 enable
  peer 2.2.2.2 advertise encap-type vxlan
  peer 2.2.2.2 import reoriginate
  peer 2.2.2.2 advertise route-reoriginated evpn mac
  peer 2.2.2.2 split-group sg1
  peer 4.4.4.4 enable
  peer 4.4.4.4 advertise encap-type vxlan
  peer 4.4.4.4 import reoriginate
  peer 4.4.4.4 advertise route-reoriginated evpn mac
