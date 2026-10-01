---
id: collect-261001-cisco/cisco/enterprise-en-multicast-mac-address-vs-broadcast-mac-address-vs-unicast-mac-addr-d4680002
title: "enterprise-en-multicast-mac-address-vs-broadcast-mac-address-vs-unicast-mac-addr-d4680002"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-multicast-mac-address-vs-broadcast-mac-address-vs-unicast-mac-addr-d4680002.md
source_anchor: ""
source_lines: [1, 8]
sha256: 886f38c48bb0e2f1860aa25ab4ab17cc0ba93c4ce6938b3432e1759cd9e3b5af
---

# enterprise-en-multicast-mac-address-vs-broadcast-mac-address-vs-unicast-mac-addr-d4680002

Hello everyone,
MAC addresses are divided into three categories, which are broadcast addresses, multicast addresses and unicast addresses. This article will give you an overview of the differences between these three types of MAC addresses.
First of all, let's understand the composition of MAC address.
The MAC address of a network device is globally unique. 48 bits in length, usually expressed in 6-byte hexadecimal, the MAC address contains two parts: the first 24 bits are the Organizational Unique Identifier OUI, which is uniformly assigned to the device manufacturer by IEEE. For example, the first 24 bits of the MAC address of Huawei's network products are 0x00e0fc. The last 24 bits of the serial number are unique values assigned to each product by the manufacturer and are assigned by each manufacturer (the product mentioned here can be a network card or other device that requires a MAC address).
Each host interface on the network is uniquely identified by a MAC address. Unicast means that a single source is sent to a single destination. For example, when Host A wants to communicate with Host B individually, Host A needs to send the MAC address of Host B as the destination MAC address to Host B. As defined by the IANA, the leftmost 24 bits of an IPv4 multicast MAC address are 0x01005E, and the 25th bit is 0.
Multicast is often used when a group of hosts on a network (rather than all hosts) need to receive the same information and the other hosts are not affected. Multicast forwarding can be understood as selective broadcasting, where a host listens for a specific multicast address and receives and processes frames with a destination MAC address of that multicast MAC address. The MAC address with the 8th bit being 1 is the unicast MAC address. The multicast MAC address starting with 01-80-c2 is called BPDU MAC, which is generally used as the destination MAC address of a protocol message to mark a certain protocol message.
Broadcast means that a single source is sent to all hosts, and all hosts that receive this broadcast frame receive and process the frame. A MAC address where every bit is 1 is a broadcast MAC address. It is usually denoted as FF-FF-FF-FF-FF-FF. A broadcast MAC address is a special case of a multicast MAC address.
Unicast MAC addresses, broadcast MAC addresses, and multicast MAC addresses are well differentiated. All bits of the broadcast MAC address are 1, the 8th bit of the multicast MAC address is 1, and the 8th bit of the unicast MAC address is 0.
