---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0041-html-f7f812db-2
title: "Configure PE1. The configurations of PE2 and P are similar to the configuration of PE1, and are not mentioned here."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0041-html-f7f812db.md
source_anchor: ""
source_lines: [84, 289]
sha256: df7a050ac3814fdbb0437aa38b08aa78c62e42291fd4714c4e57b429ee6456da
---

# Configure PE1. The configurations of PE2 and P are similar to the configuration of PE1, and are not mentioned here.

  --- 10.2.1.2 ping statistics ---                            
    5 packet(s) transmitted                                   
    5 packet(s) received                                      
    0.00% packet loss                                         
    round-trip min/avg/max = 1/1/1 ms
[PE1] tunnel-policy gre1
[PE1-tunnel-policy-gre1] tunnel select-seq gre load-balance-number 1
[PE1-tunnel-policy-gre1] quit
[PE1] interface gigabitethernet 1/0/0
[PE1-GigabitEthernet1/0/0] mpls l2vc 10.10.2.1 39 tunnel-policy gre1
[PE1-GigabitEthernet1/0/0] quit
[PE2] tunnel-policy gre1
[PE2-tunnel-policy-gre1] tunnel select-seq gre load-balance-number 1
[PE2-tunnel-policy-gre1] quit
[PE2] interface gigabitethernet 2/0/0
[PE2-GigabitEthernet2/0/0] mpls l2vc 10.10.1.1 39 tunnel-policy gre1
[PE2-GigabitEthernet2/0/0] quit
# After the configurations are complete, check the L2VPN connection on PEs. You can see that an L2VC connection has been set up and is in Up state.
[PE1] display mpls l2vc interface gigabitethernet 1/0/0
 *client interface       : GigabitEthernet1/0/0 is up
  Administrator PW       : no
  session state          : up
  AC status              : up
  Ignore AC state        : disable
  VC state               : up
  Label state            : 0
  Token state            : 0
  VC ID                  : 39
  VC type                : Ethernet
  destination            : 10.10.2.1
  local group ID         : 0            remote group ID      : 0
  local VC label         : 1025         remote VC label      : 1024
  local AC OAM State     : up
  local PSN OAM State    : up
  local forwarding state : forwarding
  local status code      : 0x0
  remote AC OAM state    : up
  remote PSN OAM state   : up
  remote forwarding state: forwarding
  remote status code     : 0x0
  ignore standby state   : no
  BFD for PW             : unavailable
  VCCV State             : up
  manual fault           : not set
  active state           : active
  forwarding entry       : exist
  link state             : up
  local VC MTU           : 1500         remote VC MTU        : 1500
  local VCCV             : alert ttl lsp-ping bfd 
  remote VCCV            : alert ttl lsp-ping bfd 
  local control word     : disable      remote control word  : disable
  tunnel policy name     : gre1
  PW template name       : --
  primary or secondary   : primary
  load balance type      : flow
  Access-port            : false
  Switchover Flag        : false
  VC tunnel/token info   : 1 tunnels/tokens
    NO.0  TNL type       : gre   , TNL ID : 0x2
    Backup TNL type      : lsp   , TNL ID : 0x0
  create time            : 0 days, 2 hours, 37 minutes, 1 seconds
  up time                : 0 days, 0 hours, 2 minutes, 11 seconds
  last change time       : 0 days, 0 hours, 2 minutes, 11 seconds
  VC last up time        : 2013/02/20 18:58:24
  VC total up time       : 0 days, 2 hours, 35 minutes, 58 seconds
  CKey                   : 2
  NKey                   : 1
  PW redundancy mode     : frr
  AdminPw interface      : --
  AdminPw link state     : --
  Diffserv Mode          : uniform
  Service Class          : --
  Color                  : --
  DomainId               : --
  Domain Name            : --
# Run the display tunnel-info tunnel-id command on PEs according to the tunnel ID in the preceding command output. You can view details of the specified tunnel ID.
[PE1] display tunnel-info tunnel-id 2
Tunnel ID:                    0x2
Tunnel Token:                 2
Type:                         gre
Destination:                  10.10.2.1
Out Slot:                     0
Instance ID:                  0
Interface:                    Tunnel0/0/1
# CE1 and CE2 can ping each other successfully.
# The display on CE1 is used as an example.
[CE1] ping 10.1.1.2
  PING 10.1.1.2: 56  data bytes, press CTRL_C to break
    Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=31 ms
    Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=10 ms
    Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=5 ms
    Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=2 ms
    Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=28 ms
  --- 10.1.1.2 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 2/15/31 ms 
Configuration file of CE1
#
 sysname CE1
#
interface GigabitEthernet1/0/0
 ip address 10.1.1.1 255.255.255.0
#
return
Configuration file of PE1
#
 sysname PE1
#
mpls lsr-id 10.10.1.1
mpls
#
mpls l2vpn
#
mpls ldp
#
mpls ldp remote-peer 10.10.2.1
 remote-ip 10.10.2.1
#
interface GigabitEthernet1/0/0
 mpls l2vc 10.10.2.1 39 tunnel-policy gre1
#
interface GigabitEthernet2/0/0
 ip address 172.1.1.1 255.255.255.0
#
interface LoopBack1
 ip address 10.10.1.1 255.255.255.255
#
interface Tunnel0/0/1
 ip address 10.2.1.1 255.255.255.0
 tunnel-protocol gre
 source 10.10.1.1
 destination 10.10.2.1
# 
ospf 1
 area 0.0.0.0
  network 10.10.1.1 0.0.0.0
  network 172.1.1.0 0.0.0.255
#
tunnel-policy gre1 
 tunnel select-seq gre load-balance-number 1
#
return
Configuration file of P
#
 sysname P
#
interface GigabitEthernet2/0/0
 ip address 172.1.1.2 255.255.255.0
#
interface GigabitEthernet1/0/0
 ip address 172.2.1.2 255.255.255.0
#
ospf 1
 area 0.0.0.0
  network 172.1.1.0 0.0.0.255
  network 172.2.1.0 0.0.0.255
#
return
Configuration file of PE2
#
 sysname PE2
#
mpls lsr-id 10.10.2.1
mpls
#
mpls l2vpn
#
mpls ldp
#
mpls ldp remote-peer 10.10.1.1
 remote-ip 10.10.1.1
#
interface GigabitEthernet1/0/0
 ip address 172.2.1.1 255.255.255.0
#
interface GigabitEthernet2/0/0
 mpls l2vc 10.10.1.1 39 tunnel-policy gre1
#
interface LoopBack1
 ip address 10.10.2.1 255.255.255.255
#
interface Tunnel0/0/1
 ip address 10.2.1.2 255.255.255.0
 tunnel-protocol gre
 source 10.10.2.1
 destination 10.10.1.1
#
ospf 1
 area 0.0.0.0
  network 10.10.2.1 0.0.0.0
  network 172.2.1.0 0.0.0.255
#
tunnel-policy gre1 
 tunnel select-seq gre load-balance-number 1
#
return
Configuration file of CE2
#
 sysname CE2
#
interface GigabitEthernet1/0/0
 ip address 10.1.1.2 255.255.255.0
#
return
