---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa-1
title: "switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "training"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa.md
source_anchor: ""
source_lines: [1, 75]
sha256: 2469c3fe16ee0450c637fda9b87e3aae09c10921aeb88a8eef19c386ca6fe642
---

# switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa

Switch Stacks
Cloud managed switching supports several types of switch stacking on select switch models. Switch stacking allows several switches to be managed as a single, larger switch which can forward traffic over dedicated stack links rather than front-side network links. In some cases, power redundancy options are available for stacks to survive power supply failures as well.
Learn more with this free online training course on the Meraki Learning Hub:
Switch stack management IP addressing
For MS model switches, such as the MS150 series, each switch in a switch stack requires it's own management IP address. This means that a stack of 8 MS switches require 8 IP management addresses. The only exception to this requirement is the MS390 (see below).
For Catalyst model switches, such as C9300, C9200 series switches, and MS390, the switch stack uses one single management IP address. This means that a stack of 8 Catalyst switches only requires a single management IP address.
Determining How to Stack
Stacking to Fit Your Network
Meraki switches have multiple options to best fit your network deployment. This article discusses the switch stacking features that can be leveraged to best suit your deployment, specifically: Physical Stacking, and Flexible Stacking.
Understanding Physical Stacking
Physical Stacking helps provide easy management and physical redundancy. Utilizing two physical stacking ports on the back of each switch, a stack can provide for gateway redundancy at Layer 3 and dual-homing redundancy at Layer 2. Only a single uplink is required to provide connectivity to the stack once all stacking cables are installed.
For step-by-step instructions, refer to the section of this article titled Configuring a Physical Switch Stack.
When a new switch stack is created, or a new switch is added to an existing stack, the below configurations will be removed from the stand-alone switch(es) and will need to be reconfigured on the stack:
- Link aggregates
- Mirrored switch ports
- Switched virtual interface (SVIs)
- Internet Group Management Protocol (IGMP) snooping (if switch-specific settings are configured)
- Spanning-Tree Protocol (STP) priority (if switch-specific settings are configured)
Features like the ones above run one instance for the entire switch. When a stack is created, you are combining multiple physical switches all running their own instances of the feature to a single logical switch, which is why some features need to be reconfigured.
Understanding StackWise Virtual
StackWise Virtual is a Cisco technology that combines two physical switches into a single logical switch for simplified management, high availability, and scalability. It uses a high-speed StackWise Virtual Link (SVL) to synchronize the switches, providing redundancy, increased bandwidth, and seamless failover for campus core and distribution networks.
StackWise Virtual uses two interfaces
In Cisco StackWise Virtual, the StackWise Virtual Link (SVL) is used to connect two physical switches to operate as a single logical switch. The SVL synchronizes control plane information, forwards data packets, and ensures redundancy between the two switches for seamless failover and high availability.
The Dual-Active Detection (DAD) interface in Cisco StackWise Virtual is used to detect and mitigate a dual-active scenario, where both switches in the virtual stack accidentally operate as active devices due to a StackWise Virtual Link (SVL) failure.
Configuration source: Device
C9500 and C9550 StackWise Virtual (SVL) pairs can be onboarded to the Meraki dashboard.
Configuration source: Cloud
StackWise Virtual (SVL) can be configured through the Meraki dashboard on supported C9500 High Performance models (C9500-24Y4C, C9500-48Y4C, C9500-32QC, and C9500-32C) and C9550. However, existing SVL pairs cannot be onboarded directly to Configuration source: Cloud. Users must first remove the SVL configuration, migrate each switch to Configuration source: Cloud, and then reconfigure SVL through the Meraki dashboard.
Understanding Flexible Stacking
Availability and redundancy are most helpful at the distribution layer of a network. On MS420 and MS425 series switches, any two switch ports can be configured as stack ports. This allows for full redundancy setup for your gateway and minimizes the impact of a failure in the network.
To achieve flexible stacking, go to Switching > Monitor > Switch Ports. Check two switch ports from each eligible switch in the list. Select the Edit button > Stacking port > Enabled. Use the Name field to identify the stack ports as needed.
For step-by-step instructions, refer to the section of this article titled Configuring a Flexible Switch Stack.
Understanding how the Active Stack Member is Elected
- If all stack members are powered up around the same time, the MS with the lowest MAC address will be elected the active switch.
- If the stack members are not powered up around the same time, the MS with the highest uptime will be elected the active switch, regardless of its MAC address.
- The Active switch will remain the Active switch until it has either rebooted or an event occurs that forces another active switch election (such as stack ports being reconnected).
Stacking Availability
Unless specifically noted, only like-models, regardless of switch port density, can be stacked. For example, MS350-48 and MS350-24X can be stacked, but MS250-48 cannot be stacked with an MS350-48. The exception to this rule is for the C9300X which is backward compatible and will stack with C9300 series switches.
| Model | Physical Stacking | Flexible Stacking | 
|---|---|---|
| MS120 | No | No | 
| MS125 | No | No | 
| MS130 | No | No | 
| MS150 | Yes | No | 
| MS210 | Yes (Compatible with MS225) | No | 
| MS220 | No | No | 
| MS225 | Yes (Compatible with MS210) | No | 
| MS250 | Yes | No | 
| MS320 | No | No | 
| MS350 | Yes | No | 
| MS355 | Yes | No | 
| MS390 | Yes | No | 
| C9300-M | Yes, Data + Power | No | 
| C9300X-M | Yes, Data + Power | No | 
| C9300L-M | Yes*, Data Only | No | 
| MS410 | Yes | No | 
| MS420 | No | Yes | 
| MS425 | No | Yes | 
| MS450 | Yes | No | 
Stacking the 9300L requires a stacking kit, it is not included by default. Please order a C9300L-STAK-KIT2-M with every C9300L that you plan on stacking. The stack kit includes two stacking ports and a stack cable
For switches that support Physical/Flexible Stacking:
| Stacking Cable Data Rate | Compatible Series | 
|---|---|
| 40 Gigabit |  | 
| 100 Gigabit |  | 
| 320 Gigabit |  | 
| 480 Gigabit / 1 Terabit |  | 
For full information about stacking cable compatibility, available options, and product IDs, see the Stacking Cables section of the SFP and Stacking Accessories datasheet.
Configuring a Physical Switch Stack
Up to eight switches can be configured in a physical stack to allow for high-speed communication between devices.
Only like-models can be stacked. For example, MS350-48 and MS350-24X can be stacked, but MS250-48 cannot be stacked with a MS350-48. The only exception is that MS210 and MS225 models can be members of the same stack.
Physical stacking is available on MS150, MS210, MS225, MS250, MS350, MS355, MS390, C9300/L/X, MS410, and MS450 switches, which include dedicated stacking ports. This section describes physical stacking.
Flexible stacking is available on MS420 and MS425 switches which do not have dedicated stacking ports; any switch port on these switches can be configured as a stack port. For flexible stacking, check the Configuring a Flexible Switch Stack section of this article.
Physical Switch Stack Configuration Steps
The steps below explain how to prepare a group of switches for physical stacking, how to stack them together, and how to configure the stack in dashboard. Refer to the Meraki Physical Switch Stacking Configuration video.
