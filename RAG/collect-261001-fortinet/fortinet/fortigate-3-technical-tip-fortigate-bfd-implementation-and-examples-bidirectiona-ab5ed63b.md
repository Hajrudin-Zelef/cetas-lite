---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-fortigate-bfd-implementation-and-examples-bidirectiona-ab5ed63b
title: "fortigate-3-technical-tip-fortigate-bfd-implementation-and-examples-bidirectiona-ab5ed63b"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-fortigate-bfd-implementation-and-examples-bidirectiona-ab5ed63b.md
source_anchor: ""
source_lines: [1, 178]
sha256: bbe4add25ab9a337edbb506ce58a7aa1cd4edec136483453c4f8c57673b53080
---

# fortigate-3-technical-tip-fortigate-bfd-implementation-and-examples-bidirectiona-ab5ed63b

**Description**


This article describes the Bidirectional Forwarding Detection implementation and examples.

- Mechanism detecting a one-way device failure.
- Used for faster convergence of routing protocols.
- Independent interface media, routing, or data protocol.
- Draft RFCs with multiple encapsulation types.
- FortiGate uses unicast packets, UDP port 3784 (not routable).
- BFD vs 'Fortinet Dead Gateway Detection' (DGD).


**DGD**:

- Simple mechanism, and no specific protocol is needed on the 'ping server'.
- Based on ping, hence an L3 routable packet that can detect problems on a far-end network.


**BFD**:

- Fast convergence (<1s).
- Devices must be on the same subnet.
- Interoperable support of BFD is required between vendors.
- Designed for OSPF and BGP.


BFD can be configured at multiple levels:

- **Device level:** Enables BFD for all interfaces and protocols on the FortiGate.
- **Interface level:** Overrides the device-level setting, allowing BFD to be enabled or disabled on specific interfaces.
- **Protocol level:** Enables BFD for a specific routing protocol, such as BGP, regardless of interface settings.


The configuration hierarchy allows each lower level to override the BFD setting of the upper level. For example, if BFD is enabled at the device level, it can still be disabled at the interface or protocol level. This allows for flexible control over BFD behavior.


To configure BFD:

1. Enable BFD at the required level (device, VDOM, or interface).
2. Define a BFD neighbor.
3. Enable BFD on the relevant route or routing protocol.


FortiGate BFD/OSPF operation is described in the following scenarios.

If BFD is configured but not OSPF, no BFD packets are sent.


When OSPF is operational, it is possible to see BFD neighbors together with OSPF neighbors.


BFD failure due to remote router (neighbor) failure.


Starting from the previous state (BFD neighbor is up), the BFD failure detection in this case is immediately followed by a withdrawal of the failed OSPF neighbor, triggering route reconvergence. 


BFD packets are seen from the CLI sniffer:


```
FGT # diagnose sniffer packet any "udp port  3784" 6
0.514603 port7 in 192.168.11.53.49161 -> 192.168.11.54.3784: udp 24
0x0000   0000 0000 0001 0009 0f12 b95e 0800 4500        ...........^..E.
0x0010   0034 508b 0000 ff11 d371 c0a8 0b35 c0a8        .4P......q...5..
0x0020   0b36 c009 0ec8 0020 ee8f 20c0 0318 0000        .6..............
0x0030   000a 0000 000d 0000 c350 0000 c350 0000        .........P...P..
0x0040   0000                                           ..
0.530202 port7 out 192.168.11.54.49164 -> 192.168.11.53.3784: udp 24
0x0000   0000 0000 0000 0009 0f12 bcfe 0800 4500        ..............E.
0x0010   0034 c08f 0000 ff11 636d c0a8 0b36 c0a8        .4......cm...6..
0x0020   0b35 c00c 0ec8 0020 ee8c 20c0 0318 0000        .5..............
0x0030   000d 0000 000a 0000 c350 0000 c350 0000        .........P...P..
0x0040   0000                                           ..    
```
                                    

BFD packets are seen with a packet decoder:

**Scope**


FortiGate or VDOM operating in NAT Mode and running OSPF or BGP.**Solution**


**Step 1: BFD must be configured globally and per interface (per neighbor if used for BGP).**Default = 250ms ; threshold = 3.


```
FGT # config system settings
FGT(settings) # set bfd enable
FGT # config system interface
FGT(interface) # edit port7
FGT(port7) # set bfd enable
```

**Step 2: Configuration at the protocol level.**

Configuration example at the OSPF level.


```
config router ospf
    set bfd enable  --> Default = disabled.
        config ospf-interface
            edit dmz
                set bfd enable
```

Configuration example at the BGP level:


```
config router bgp
set as 65250
    config neighbor
        edit 192.168.3.254
            set bfd enable
            set remote-as 65254
```

Verification with CLI commands:

```
TestFGT_5 # get router info bgp neighbors --> Shows details of the neighbors, including: Peer IP address, router ID, remote AS, BGP state, and the negotiated capabilities.
BGP neighbor is 192.168.3.254, remote AS 65254, local AS 65250, external link
TestFGT_5 # get router info bfd neighbor
OurAddr         NeighAddr       LD/RD   State   Int
192.168.3.250   192.168.3.254   4/1     UP      port7
```

BFD troubleshooting commands (CLI):

 For IPv4:

```
get router info bfd neighbor
get router info bfd neighbor detail
get router info bfd requests
diagnose test application bfd 1
```
   

For IPv6:

```
get router info6 bfd neighbor
get router info6 bfd neighbor detail
get router info6 bfd requests
```

BFD debug (CLI):

```
diagnose debug reset
diagnose ip router bfd all enable
diagnose debug app bfdd -1
diagnose debug console timestamp enable
diagnose debug enable
```
   To stop debug:

```
diagnose ip router bfd all disable
diagnose debug disable
diagnose debug reset
```
   

**Notes**:

- State: returns the current state of BFD (UP).
- LD/RD: BFD Local Discriminator / Remote Discriminator used in this BFD session.
- From v7.2.0 onward, it is possible to configure the multi-hop BFD.


**Related documents:**
