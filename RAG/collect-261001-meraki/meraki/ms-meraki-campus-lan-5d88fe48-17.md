---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-17
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [656, 725]
sha256: 099a98a6313730589d21ba19f68566b93134e9382123d0c62641299c2d17677f
---

# ms-meraki-campus-lan-5d88fe48

  - To RMA a switch: Select the old switch (the one being replaced) and the new switch (that one replacing the old switch) and click clone switch
  - To replace a member with a new switch: Instead, click on Manage members and select the new switch and add it to the stack (The new switch can then be configured as part of the stack with the desired configuration)
- Physically swap the switches
- Remove the old switch from the stack
- Remove the old switch from the network
After the switch has been added to the network and before it is added to the stack or replaced, it should be brought online individually and updated to the same firmware build as the rest of the stack. Failing to do so can prevent the switch from stacking successfully. The configured firmware build for the network can be verified under Organization > Firmware Upgrades. A flashing white or green LED on the status light on the switch indicates that a firmware upgrade is in progress.
If the network is bound to a template, please follow the instructions here instead.
If you face problems with stacking switches, please check the common alerts here.
Physical Stacking - Cloning a stack member
Cloning a stack member can be useful in one of these occasions:
- An identical switch is being added to the network and configuration needs to be cloned from an existing member in one of your stacks
- A failed switch that is RMA'd and needs to be replaced with a new one (like for like) but needs to be operational before the replacement occurs
General Guidance
- Claim the new/replacement switch in the inventory
- Add the switch to the network containing the stack
- Edit the name of the switch if required (For instance, to resemble the old switch e.g. SW-SFO-#5-02)
- Navigate to Switch > Switches
- Select the new/replacement switch and click on Edit > Clone
- Choose the switch that you want to clone the config from
- Click clone
- Power on the switch that is replacing the old one
- Connect a functional uplink to one of the ports on the switch
- Wait for the switch to come online and update its firmware to the one configured on your network (Refer to Organization > Firmware Upgrades and check the Switch details page)
- Navigate to Switch > Switch stacks
- Select the existing stack
- Navigate to Manage members and add the new switch
- Physically swap the switches
- Remove the old switch from the stack
- Remove the old switch from the network
If the network is bound to a template, please follow the instructions here instead.
If you face problems with stacking switches, please check the common alerts here.
Layer 2 Loop-Free Topology
Introduction
Layer 2 loop-free topology and the possibility of enabling Layer 3 on the access switches is an emerging design blueprint due to the following reasons:
- Better convergence results than designs that rely on STP to resolve convergence events
- A routing protocol can even achieve better convergence results than the time-tested L2/L3 boundary hierarchical design
- Convergence based on the up or down state of a point-to-point physical link is faster than timer-based non-deterministic* convergence
- The default gateway is at the Access switch/stack, and a first-hop redundancy protocol is not needed
- Instead of indirect neighbour or route loss detection using hellos and dead timers, physical link loss indicates that a path is unusable; all traffic is rerouted to the alternative equal-cost path
- Using all links from access to core (no STP blocking) thanks to ECMP
* Non-deterministic means that the path of execution isn't fully determined by the specification of the computation, so the same input can produce different outcomes, while deterministic execution is guaranteed to be the same, given the same input
Please check the following diagrams for better understanding the benefits of layer 2 loop-free topology:
Option 1: Gateway Redundancy Protocol (e.g. VRRP) Model
Per the above diagram, L2 links are deployed between the access and distribution nodes. However, no VLAN exists across multiple access layer switches. Additionally, the distribution-to-distribution link is an L3 routed link. This results in an L2 loop-free topology in which both uplinks from the Access layer are forwarding from an L2 perspective and are available for immediate use in the event of a link or node failure. This architecture is ideal for multiple buildings that are linked via fiber connections
In a less-than-optimal design where VLANs span multiple Building Access layer switches, the Building Distribution switches must be linked by a Layer 2 connection. That extends the layer 2 domain from the access layer to the distribution layer. Also, set your STP root and primary gateway on the same Distribution switch
Option 2: Dynamic Routing Protocol (e.g. OSPF) Model
Per the above diagram, L3 links are deployed between the access and distribution nodes using Transit VLANs and SVIs are hosted on the Access Switches. Nonetheless, Etherchannels are used between Access and Distribution Stacks. This results in an L2 loop-free topology in which both uplinks from the Access layer are forwarding from an L2 perspective and are available for immediate use in the event of a link or node failure. This architecture is ideal for a single building where the distribution switches are stacked together in the same rack/cabinet
As seen with the above two options, you can achieve a layer 2 loop free topology. However, Please note that some additional complexity (uplink IP addressing and subnetting) and loss of flexibility are associated with this design alternative.
Now compare that to a Layer 2 looped topology as shown in the following diagram:
As you can see, some L2 links are blocked because of the loop prevention mechanism that is being used (i.e. STP). You must make sure that the STP root and default gateway (HSRP or VRRP) match. STP/RSTP convergence is required for several convergence events. Depending on the version of STP, convergence could take as long as 90 seconds.
General Guidance
- Localize your VLANs to an access switch/stack where possible (Mapping your broadcast domain to your physical space can be beneficial for more than one reason)
- There are many reasons why STP/RSTP convergence should be avoided for the most deterministic* and highly available network topology
- In general, when you avoid STP/RSTP, convergence can be predictable, bounded, and reliably tuned
- L2 environments fail open, forwarding traffic with unknown destinations on all ports and causing potential broadcast storms
- 
    L3 environments fail closed, dropping routing neighbor relationships, breaking connectivity, and isolating the soft failed devices
- 
    If you are running a routed access layer, it is recommended to set the uplink ports as access and keep STP enabled as a failsafe
- 
    If you're running routed distribution layer, it is recommended to summarize routes to the core (where applicable)
* Non-deterministic means that the path of execution isn't fully determined by the specification of the computation, so the same input can produce different outcomes, while deterministic execution is guaranteed to be the same, given the same input
Layer 3 Features
L3 configuration changes on MS210, MS225, MS250, MS350, MS355, MS410, MS425, MS450 require the flushing and rebuilding of L3 hardware tables. As such, momentary service disruption may occur. We recommend making such changes only during scheduled downtime/maintenance window
OSPF
General Guidance
- All Meraki MS switches support OSPF as a dynamic routing protocol
- All configured interfaces should use broadcast mode for hello message
- The following area types are supported: 
    
