---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a-1
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a.md
source_anchor: ""
source_lines: [1, 41]
sha256: 120dacd9f2b93c975a179758ea9f07a1f2c34e82695d74cee9eb34b1f7151d89
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a

MX Warm Spare - High-Availability Pair
MX High Availability Overview
This page describes how to set up a high-availability (HA) pair using Virtual Router Redundancy Protocol (VRRP) between two MX security appliances. The setup includes either one-arm concentrator mode or routed mode, as well as the expected behavior of the HA pairs. High availability can be used to minimize downtime in the event of a hardware failure.
Only one license is required for an HA pair, so the warm spare unit does not require a separate license. Alerts for warm spare failover can be configured on the Network-wide > Configure > Alerts page.
Note: The spare MX must be the same MX model as the primary. Warm spare functionality is not supported between different MX models (for example, MX85 and MX105) and also MXs with the same model but different territorial designations (e.g., MX67C-WW and MX67C-NA).
The swap button changes the primary and spare roles of the two MX devices and is not meant to test HA failover. For a failover test, you must completely disconnect the uplink to the primary MX.
If the Primary MX goes offline and fails over to the Secondary for an extended period, it is recommended that you perform a swap of the Secondary MX to the Primary. This will prevent any issues with VPN Status and DDNS updates while the original primary is offline if Site-to-Site VPN and DDNS are enabled.
For more information about DDNS, refer to the documentation below:
Dynamic DNS (DDNS)
For more information about VPN Status, refer to the documentation below: 
VPN Status Page
Use Case and Benefits
In most customer deployments, network downtime has a direct impact on the business and should be avoided to prevent service interruptions. Warm spare functionality prevents the network from having a single point of failure, allowing for fast, automatic recovery in the event of device failure. This functionality not only reduces the negative impact on end-user services but also offers significant benefits:
- 
    Reduced Network Downtime: In the event of hardware failure, network downtime will be greatly reduced or eliminated entirely, depending on the architecture being used.
- 
    No Manual Intervention Required: There is no need for manual intervention by the network administration team to facilitate recovery from a hardware failure.
- 
    Zero-Downtime MX Upgrades: When MX appliances are configured to operate in HA (High Availability), the dashboard will automatically take steps to minimize downtime during upgrades. This is achieved through the automated process detailed in the "Appliance Network with Two MXs in an HA Configuration" section of the "Best Practices for Meraki Firmware" document.
Terminology
For purposes of this document, it is important to understand the following terms and their meaning:
Primary: The MX that is configured as the main MX for the network. If both MXs are online, this is the MX that traffic should be flowing through. This is a static designation, meaning that regardless of the current state of the network, the primary will always be the primary.
Spare: The MX that is configured as the secondary MX for the network. If both MXs are online, this is the MX that is the inactive warm spare. This is a static designation, meaning that regardless of the current state of the network, the secondary will always be the secondary.
Active: The MX that is currently acting as the edge firewall/security appliance for the network. This is a dynamic designation.
Passive: The MX that is currently acting as an inactive warm spare with no traffic passing through it. This is a dynamic designation.
Dual active: Dual active describes a scenario in which both the primary and the spare are in the active state. This occurs when both MXs are online and communicating with the cloud, but the spare is not receiving heartbeat packets (see VRRP heartbeats in the next section) from the primary. This can cause several issues with dynamic DNS, VPN, and traffic processing in general and should be avoided at all costs. The Physical Architectures section of this document describes how to deploy an MX warm spare pair in order to minimize the chances of a dual active scenario occurring.
Underlying Concepts and Technologies
VRRP Heartbeats
Failure detection for an MX warm spare pair uses VRRP heartbeat packets. These heartbeat packets are sent from the primary MX to the spare MX on all configured VLANs in order to indicate that the primary is online and functioning properly. As long as the secondary is receiving these heartbeat packets, it functions in the spare state. If the secondary stops receiving these heartbeat packets for 3 seconds, it will assume that the primary is offline and will transition into the active state. When the MX is in routed mode, VRRP heartbeats are not sent over the WAN and there is no guarantee that the WAN interfaces can communicate with each other. See Connection Monitoring below to understand how the WAN interface can also impact how VRRP packets are sent through the LAN on routed mode.
For more in-depth information regarding the VRRP mechanics on the MX, please see the Routed HA Failover Behavior documentation.
Connection Monitoring
Connection monitor is an uplink monitoring engine built into every MX security appliance. The mechanics of the engine are described in the Connection Monitoring for WAN Failover article. When all uplinks of a primary MX are marked as failed by connection monitor, that MX will stop sending VRRP heartbeat packets, which will initiate a warm spare failover. Once there is at least one working uplink, the primary returns to a working state and resumes sending heartbeat packets and the secondary relinquishes the active role back to the primary.
DHCP Synchronization
The DHCP lease table synchronizes regularly between the primary and spare over UDP port 3483. Synchronization prevents a scenario where an IP address is assigned by the primary via DHCP, and then the same IP address is assigned to another client by the spare after a failover.
Dashboard Configuration
To configure warm spare failover for an existing dashboard network, navigate to the Security & SD-WAN > Monitor > Appliance status and select Configure warm spare near the upper-left side of the page, below the device name. In the window that appears, select Enabled. Enter the serial number of the secondary MX and select the desired uplink IP configuration, then select Update to enable warm spare.
Active/Passive High Availability (HA) Mode
Active/Passive High Availability is the new reference for the traditional warm spare operation from MX 26.2+. Onboarding a Warm spare will configure the secondary device as passive with the option to use Uplink IPs or virtual uplink IPs for connectivity.
Adding or replacing an online/offline spare MX in the dashboard network will cause a brief connectivity loss on the primary MX due to the initialization of the HA configuration.
Furthermore, when a warm spare is added to a network, you will lose the ability to use VLAN objects. Any existing L3 rules utilizing VLAN objects will be removed as VLAN objects are not compatible with Warm Spare.
Use MX uplink IPs: When using this option, the current active MX will use its distinct uplink IP or IPs when sending traffic out to the internet. This option does not require additional public IPs for internet-facing MXs, but also results in more disruptive failover. This is because the IP of the outbound flows on the MX will change, which will result in a need for clients to reestablish all live sessions (e.g. web pages, applications, etc).
