---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vrrp-0049-html-d3a8f664-2
title: "Assign an IP address to each interface. RouterA is used as an example. The configurations of other routers are similar to the configuration of routerA, and are not mentioned here."
domain: huawei
role: reference
task: reference
actors: []
dates: ["2012-05-22"]
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vrrp-0049-html-d3a8f664.md
source_anchor: ""
source_lines: [170, 299]
sha256: a7704dc88aa95908c9b88e4e2f16a0bef978c29dbe94c144b2efca3b93fe39a5
---

# Assign an IP address to each interface. RouterA is used as an example. The configurations of other routers are similar to the configuration of routerA, and are not mentioned here.

    Backup-forward   : disabled
    Track NQA : user  test   Priority reduced : 40
    NQA state : success
    Create time      : 2012-05-22 17:36:56
    Last change time : 2012-05-22 17:37:00
<RouterB> display vrrp
  GigabitEthernet1/0/0 | Virtual Router 1
    State            : Backup
    Virtual IP       : 10.1.1.10
    Master IP        : 10.1.1.1
    PriorityRun      : 100
    PriorityConfig   : 100
    MasterPriority   : 120
    Preempt          : YES   Delay Time : 0 s 
    TimerRun         : 1 s 
    TimerConfig      : 1 s
    Auth Type        : NONE
    Virtual Mac      :  0000-5e00-0101
    Check TTL        : YES
    Config type      : normal-vrrp
    Backup-forward   : disabled
    Create time      : 2012-05-22 17:37:00
    Last change time : 2012-05-22 17:37:04
RouterA configuration file
#
 sysname RouterA
#
interface GigabitEthernet1/0/0
 ip address 10.1.1.1 255.255.255.0
 vrrp vrid 1 virtual-ip 10.1.1.10
 vrrp vrid 1 priority 120
 vrrp vrid 1 preempt-mode timer delay 20
 vrrp vrid 1 track nqa user test reduced 40
#
interface GigabitEthernet2/0/0
 ip address 192.168.1.1 255.255.255.0
#
nqa test-instance user test                
 test-type icmp                             
 destination-address ipv4 20.1.1.2               
 frequency 20                              
 fail-percent 80
 probe-count 5                
#           
ospf 1
 area 0.0.0.0
  network 192.168.1.0 0.0.0.255
  network 10.1.1.0 0.0.0.255
#
return
RouterB configuration file
#
 sysname RouterB
#
interface GigabitEthernet1/0/0
 ip address 10.1.1.2 255.255.255.0
 vrrp vrid 1 virtual-ip 10.1.1.10
#
interface GigabitEthernet2/0/0
 ip address 192.168.2.1 255.255.255.0
#
ospf 1
 area 0.0.0.0
  network 192.168.2.0 0.0.0.255
  network 10.1.1.0 0.0.0.255
#
return
RouterC configuration file
#
 sysname RouterC
#
interface GigabitEthernet1/0/0
 ip address 192.168.1.2 255.255.255.0
#
interface GigabitEthernet2/0/0
 ip address 20.1.1.1 255.255.255.0
#
ospf 1
 area 0.0.0.0
  network 192.168.1.0 0.0.0.255
  network 20.1.1.0 0.0.0.255
#
return
RouterD configuration file
#
 sysname RouterD
#
interface GigabitEthernet1/0/0
 ip address 192.168.2.2 255.255.255.0
#
interface GigabitEthernet2/0/0
 ip address 30.1.1.1 255.255.255.0
#
ospf 1
 area 0.0.0.0
  network 192.168.2.0 0.0.0.255
  network 30.1.1.0 0.0.0.255
#
return
RouterE configuration file
#
 sysname RouterE
#
interface GigabitEthernet1/0/0
 ip address 20.1.1.2 255.255.255.0
#
interface GigabitEthernet2/0/0
 ip address 30.1.1.2 255.255.255.0
#
ospf 1
 area 0.0.0.0
  network 20.1.1.0 0.0.0.255
  network 30.1.1.0 0.0.0.255
#
return
Switch configuration file
#
sysname Switch
#
vlan batch 10
#
interface GigabitEthernet1/0/0
 port hybrid pvid vlan 10
 port hybrid untagged vlan 10
#
interface GigabitEthernet2/0/0
 port hybrid pvid vlan 10
 port hybrid untagged vlan 10
#
return
