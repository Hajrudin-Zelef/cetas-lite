---
id: collect-261001-cisco/cisco/enterprise-en-establishing-a-link-aggregation-control-protocol-lacp-port-channel-683babdc
title: "--- First, clear the existing configuration on the pyhsical interfaces:"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-establishing-a-link-aggregation-control-protocol-lacp-port-channel-683babdc.md
source_anchor: ""
source_lines: [1, 70]
sha256: 16f3f2ae27c11d0a12e46a6dd18008831ab52de7ca5d075c0032494b08a80f6d
---

# --- First, clear the existing configuration on the pyhsical interfaces:

Setting up a Link Aggregation Control Protocol (LACP) link aggregation, which may be referred to as EtherChannel, port-channel, or Eth-trunk, between a Huawei switch and a Cisco switch is a frequent networking task. However, due to diferences in configuration syntax between the two vendors, it can be perplexing.
In this article, I will demonstrate the correct steps for configuring an LACP port-channel, known as Eth-trunk on Huawei devices, between a Cisco Catalyst switch running IOS or IOS-XE and a Huawei switch, specifically the 6700 model in this instance.
To make things even more intriguing, let's specify that we aim to allow only VLANs 100 and 200 to traverse this connection.
! --- First, cleanup the existing configuration on interfaces ten1/1 and 1/2:
default int range ten1/1-2
!
! --- Then, configure the port-chanel, in shutdown:
int range ten1/1-2
shut
channel-group 1 mode active
!
! --- Configure the port-channel itself:
interface Port-channel1
description --- LACP to Huawei Switch
switchport
switchport trunk allowed vlan 100,200
switchport mode trunk
!
! --- Finally, unshut the interfaces:
int range ten1/1-2
no shut
end
# --- First, clear the existing configuration on the pyhsical interfaces:
int range xgi1/0/1 to xgi1/0/2
clear configuration this
shut
#
# --- Then, configure the eth-trunk:
eth-trunk 1 mode active
desc --- LACP to Cisco Switch
# Note: a description, or other configuration, must be written AFTER the eth-trunk configuration line. Otherwise a warning mesage says the interface configuration is not empty so it's not possible to configure an eth-trunk.
#
# --- Configure the port-channel itself: 
interface Eth-Trunk1
desc --- LACP to Cisco Switch
port link-type trunk
port trunk allow-pass vlan 100 200
mode lacp
# Depending of the software version or switch model, the "mode lacp-static" command is necessary.
#
# --- Finally, undo shut the interfaces:
int range xgi1/0/1 to xgi1/0/2
undo shut
#
These are the fundamental configurations.
Further down, you'll discover additional configuration options and troubleshooting commands for both Cisco and Huawei switches.
LACP Priority
You can aggregate up to 8 physical ports into a single logical LACP link. However, you can configure more than 8 ports as hot-standby. To determine which port(s) are active members of the logical link and which port(s) are in hot-standby, you need to set the LACP priority. It's important to note that the lower priorty value takes precedence.
Cisco:
interface ten1/3
lacp port-priority <0-65535>
Huawei:
interface xgi1/0/3
lacp priority <0-65525>
The choice of load-balancing method should be adapted according to the network topology. On Cisco switches, the mode you select applies to all EtherChannels configured on the switch. In contrast, on Huawei switches, it can be configured diferently for each eth-trunk:
port-channel load-balance {src-mac | dst-mac | src-dst-mac | src-ip | dst-ip | src-dst-ip | src-port | dst-port | src-dst-port}
(This list varies depending on the platform)
int eth-trunk 1
load-balance {dst-ip | dst-mac | src-ip | src-mac | src-dst-ip | src-dst-mac}
(This list varies depending on the platform)
If the EtherChannels fail to establish a connection, there are numerous factors to investgate, ranging from physical interfaces to the LACP protocol. Here, I'll mention only the most prevalent issues:
1. Check the physical interfaces:
show int ten1/1dis int xgi1/0/1
2. Check the status of a link aggregation group:
show int po1display eth-trunk 1
3. Check the member interfaces of an EtherChannel interface:
show int po1 etherchannel and show etherchannel detailsdisplay trunkmembership eth-trunk
4. Debugging of LACP packets:
debug lacp packetdebugging trunk lacp-pdu 
Thanks,
