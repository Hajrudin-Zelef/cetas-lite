---
id: collect-261001-cisco/cisco/vlan-configuration-on-cisco-packet-tracer-2026-updated-2
title: "vlan-configuration-on-cisco-packet-tracer-2026-updated"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["pruning", "voice"]
source: docs/RAG/collect-261001-cisco/vlan-configuration-on-cisco-packet-tracer-2026-updated.md
source_anchor: ""
source_lines: [193, 435]
sha256: 66e93d5277d0b57122a22862b898c0a688574d03586342d9968d99a7d8adae67
---

# vlan-configuration-on-cisco-packet-tracer-2026-updated

```
Switch1#
```
**show interfaces fastethernet 0/1 switchport**
Name: Fa0/1
Switchport: Enabled
Administrative Mode: trunk
Operational Mode: trunk
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: dot1q
Negotiation of Trunking: On
Access Mode VLAN: 1 (default)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: 1-3
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none

```
Switch1#
```
**show interfaces fastethernet 0/2 switchport**
Name: Fa0/2
Switchport: Enabled
Administrative Mode: static access
Operational Mode: static access
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: native
Negotiation of Trunking: Off
Access Mode VLAN: 2 (HR)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: All
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none

```
Switch1#
```
**show interfaces fastethernet 0/3 switchport**
Name: Fa0/3
Switchport: Enabled
Administrative Mode: static access
Operational Mode: static access
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: native
Negotiation of Trunking: Off
Access Mode VLAN: 1 (default)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: All
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none

```
Switch1#
```
**show interfaces fastethernet 0/4 switchport**
Name: Fa0/4
Switchport: Enabled
Administrative Mode: static access
Operational Mode: static access
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: native
Negotiation of Trunking: Off
Access Mode VLAN: 1 (default)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: All
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none

```
Switch1#
```
**show interfaces fastethernet 0/5 switchport**
Name: Fa0/5
Switchport: Enabled
Administrative Mode: static access
Operational Mode: static access
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: native
Negotiation of Trunking: Off
Access Mode VLAN: 3 (R&D)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: All
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none

Another important VLAN command is “**show interfaces trunk**” command.  With this command, we can see the details of VLAN Trunk ports.

```
Switch1#
```
**show interfaces trunk**
Port      Mode     Encapsulation     Status      Native vlan
Fa0/1      on             802.1q            trunking            1
Port Vlans allowed on trunk
Fa0/1 1-3
Port Vlans allowed and active in management domain
Fa0/1 1,2,3
Port Vlans in spanning tree forwarding state and not pruned
Fa0/1 1,2,3

You can check the same outputs for switch 2. The outputs for both switch 1 and swicth too are also in the below configuration documents.

```
Switch 2#
```
**show vlan brief**
Switch 2#**show interfaces fastEthernet 0/1 switchport**
Switch 2#**show interfaces fastEthernet 0/2 switchport**
Switch 2#**show interfaces fastEthernet 0/3 switchport**
Switch 2#**show interfaces fastEthernet 0/4 switchport**
Switch 2**#show interfaces fastEthernet 0/5 switchport**
Switch 2#**show interfaces trunk**

To **verify** the communication between **same VLANs** now we will use ping command to check the communication between two PCs in the same VLAN. Here, if the PCs are in the same VLAN, the ping will successful. If they are in different VLANs, ping will not be successful.


```
PC1:\>
```
**ipconfig**
FastEthernet0 Connection:(default port)
Connection-specific DNS Suffix..:
Link-local IPv6 Address.........: FE80::201:C7FF:FE36:68DA
IPv6 Address....................: ::
IPv4 Address....................: 10.0.0.2
Subnet Mask.....................: 255.255.255.0
Default Gateway.................: :: 0.0.0.0

```
PC1:\>
```
**ping 10.0.0.6**
Pinging 10.0.0.6 with 32 bytes of data:
Reply from 10.0.0.6: bytes=32 time<1ms TTL=128
Reply from 10.0.0.6: bytes=32 time<1ms TTL=128
Reply from 10.0.0.6: bytes=32 time<1ms TTL=128
Reply from 10.0.0.6: bytes=32 time<1ms TTL=128
Ping statistics for 10.0.0.6:
Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
Minimum = 0ms, Maximum = 0ms, Average = 0ms

```
PC1:\>
```
**ping 10.0.0.7**
Pinging 10.0.0.7 with 32 bytes of data:
Reply from 10.0.0.7: bytes=32 time<1ms TTL=128
Reply from 10.0.0.7: bytes=32 time<1ms TTL=128
Reply from 10.0.0.7: bytes=32 time<1ms TTL=128
Reply from 10.0.0.7: bytes=32 time<1ms TTL=128
Ping statistics for 10.0.0.7:
Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
Minimum = 0ms, Maximum = 0ms, Average = 0ms

Now, let’s ping from PC3 tom PC7. They are both in VLAN 3 but in different switches.


```
PC3:\>
```
**ping 10.0.0.8**
Pinging 10.0.0.8 with 32 bytes of data:
Reply from 10.0.0.8: bytes=32 time<1ms TTL=128
Reply from 10.0.0.8: bytes=32 time<1ms TTL=128
Reply from 10.0.0.8: bytes=32 time=1ms TTL=128
Reply from 10.0.0.8: bytes=32 time<1ms TTL=128
Ping statistics for 10.0.0.8:
Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
Minimum = 0ms, Maximum = 1ms, Average = 0ms

It is also successful as expected.


We have successfully tested VLAN communication between two switches in the same VLAN. Now, let’s see that, a PC can not ping another PC in another VLAN. To do this, we will ping PC7 from PC1.

