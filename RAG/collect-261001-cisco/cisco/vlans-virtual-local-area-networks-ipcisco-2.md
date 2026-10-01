---
id: collect-261001-cisco/cisco/vlans-virtual-local-area-networks-ipcisco-2
title: "vlans-virtual-local-area-networks-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/vlans-virtual-local-area-networks-ipcisco.md
source_anchor: ""
source_lines: [178, 370]
sha256: 3ffc11905631e3bfbf86c4d6de8cb897f1cd4e3d83f80b4830d6b075ffa3e2ae
---

# vlans-virtual-local-area-networks-ipcisco

```
Switch C (config)# 
```
**interface FastEthernet0/1**
Switch C (config-if)# **switchport mode trunk**
Switch C (config-if)# **switchport nonegotiate**
Switch C (config-if)# **switchport trunk allowed vlan 10,20**
Switch C (config-if)# **interface FastEthernet0/2**
Switch C (config-if)# **switchport mode trunk**
Switch C (config-if)# **switchport nonegotiate**
Switch C (config-if)# **switchport trunk allowed vlan 10,20**


At last step, we will verify our configuration. To do this, we will use “**show vlan brief**” and “**show interface trunk**” commands for the verification.


```
Switch A# 
```
**show vlan brief**
VLAN Name                             Status    Ports
—- ——————————– ——— ——————————-
1    default                          active    Fa0/5, Fa0/6, Fa0/7, Fa0/8
Fa0/9, Fa0/10, Fa0/11, Fa0/12
Fa0/13, Fa0/14, Fa0/15, Fa0/16
Fa0/17, Fa0/18, Fa0/19, Fa0/20
Fa0/21, Fa0/22, Fa0/23, Fa0/24
Gig0/1, Gig0/2
10   VLAN0010                         active    Fa0/2, Fa0/3
20   VLAN0020                         active    Fa0/4
1002 fddi-default                     active
1003 token-ring-default               active
1004 fddinet-default                  active
1005 trnet-default                    active

```
Switch A# 
```
**show interfaces trunk**
Port        Mode         Encapsulation  Status        Native vlan
Fa0/1       on           802.1q         trunking      1
Port        Vlans allowed on trunk
Fa0/1       10,20
Port        Vlans allowed and active in management domain
Fa0/1       10,20
Port        Vlans in spanning tree forwarding state and not pruned
Fa0/1       10,20
 

 

```
Switch B# 
```
**show vlan brief**
VLAN Name                             Status    Ports
—- ——————————– ——— ——————————-
1    default                          active    Fa0/5, Fa0/6, Fa0/7, Fa0/8
Fa0/9, Fa0/10, Fa0/11, Fa0/12
Fa0/13, Fa0/14, Fa0/15, Fa0/16
Fa0/17, Fa0/18, Fa0/19, Fa0/20
Fa0/21, Fa0/22, Fa0/23, Fa0/24
Gig0/1, Gig0/2
10   VLAN0010                         active    Fa0/3, Fa0/4
20   VLAN0020                         active    Fa0/2
1002 fddi-default                     active
1003 token-ring-default               active
1004 fddinet-default                  active
1005 trnet-default                    active

```
Switch B# 
```
**show interfaces trunk**
Port        Mode         Encapsulation  Status        Native vlan
Fa0/1       on           802.1q         trunking      1
Port        Vlans allowed on trunk
Fa0/1       10,20
Port        Vlans allowed and active in management domain
Fa0/1       10,20
Port        Vlans in spanning tree forwarding state and not pruned
Fa0/1       10,20

```
Switch C# 
```
**show vlan brief** 
VLAN Name                             Status    Ports
—- ——————————– ——— ——————————-
1    default                          active    Fa0/6, Fa0/7, Fa0/8, Fa0/9
Fa0/10, Fa0/11, Fa0/12, Fa0/13
Fa0/14, Fa0/15, Fa0/16, Fa0/17
Fa0/18, Fa0/19, Fa0/20, Fa0/21
Fa0/22, Fa0/23, Fa0/24, Gig0/1
Gig0/2
10   VLAN0010                         active    Fa0/3, Fa0/4
20   VLAN0020                         active    Fa0/5
1002 fddi-default                     active
1003 token-ring-default               active
1004 fddinet-default                  active
1005 trnet-default                    active

```
Switch C# 
```
**show interfaces trunk**
Port        Mode         Encapsulation  Status        Native vlan
Fa0/1       on           802.1q         trunking      1
Fa0/2       on           802.1q         trunking      1
Port        Vlans allowed on trunk
Fa0/1       10,20
Fa0/2       10,20
Port        Vlans allowed and active in management domain
Fa0/1       10,20
Fa0/2       10,20
Port        Vlans in spanning tree forwarding state and not pruned
Fa0/1       10,20
Fa0/2       10,20

Beside we will ping PCs in the same VLAN but connected to different VLANs. For example, we will ping from PC1 to PC5 or PC3 to PC4.


Firstly, let’s ping from **PC 1 to PC 5** and **PC8**. These PCs are in **VLAN 10**.


```
C:\> 
```
**ping 10.10.10.5**
Pinging 10.10.10.5 with 32 bytes of data:
Reply from 10.10.10.5: bytes=32 time<1ms TTL=128
Reply from 10.10.10.5: bytes=32 time=1ms TTL=128
Reply from 10.10.10.5: bytes=32 time<1ms TTL=128
Reply from 10.10.10.5: bytes=32 time<1ms TTL=128
Ping statistics for 10.10.10.5:
Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
Minimum = 0ms, Maximum = 1ms, Average = 0ms

```
C:\> 
```
**ping 10.10.10.8**
Pinging 10.10.10.8 with 32 bytes of data:
Reply from 10.10.10.8: bytes=32 time<1ms TTL=128
Reply from 10.10.10.8: bytes=32 time=1ms TTL=128
Reply from 10.10.10.8: bytes=32 time<1ms TTL=128
Reply from 10.10.10.8: bytes=32 time<1ms TTL=128
Ping statistics for 10.10.10.8:
Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
Minimum = 0ms, Maximum = 1ms, Average = 0ms

Then, let’s ping from **PC3 to PC4** and **PC9**. These PCs are in **VLAN 20**.


```
C:\> 
```
**ping 20.20.20.4**
Pinging 20.20.20.4 with 32 bytes of data:
Reply from 20.20.20.4: bytes=32 time=1ms TTL=128
Reply from 20.20.20.4: bytes=32 time<1ms TTL=128
Reply from 20.20.20.4: bytes=32 time<1ms TTL=128
Reply from 20.20.20.4: bytes=32 time=1ms TTL=128
Ping statistics for 20.20.20.4:
Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
Minimum = 0ms, Maximum = 1ms, Average = 0ms

```
C:\> 
```
**ping 20.20.20.9**
Pinging 20.20.20.9 with 32 bytes of data:
Reply from 20.20.20.9: bytes=32 time<1ms TTL=128
Reply from 20.20.20.9: bytes=32 time<1ms TTL=128
Reply from 20.20.20.9: bytes=32 time=2ms TTL=128
Reply from 20.20.20.9: bytes=32 time=1ms TTL=128
Ping statistics for 20.20.20.9:
Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
Minimum = 0ms, Maximum = 2ms, Average = 0ms

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
