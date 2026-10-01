---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-b59e0f13-2
title: "platform-management-dashboard-administration-design-and-configure-architectures--b59e0f13"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--b59e0f13.md
source_anchor: ""
source_lines: [129, 223]
sha256: 0ad25f1ff0f88e1cf529372cbb7e4a804c4b78a1d2d77d717dea8dde414ee20a
---

# platform-management-dashboard-administration-design-and-configure-architectures--b59e0f13

- 
    IGMP Snooping 
  - 
        Disable IGMP Snooping if there are no layer 2 multicast requirements. IGMP Snooping is a CPU dependent feature, therefore it is recommended to utilize this feature only when required. For example, IPTV.
  - 
        It is recommended to use 239.0.0.0/8 multicast address space for internal applications
  - 
        Always configure an IGMP Querier if IGMP snooping is required and there are no Multicast routing enabled switches/routers in the network. A querier or PIM enabled switch/router is required for every VLAN that carries multicast traffic.
- 
        
High Availability and Redundancy
Switch Stacking
The following steps explain how to prepare a group of switches for physical stacking, how to stack them together, and how to configure the stack in the dashboard:
- 
    Add the switches into a dashboard network. This can be a new dashboard network for these switches, or an existing network with other switches. Do not configure the stack in the dashboard yet.
- 
    Connect each switch with individual uplinks to bring them both online and ensure they can check in with the dashboard.
- 
    Download the latest firmware build using the Firmware Upgrade Manager under Organization > Monitor > Firmware Upgrades. This helps ensure each switch is running the same firmware build.
- 
    With all switches powered off and links disconnected, connect the switches together via stacking cables in a ring topology (as shown in the following image). To create a full ring, start by connecting switch 1/stack port 1 to switch 2/stack port 2, then switch 2/stack port 1 to switch 3/stack port 2 and so forth, with the bottom switch connecting to the top switch to complete the ring.
- 
    Connect one uplink for the entire switch stack.
- 
    Power on all the switches, then wait several minutes for them to download the latest firmware and updates from the dashboard. The switches may reboot during this process. 
  - 
        The power LEDs on the front of each switch will blink during this process.
  - 
        Once the switches are done downloading and installing firmware, their power LEDs will stay solid white or green.
- 
        
- 
    Navigate to Switching > Monitor > Switch stacks.
- 
    Configure the switch stack in the dashboard. If the dashboard has already detected the correct stack under Detected potential stacks, click Provision this stack to automatically configure the stack.
- 
    Otherwise, to configure the stack manually:
- 
    Navigate to Switching > Monitor > Switch stacks.
- 
    Click add one / Add a stack:
- 
    Select the checkboxes of the switches you would like to stack, name the stack, and then click Create.
The configuration is complete and the stack should be up and running.
- Use 2 ports on each of “top” and “bottom” switches of the stack for uplink connectivity and redundancy.
- Configure cross-stack link aggregation for uplink connectivity
Warm Spare for Layer 3 Switches
MS Series switches configured for layer 3 routing can also be configured with a “warm spare” for gateway redundancy. This allows two identical switches to be configured as redundant gateways for a given subnet, thus increasing network reliability for users.
Note that, while warm spare is a method to ensure reliability and high availability, generally, we recommend using switch stacking for layer 3 switches, rather than warm spare, for better redundancy and faster failover.
Warm Spare is built on VRRP to provide clients with a consistent gateway. The switch pair will share a virtual MAC address and IP address for each layer 3 interface. The MAC address will always begin with 00-00-5E-00-01, and the IP address will always be the configured interface IP address on the primary. Clients will always use this virtual IP and MAC address to communicate with their gateway.
- 
    For redundancy, ensure an alternate path exists for the exchange of VRRP messages between the Primary and Spare. A direct connection between the Primary and Spare is recommended
Any changes made to L3 interfaces of MS Switches in Warm Spare may cause VRRP Transitions for a brief period of time. This might result in a temporary suspension in the routing functionality of the switch for a few seconds. We recommend making any changes to L3 interfaces during a change window to minimize the impact of potential downtime.
Quality of Service
- 
    Classification 
  - 
        Identify different traffic classes within the network for prioritization. On a high level, traffic can be classified based on VLAN (user, voip, network control etc)
  - 
        You can further classify traffic within a VLAN by adding a QoS rule based on protocol type, source port and destination port as data, voice, video etc.
  - 
        Typical enterprise traffic classes are listed below:
- 
        
- 
    Marking 
  - 
        Meraki MS supports trusting or remarking of incoming DSCP values. Meraki MS supports marking (remarking/trusting) based on DSCP values only. CoS values carried within Dot1q headers are not acted upon. If the end device does not support automatic tagging with DSCP, configure a QoS rule to manually set the appropriate DSCP value.
- 
        
- CoS markings within a Dot1q header are not preserved by default since MS switches support DSCP markings only.
- 
    Queueing and Scheduling 
  - 
        Assign an appropriate Class-of-Service queue to each DSCP value
- 
        
- An MS network has 6 configurable CoS queues labeled 0-5. Each queue is serviced using FIFO. Without QoS enabled, all traffic is serviced in queue 0 (default class) using a FIFO model. The queues are weighted as follows:
| CoS | Weight | 
| 0 (default class) | 1 | 
| 1 | 2 | 
| 2 | 4 | 
| 3 | 8 | 
| 4 | 16 | 
| 5 | 32 | 
Take, for example, a switched environment where VoIP traffic should be in CoS queue 3, an enterprise application in CoS queue 2, and all of other traffic is unclassified. The percentage of bandwidth allocation can be calculated using the weight of the individual CoS queue as the numerator and the sum of all configured CoS queues as the denominator (in this example 8+4+1=13):
- 
    VoIP would be guaranteed 8/13 or ~62% percent of the bandwidth. The switch would forward 8 frames from the CoS queue 3 and move to CoS queue 2.
- 
    The enterprise application would be guaranteed 4/13 or ~30% bandwidth.The switch would forward 4 frames from the CoS queue 2 and move to the default queue.
- 
    All other traffic would receive 1/13 or ~8% of the bandwidth. The switch would forward 1 frame from the default queue, then cycle back to CoS queue 3.
Based on the information above, determine the appropriate CoS queue for each class of traffic in your network. Remember, QoS kicks in only when there is congestion so planning ahead for capacity is always a best practice.
Cabling Best Practices for Multi-Gigabit operations
While Category-5e cables can support multigigabit data rates upto 2.5/5 Gbps, external factors such as noise, alien crosstalk coupled with longer cable/cable bundle lengths can impede reliable link operation. Noise can originate from cable bundling, RFI, cable movement, lightning, power surges and other transient event. It is recommended to use Category-6a cabling for reliable multigigabit operations as it mitigates alien crosstalk by design.
