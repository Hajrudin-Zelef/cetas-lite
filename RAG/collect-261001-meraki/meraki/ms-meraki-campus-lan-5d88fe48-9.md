---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-9
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "cost", "distribution"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [175, 233]
sha256: c3348a251c80459ff969811fb4e34f7a4ec074bd93a8b1a5e5d6836f52c33868
---

# ms-meraki-campus-lan-5d88fe48

- Connect the aggregation switches with uplinks and get them online on dashboard so they can download available firmware and configuration files (Refer to the installation guide for your Meraki aggregation switches)
- Configure stacking for your aggregation switches and connect stacking cables to bring the stack online (Please follow stacking best practices)
- Enable OSPF where applicable and choose what interfaces should be advertised (Please refer to routing best practices)
- Connect access switches with uplinks and get them online on dashboard so they can download available firmware and configuration files (Refer to the installation guide of your Meraki access switches)
- Configure stacking for your access switches and connect stacking cables to bring the stack online (Please follow stacking best practices)
- Ensure that your security settings (e.g. Switch ACL) have been completed
- Connect access points with uplinks to your access switches and get them online on dashboard so they can download available firmware and configure files (Refer to the installation guide for your Meraki access point)
- Ensure that your switch QoS settings match incoming DSCP values from your APs
- Check your administration settings and adjust dashboard access as required (e.g. Tag based port access)
- Complete other settings in dashboard as required (e.g. Traffic analytics)
- Revisit your dashboard after 7 days to monitor activity and configure tweaks based on actual traffic profiles (e.g. Traffic Shaping on MR APs and switch QoS) and also monitor security events (e.g. DHCP snooping)
- Remember that Campus LAN design is like any other design process and should run in iterations for continuous enhancements and development
- For any Client VLAN changes, start from where your SVI resides (assuming its within the Campus LAN)
- For any native VLAN changes, start from. the lowest layer (e.g. Access Layer) working your way upwards. This will prevent losing access to downstream devices which might require Factory reset
- For any management VLAN changes, attempt to change your IP address settings to DHCP first allowing the switch to acquire an IP address in the designated VLAN automatically. When back online in Dashboard with the new IP address, change the settings to Static assigning the required IP address
- Any SVI or routing changes should be done in a maintenance window as it will result in a brief outage in traffic forwarding
- Always pay attention to platform specific requirements/restrictions. Please refer to the following sections below for further guidance
Redundancy & Resiliency
For optimum distribution-to-core layer convergence, build redundant triangles, not squares, to take advantage of equal-cost redundant paths for the best deterministic convergence. See the below figure for an illustration:
Redundant Triangles
The multilayer switches are connected redundantly with a triangle of links that have Layer 3 equal costs. Because the links have equal costs, they appear in the routing table (and by default will be used for load balancing). If one of the links or distribution layer devices fails, convergence is extremely fast, because the failure is detected in hardware and there is no need for the routing protocol to recalculate a new path; it just continues to use one of the paths already in its routing table.
Redundant Squares
In contrast, only one path is active by default, and link or device failure requires the routing protocol to recalculate a new route to converge.
General Guidance:
- Consider default gateway redundancy (where applicable) using dual connections to redundant distribution layer switches that use VRRP/HSRP/GLBP such that it provides fast failover from one switch to the other at the distribution layer
- Link Aggregation (Ether-Channel or 802.3ad) between switches And/Or switch stacks which provide higher effective bandwidth while reducing complexity and improving service availability
- Deploy redundant triangles as opposed to redundant squares
- Deploy redundant distribution layer switches (preferably stacked together)
- Deploy redundant point-to-point L3 interconnections in the core
- High availability in the distribution layer should be provided through dual equal-cost paths from the distribution layer to the core and from the access layer to the distribution layer. This results in fast, deterministic convergence in the event of a link or node failure
- Redundant power supplies to enhance the overall service availability
- High availability in the distribution layer is achieved through dual equal-cost paths from the distribution layer to the core and from the access layer to the distribution layer. (This results in fast, deterministic convergence in the event of a link or node failure).
The following Meraki MS platforms support Power Supply resiliency:
- MS250
- MS350
- MS355
- MS390
- MS420
- MS425
Meraki MS390 switches support StackPower in addition to Power resiliency and are in combined power mode by default
Firmware
General Guidance
- It’s always important to consider the topology of your switches as, when you drive closer to the network core and away from the access layer, the risk during a firmware upgrade increases
- For Large Campus LAN, it is recommended to start the upgrade closest to the access layer
- For Core switches, it is recommended to reschedule the upgrade to your desired maintenance window
- Staged Upgrades allows you to upgrade in logical increments (For instance, starting from low-risk locations at the access layer and moving onto the higher risk core)
- Firmware for MS switches is set on the network level and therefore all switches in that network will have the same firmware
- Major releases; A new major firmware is released with the launch of new products, technologies and/or major features. New major firmware may also include additional performance, security and/or stability enhancements
- Minor releases; A new minor firmware version is released to fix any bugs or security vulnerabilities encountered during the lifecycle of a major firmware release
- On average, Meraki deploys a new firmware version once a quarter for each product family
- Please plan for sufficient bandwidth to be available for firmware downloads as they can be large in size
- It is recommended to set the out-of-hours preferred upgrade date and time in your network settings for automatic upgrades (remember to set the network's timezone)
- You can also manually upgrade network firmware from Organization > Firmware Upgrades (Meraki will notify you 2 weeks in advance of the scheduled upgrade and, within this two week time window, you have the ability to reschedule to a day and time of your choice)
- Meraki MS devices use a “safe configuration” mechanism, which allows them to revert to the last good (“safe”) configuration in the event that a configuration change causes the device to go offline or reboot.
- During routine operation, if a device remains functional for a certain amount of time (30 minutes in most circumstances, or 2 hours on the MS after a firmware upgrade), a configuration is deemed safe
- When a device comes online for the first time or immediately after a factory reset, a new safe configuration file is generated since one doesn’t exist previously
- It is recommended to leave the device online for 2 hours for the configuration to be marked safe after the first boot or a factory reset.
Multiple reboots in quick succession during initial bootup may result in a loss of this configuration and failure to come online. In such events, a factory reset will be required to recover
General Recommendation
