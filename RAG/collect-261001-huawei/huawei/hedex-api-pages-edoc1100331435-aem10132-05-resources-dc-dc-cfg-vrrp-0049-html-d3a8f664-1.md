---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vrrp-0049-html-d3a8f664-1
title: "Assign an IP address to each interface. RouterA is used as an example. The configurations of other routers are similar to the configuration of routerA, and are not mentioned here."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["0000-00-00", "2012-05-22"]
keywords: ["preemption"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-vrrp-0049-html-d3a8f664.md
source_anchor: ""
source_lines: [1, 169]
sha256: 8d46d23963a06619d205ba8e07352e2da5fc926cbc6dbe28704572601815c15c
---

# Assign an IP address to each interface. RouterA is used as an example. The configurations of other routers are similar to the configuration of routerA, and are not mentioned here.

As shown in Figure 1, hosts on a LAN are dual-homed to RouterA and RouterB through the switch. A VRRP group is established on RouterA and RouterB, and RouterA is the master. Generally, RouterA functions as the gateway and user traffic is transmitted along the path Switch -> RouterA -> RouterC -> RouterE.
When the link between RouterC and RouterE is faulty or unstable, the VRRP group can detect the fault and an active/standby switchover is performed rapidly. Then RouterB forwards services, so the impact of the link fault on service transmission is reduced.
# Assign an IP address to each interface. RouterA is used as an example. The configurations of other routers are similar to the configuration of routerA, and are not mentioned here.
<Huawei> system-view
[Huawei] sysname RouterA
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ip address 10.1.1.1 24
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 2/0/0
[RouterA-GigabitEthernet2/0/0] ip address 192.168.1.1 24
[RouterA-GigabitEthernet2/0/0] quit
# Configure Layer 2 transparent transmission on the switch.
<Huawei> system-view
[Huawei] sysname Switch
[Switch] vlan 10
[Switch-vlan10] quit
[Switch] interface gigabitethernet 1/0/0
[Switch-GigabitEthernet1/0/0] port hybrid pvid vlan 10
[Switch-GigabitEthernet1/0/0] port hybrid untagged vlan 10
[Switch-GigabitEthernet1/0/0] quit
[Switch] interface gigabitethernet 2/0/0
[Switch-GigabitEthernet2/0/0] port hybrid pvid vlan 10
[Switch-GigabitEthernet2/0/0] port hybrid untagged vlan 10
[Switch-GigabitEthernet2/0/0] quit
# Configure OSPF between devices. RouterA is used as an example. The configurations of other Routers are similar to the configuration of RouterA, and are not mentioned here.
[RouterA] ospf 1
[RouterA-ospf-1] area 0
[RouterA-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
[RouterA-ospf-1-area-0.0.0.0] network 192.168.1.0 0.0.0.255
[RouterA-ospf-1-area-0.0.0.0] quit
[RouterA-ospf-1] quit
# Configure VRRP group 1 on RouterA, and set the priority of RouterA to 120 and the preemption delay to 20s.
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] vrrp vrid 1 virtual-ip 10.1.1.10
[RouterA-GigabitEthernet1/0/0] vrrp vrid 1 priority 120
[RouterA-GigabitEthernet1/0/0] vrrp vrid 1 preempt-mode timer delay 20
[RouterA-GigabitEthernet1/0/0] quit
# Configure VRRP group 1 on RouterB. RouterB uses default priority 100.
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] vrrp vrid 1 virtual-ip 10.1.1.10
[RouterB-GigabitEthernet1/0/0] quit
[RouterA] nqa test-instance user test
[RouterA-user-test] test-type icmp
[RouterA-user-test] destination-address ipv4 20.1.1.2
[RouterA-user-test] frequency 20
[RouterA-user-test] probe-count 5
[RouterA-user-test] fail-percent 80
[RouterA-user-test] start now
[RouterA-user-test] quit
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] vrrp vrid 1 track nqa user test reduced 40
[RouterA-GigabitEthernet1/0/0] quit
# After the configuration is complete, run the display vrrp command on RouterA and RouterB. You can see that RouterA is the master, RouterB is the backup, and the associated NQA test instance is success.
<RouterA> display vrrp
  GigabitEthernet1/0/0 | Virtual Router 1
    State           : Master
    Virtual IP       : 10.1.1.10
    Master IP        : 10.1.1.1
    PriorityRun      : 120
    PriorityConfig   : 120
    MasterPriority   : 120
    Preempt          : YES   Delay Time : 20 s 
    TimerRun         : 1 s 
    TimerConfig      : 1 s
    Auth Type        : NONE
    Virtual Mac      :  0000-5e00-0101
    Check TTL        : YES
    Config type      : normal-vrrp
    Backup-forward   : disabled
    Track NQA : user  test   Priority reduced : 40
    NQA state : success
    Create time      : 2012-05-22 17:32:56
    Last change time : 2012-05-22 17:33:00
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
    Create time      : 2012-05-22 17:33:00
    Last change time : 2012-05-22 17:33:04
# Run the shutdown command on GE1/0/0 of RouterE to simulate a link fault.
[RouterE] interface gigabitethernet 1/0/0
[RouterE-GigabitEthernet1/0/0] shutdown
[RouterE-GigabitEthernet1/0/0] quit
# Run the display nqa results test-instance user test command on RouterA. The command output shows that the NQA test instance status is failed.
<RouterA> display nqa results test-instance user test
NQA entry(user, test) :testflag is active ,testtype is icmp
  1 .Test 1 result   The test is finished
   Send operation times: 5              Receive response times: 0
   Completion:failed                   RTD OverThresholds number: 0
   Attempts number:1                    Drop operation number:0
   Disconnect operation number:0        Operation timeout number:5
   System busy operation number:0       Connection fail number:0
   Operation sequence errors number:0   RTT Stats errors number:0
   Destination ip address:20.1.1.2
   Min/Max/Average Completion Time: 0/0/0                                       
   Sum/Square-Sum  Completion Time: 0/0                                         
   Last Good Probe Time: 0000-00-00 00:00:00.0                                  
   Lost packet ratio: 100 %                                                     
# Run the display vrrp command on RouterA and RouterB. You can see that RouterA is in Backup state, RouterB becomes the master, and the NQA test instance is failed.
<RouterA> display vrrp
  GigabitEthernet1/0/0 | Virtual Router 1
    State           : Backup
    Virtual IP       : 10.1.1.10
    Master IP        : 10.1.1.1
    PriorityRun      : 80
    PriorityConfig   : 120
    MasterPriority   : 100
    Preempt          : YES   Delay Time : 20 s 
    TimerRun         : 1 s 
    TimerConfig      : 1 s
    Auth Type        : NONE
    Virtual Mac      :  0000-5e00-0101
    Check TTL        : YES
    Config type      : normal-vrrp
    Backup-forward   : disabled
    Track NQA : user  test   Priority reduced : 40
    NQA state : failed
    Create time      : 2012-05-22 17:34:56
    Last change time : 2012-05-22 17:35:00
<RouterB> display vrrp
  GigabitEthernet1/0/0 | Virtual Router 1
    State            : Master
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
    Create time      : 2012-05-22 17:35:00
    Last change time : 2012-05-22 17:35:04
# Run the undo shutdown command on GE1/0/0 of RouterE.
[RouterE] interface gigabitethernet 1/0/0
[RouterE-GigabitEthernet1/0/0] undo shutdown
[RouterE-GigabitEthernet1/0/0] quit
# After 20s, run the display vrrp command on RouterA and RouterB. You can see that RouterA restores to be the master and RouterB the backup, and the associated NQA test instance status is success.
<RouterA> display vrrp
  GigabitEthernet1/0/0 | Virtual Router 1
    State           : Master
    Virtual IP       : 10.1.1.10
    Master IP        : 10.1.1.1
    PriorityRun      : 120
    PriorityConfig   : 120
    MasterPriority   : 120
    Preempt          : YES   Delay Time : 20 s 
    TimerRun         : 1 s 
    TimerConfig      : 1 s
    Auth Type        : NONE
    Virtual Mac      :  0000-5e00-0101
    Check TTL        : YES
    Config type      : normal-vrrp
