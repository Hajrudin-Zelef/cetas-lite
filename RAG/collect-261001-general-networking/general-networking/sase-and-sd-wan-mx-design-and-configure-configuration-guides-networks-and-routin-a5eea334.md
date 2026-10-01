---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-a5eea334
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-a5eea334"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-a5eea334.md
source_anchor: ""
source_lines: [1, 81]
sha256: 1dc6bf8af2cdfbd80d2ee72b8a5be9f68ef89f6316b00b1884779e7f8fc549ab
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-a5eea334

Routed HA Failover Behavior
日本語はこちら
Overview
The Cisco Meraki MX offers seamless hardware failover using a warm spare, high availability configuration. This article will detail how an HA pair of MX use Virtual Router Redundancy Protocol (VRRP) to fail over and maintain connectivity for downstream clients.
Please note, this article assumes working knowledge of VRRP and Routed HA.
For more information about VRRP, please reference the RFC.
For more information about Routed HA, please reference our documentation.
VRRP Mechanics for HA
A pair of MX in an HA configuration will use VRRP advertisements to monitor the status of the current active. In a working state, the active MX will send VRRP advertisements out to the LAN every second. If the passive MX does not receive any advertisements for 3 seconds, it assumes that the active MX has failed and will take over as the new active (including sending its own advertisements). This mechanism allows a spare MX to take over in the event of a hardware failure.
In addition to this simple heartbeat mechanic, the active will also report its VRRP priority in the advertisements it sends. For reference, the following priority values are used by each MX, which also depends on whether or not the MX has uplink connectivity:
VRRP Priority Values set on the MX
|  | Primary MX | Spare MX | 
|---|---|---|
| Working Uplinks on primary | 255 | 235 | 
| No Uplink Connection on primary | 75 | 235 | 
| No Uplink Connection on spare | 255 | 55 | 
VRRP Priority Values sent by the MX
|  | Primary MX | Spare MX | 
|---|---|---|
| Working Uplinks on primary | 255 | Nothing | 
| No Uplink Connection on primary | 0 (single advertisement) | 235 | 
| No Uplink Connection on spare | 255 | Nothing | 
If either MX sees a VRRP advertisement with a lower priority than its own, that MX will take over as active.
For example: If the active/primary MX loses all uplink connectivity, it changes its own internal VRRP priority to 75 and sends one-time advertisement with a priority of 0 - a priority of 0 indicates that the sender will no longer be sending advertisements. When the spare MX receives the advertisement with priority 0, it sees that its own priority (235) is greater than the priority within the advertisement, so the spare takes over as the current active and begins sending advertisements with the priority of 235. The primary MX stops sending advertisements until it goes back into a working state.
This mechanism allows a spare MX to take over in the event of an upstream failure on the primary MX.
VRRP with WAN Uplink Sharing
To allow for adequate failover for client traffic between devices with failed uplinks, the VRRP priorities are updated when WAN Uplink Sharing has been enabled.
VRRP Priority Values set on the MX
|  | Primary MX | Spare MX | 
|---|---|---|
| Working Uplinks on primary | 195 | 175 | 
| No Uplink Connection on primary | 195 | 235 | 
| No Uplink Connection on spare | 255 | 175 | 
VRRP Priority Values sent by the MX
|  | Primary MX | Spare MX | 
|---|---|---|
| Working Uplinks on primary | 195 | Nothing | 
| No Uplink Connection on primary | 0 (single advertisement) | 235 | 
| No Uplink Connection on spare | 255 | Nothing | 
Eventlog
| Message | Meaning | 
|---|---|
| VRRP Transition | Event log Event message denoting a change | 
| if_up | VRRP interface state after the event | 
| old_if_up | VRRP interface state before the event | 
| mode | Mode after the event | 
| old_mode | Mode before the event | 
| prio | VRRP Priority after the event | 
| old_prio | VRRP Priority before the event | 
| elector_state | State of MX after the event | 
| last_state_change_reason | Reason for the state change | 
Additional VRRP Notes
Only the current active MX will send VRRP advertisements. In addition to the VRRP priority, there are two key values used by the HA pair:
- VRRP Router ID - A shared router ID that is also used by both of the MX in the warm spare pair.
- VRRP MAC address - The virtual MAC address used on the LAN by both MX.
These two fields are used in conjunction to indicate that a VRRP advertisement is sent by the other MX; they will ignore any VRRP advertisements that do not match these values.
In addition, the VRRP MAC address is shared by both MXs for LAN communication. Clients on the LAN will associate this shared MAC address with the MX's LAN IPs. As such, in the event of failover, LAN clients won't need to update their ARP table with a new MAC address.
MX vs MS Advertisement Timers
MXs use a 1-second timer for VRRP advertisements. This is in contrast to the advertisement timers used by MS switches, where the advertisements are sent every 0.3 seconds. That is why MS switches will failover in 0.9 seconds as opposed to the expected 3-second failover for MX.
Typical Failover Scenario
The following sections walk step-by-step through a common HA failover scenario, wherein the primary MX loses all uplink connectivity and the spare MX takes over.
The following scenario assumes that the primary and spare MXs are connected on the LAN side, and that they are able to exchange VRRP advertisements across all configured VLANs.
Normal State
Starting from a baseline working state, both the primary and spare MX are online with dual uplinks. Everything is normal, so the primary MX is the current active:
In this state, the primary MX sends VRRP advertisements (with a priority of 255) every second:
Primary Uplink Failure
After the primary MX loses all uplink connectivity, it will send a VRRP advertisement with a priority of 0.
The priority value zero (0) has special meaning indicating that the current active has stopped participating in VRRP. This is used to trigger backup routers to immediately transition to active without having to wait for the current active to timeout.
Failover to Spare MX
Once the spare MX receives the 0-priority VRRP advertisement, it will become the new active.
As the new active, the spare MX takes over the LAN by sending its own advertisements with a priority of 235:
Additional Failover Scenarios
The following sections outline some less common failover scenarios:
Both MXs Lose Uplink Connectivity
Assume the end of the scenario above, where the primary MX has no uplink connectivity and the spare MX is the current active.
If the spare MX also loses all uplink connectivity, it will send a VRRP message with a priority of 0:
In this scenario, the primary MX will transition back into the current active role. Without any working uplinks, it will only provide LAN routing:
When the primary MX receives the 0-priority VRRP advertisement, the primary starts sending its own VRRP advertisements with a priority of 75, indicating that it does not have uplink connectivity:
Uplinks and Primary MX Down
In the unlikely scenario that the primary MX's hardware goes down entirely while the spare has no working uplinks, the spare will transition back to the current active role in order to provide LAN routing:
When the spare MX stops seeing any VRRP messages from the primary, the spare MX takes over the LAN by sending its own advertisements with a priority of 55, indicating that it does not have uplink connectivity:
