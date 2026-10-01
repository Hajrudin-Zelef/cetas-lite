---
id: collect-261001-huawei/huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab7-acl-md-5fb99bc1
title: "kaiyrkhan-huawei-datacom-ensp-blob-head-lab7-acl-md-5fb99bc1"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/kaiyrkhan-huawei-datacom-ensp-blob-head-lab7-acl-md-5fb99bc1.md
source_anchor: ""
source_lines: [1, 132]
sha256: 270aecbcd204b0ac1f26650375646d571ef65e53ad063d426f7f334218769344
---

# kaiyrkhan-huawei-datacom-ensp-blob-head-lab7-acl-md-5fb99bc1

Download Link for eNSP Topology File
- Configure IP Address;
- Configure OSPF;
- Configure Remote Access (Telnet);
- Create an ACL;
- Verify the Configuration.
R1
undo terminal monitor
system-view
sysname R1
int g0/0/0
 ip address 10.1.1.101 30
int loopback0
 ip address 50.1.1.1 32
int loopback1
 ip address 50.2.2.2 32
quit
R2
undo terminal monitor
system-view
sysname R3
int g0/0/1
 ip address 10.1.1.102 30
quit
R1
ospf
area 0
network 50.1.1.1 0.0.0.0
network 50.2.2.2 0.0.0.0
network 10.1.1.100 0.0.0.3
return
R2
ospf
area 0
network 10.1.1.100 0.0.0.3
returnping 50.1.1.1
ping 50.2.2.2
ping 10.1.1.101
R2
telnet server enable
user-interface vty 0 4
 user privilege level 3
 set authentication password cipher Huawei@123
R2
acl 3001
 rule 5 permit tcp source 50.2.2.2 0 destination 10.1.1.102 0 destination-port eq 23
 rule 10 deny tcp source any
 display this
user-interface vty 0 4
acl 3001 inbound
немесе
int g0/0/1
traffic-filter inbound acl 3001
display acl 3001
display cu section acl
Verification:
<R1> telnet -a 10.1.1.101 10.1.1.102
<R1> telnet -a 50.1.1.1 10.1.1.102
<R1> telnet -a 50.2.2.2 10.1.1.102
R1
ping -a 10.1.1.101 10.1.1.102
ping -a 50.1.1.1 10.1.1.102
ping -a 50.2.2.2 10.1.1.102
R2
acl 3002
 rule 5 deny icmp source 50.2.2.2 0 destination 10.1.1.102 0
 rule 10 permit icmp source any
int g0/0/1
traffic-filter inbound acl 3002
display acl 3002
display cu section acl
Verification:
<R1> ping -a 10.1.1.101 10.1.1.102
Reply from 10.1.1.102: bytes=56 Sequence=1 ttl=255 time=20 ms
<R1> ping -a 50.1.1.1 10.1.1.102
Reply from 10.1.1.102: bytes=56 Sequence=1 ttl=255 time=20 ms
<R1> ping -a 50.2.2.2 10.1.1.102
Request time out
SW1
undo terminal monitor
system-view
sysname SW1vlan batch 10 20
display vlan
int g0/0/2
 port link-type access
 port default vlan 10
int g0/0/3
 port link-type access
 port default vlan 10
int g0/0/4
 port link-type access
 port default vlan 20
int g0/0/5
 port link-type access
 port default vlan 20
display vlanint g0/0/1
 port link-type trunk
 port trunk allow-pass vlan 10 20
 display this
RT1
undo terminal monitor
system-view
sysname RT1
int g0/0/0.10
 ip address 172.16.10.1 24
 dot1q termination vid 10
 arp broadcast enable
quit
int g0/0/0.20
 ip address 172.16.20.1 24
 dot1q termination vid 20
 arp broadcast enable
quit
display ip int brief
Verification
<PC1> ping 172.16.10.102
<PC1> ping 172.16.20.101
<PC1> ping 172.16.20.102
RT2
acl 3000
 rule 5 deny ip source 172.16.10.0 0.0.0.255 destination 172.16.20.0 0.0.0.255
 rule 10 permit ip source any destination any
 display this
int g0/0/0
 traffic-filter inbound acl 3000
 display thisdisplay acl 3000
display cu section acl
Verification
<PC1> ping 172.16.10.102
<PC1> ping 172.16.20.101
<PC3> ping 172.16.20.102
<PC3> ping 172.16.10.101
