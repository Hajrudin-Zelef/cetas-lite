---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-7
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "license"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [79, 105]
sha256: 52dbc551044e4f16ed7fa9c3b92b2d0e367dcd83f67a80d02a4e7ea06cb10e47
---

# ms-meraki-campus-lan-5d88fe48

| Adaptive Policy (aka AdP) | Adaptive policy requires MR-ADV license which is only supported in PDL license model When deploying a combined network, please ensure that all devices (e.g. MR, MS390, MX) support AdP (HW and SW) To enable Radius based tagging (e.g. Cisco ISE), please ensure that the RADIUS attribute value pair (av-pair) uses the Cisco SGT AV-Pair presented in HEX value (e.g. cisco-av-pair:cts:security-group-tag=0fa0-00). This example sends back an SGT of 4000 | AdP is supported with:  In a hybrid architecture (e.g. with Cisco Catalyst as Core), please refer to this configuration guide To enable CMD tagging on an upstream TrustSec capable switch, please to IOS-XE Trunk Port Configuration To sync Adaptive Policy between Dashboard and Cisco ISE, refer to this tool. | Tagging is only supported in NAT and Bridge mode If the upstream switch is Meraki MS and the switch port where the MR is connected is not configured with Peer SGT capable, the MR will disable tagging for that AP. The same thing will happen if the upstream switch (regardless of its make or model) does not send Cisco MetaData (CMD) encapsulated traffic for a set number of frames. This is known as fail-safe (AP will completely disable tagging until tagging is enabled on the connected switch and encapsulation is observed on the incoming frames) The RADIUS assignment of group tags is done per-session and to operate will require the av-pair in every access-accept for the client. Please note that you CAN apply a default tag to the SSID and override it with a RADIUS response. | 
|  | SecureConnect automates the process of securely provisioning Meraki MR Access Points when directly connected to switch-ports on Meraki MS Switches, without the requirement of a per-port configuration on the switch  SecureConnect-capable MR access point connected to an MS switch enabled for SecureConnect should not be configured with LAN IP VLAN number | Supported APs will start off only being able to reach dashboard on the switch management VLAN. The APs will have 3 attempts of 5 seconds each to authenticate. If this authentication fails, the switch's port will fall into a restricted state. This might show itself in a couple of ways:   | The failed authentication can happen if the AP and switch are in different organizations, or if the AP is not claimed in inventory. Supported on (with 27.6) MR20, MR30H, MR33, MR42, MR42E, MR52, MR53, MR53E, MR70, MR74, MR84, MR45, MR55, MR36, MR46, MR46E, MR56, MR76, MR86 If you have a MR44 or MR46 please contact Meraki support to check for supportability For further information on supported MS switch models, please check here | 
Wired LAN
Introduction
A traditional Campus LAN Solution will reflect a hierarchical architecture with the following layers:
- Access Layer
- Distribution Layer
- Core Layer
When designing your Wired Campus LAN, it is recommended to start planning in a bottom-up approach (i.e. start at the Access Layer and go upwards). This will simplify the design process and ensure that you have taken into account the design requirements from an end to end perspective. As always, the design process should be done in iterations revising each stage and refining the design elements until the desired outcome can be achieved.
Here's an explanation of each layer in details and what design aspects should be considered for each:
Access Layer
The access layer is the first tier or edge of the campus. It is the place where end devices (PCs, printers, cameras, and the like) attach to the wired portion of the campus network. It is also the place where devices that extend the network out one more level are attached—IP phones and wireless access points (APs) being the prime two key examples of devices that extend the connectivity out one more layer from the actual campus access switch. The wide variety of possible types of devices that can connect and the various services and dynamic configuration mechanisms that are necessary, make the access layer one of the most feature-rich parts of the campus network.
Distribution Layer
The distribution layer in the campus design has a unique role in that it acts as a services and control boundary between the access and the core. It's important for the distribution layer to provide the aggregation, policy control and isolation demarcation point between the campus distribution building block and the rest of the network. It defines a summarization boundary for network control plane protocols (OSPF, Spanning Tree) and serves as the policy boundary between the devices and data flows within the access-distribution block and the rest of the network. In providing all these functions the distribution layer participates in both the access-distribution block and the core. As a result, the configuration choices for features in the distribution layer are often determined by the requirements of the access layer or the core layer, or by the need to act as an interface to both.
Core Layer
The campus core is in some ways the simplest yet most critical part of the campus. It provides a very limited set of services but yet must operate as a non-stop 7x24x365 service. The key design objectives for the campus core must also permit the occasional, but necessary, hardware and software upgrade/change to be made without disrupting any network applications. The core of the network should not implement any complex policy services, nor should it have any directly attached user/server connections. The core should also have the minimal control plane configuration combined with highly available devices configured with the correct amount of physical redundancy to provide for this non-stop service capability.
The following table compares between the main functions and design aspects of the three campus layers:
|  | Access Layer | Distribution Layer | Core Layer | 
| Main Function  |  |  |  | 
| Design Aspects  |  * Please remember to factor for both the total power budget required per switch AND the power standard(s) required |  |  | 
Collapsed Core Layer
One question that must be answered when developing a campus design is this: Is a distinct core layer required? In those environments where the campus is contained within a single building—or multiple adjacent buildings with the appropriate amount of fiber—it is possible to collapse the core into the two distribution switches.
It is important to consider that in any campus design even those that can physically be built with a collapsed distribution core that the primary purpose of the core is to provide fault isolation and backbone connectivity. Isolating the distribution and core into two separate modules creates a clean delineation for change control between activities affecting end stations (laptops, phones, and printers) and those that affect the data center, WAN or other parts of the network. A core layer also provides for flexibility for adapting the campus design to meet physical cabling and geographical challenges.
To illustrate the differences between having a Core layer and a Collapsed Core layer and how that relates to scalability, please see the following two diagrams:
Topology without Core Layer
Topology with Core Layer
Having a dedicated core layer allows the campus to accommodate this growth without compromising the design of the distribution blocks, the data center, and the rest of the network. This is particularly important as the size of the campus grows either in number of distribution blocks, geographical area or complexity. In a larger, more complex campus, the core provides the capacity and scaling capability for the campus as a whole.
