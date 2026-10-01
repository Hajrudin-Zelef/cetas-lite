---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-7
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory", "throughput"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [513, 569]
sha256: 429ab5f38368618e7998162bc33ef063bd74f564b56ec7ea3a9abe2dfe7e48e5
---

# Example SSH command to connect to the FortiGate firewall

- Configuring DHCP: Provided guidance on configuring DHCP server pools on the LAN interface of FortiGate firewall to automate IP address assignment for local users.
- Basic Firewall Policies: Covered the configuration of firewall policies on FortiGate firewall, including setting conditions, actions, and security profiles to control traffic flow between different network segments.
- Network Address Translation (NAT): Explained the theory and configuration of NAT on FortiGate firewall, including static NAT (1-to-1 NAT) and port forwarding to facilitate communication between internal and external networks.
By understanding and implementing the concepts covered in Module 2, administrators can effectively configure interfaces, routing, DHCP, firewall policies, and NAT on FortiGate firewall to ensure efficient network connectivity and robust security measures.
Table of contents:
- Active - Standby(Theory)
- Active - Standby(Lab)
- Active - Active(Theory)
- Active - Active(Lab)
- Identical FortiGate Models: Both FortiGate units in the HA cluster must be identical models to ensure compatibility and proper synchronization.
- Sufficient Resources: Ensure that both FortiGate units have adequate CPU, memory, and storage resources to handle the expected network traffic and configurations.
- Network Interfaces: Each FortiGate unit should have the same number and type of network interfaces (e.g., Ethernet, fiber) configured identically.
- HA Ports: Both FortiGate units must have dedicated HA ports available for HA heartbeat communication and synchronization. These ports should be connected via a dedicated HA link cable or network segment.
- Power and Cooling: Ensure that the power supply and cooling systems are sufficient to support both FortiGate units and maintain optimal operating conditions.
- Compatible Firmware Versions: Both FortiGate units must run the same firmware version to ensure compatibility and proper functionality.
- HA Licensing: Ensure that both FortiGate units are licensed for HA features and functionalities. Some HA features may require specific licensing.
- Configuration Synchronization: Configure both FortiGate units with identical network configurations, security policies, routing settings, and HA settings.
- Virtual Domains (VDOMs): If using VDOMs, ensure that VDOM settings are synchronized between both units and that VDOM HA settings are properly configured.
- Monitoring and Management: Set up monitoring and management tools to monitor the health and status of the HA cluster, including CPU usage, memory utilization, and interface status.
- Dedicated HA Link: Establish a dedicated network link (HA link) between the HA ports of both FortiGate units for heartbeat communication and synchronization.
- Redundant Network Connectivity: Ensure redundant network connectivity for both FortiGate units to prevent single points of failure and ensure continuous operation.
- Network Topology: Configure network routing and VLAN settings to accommodate HA failover events and ensure seamless traffic redirection in case of unit failure.
By meeting these hardware, software, and network requirements, administrators can set up a robust high availability (HA) configuration in the FortiGate firewall to ensure continuous network operation and minimize downtime.
Active-standby mode in the FortiGate firewall, facilitated by FGCP (FortiGate Cluster Protocol), is a high availability (HA) configuration where two firewall units operate in tandem. One unit serves as the primary (active) unit, actively processing traffic, while the other unit acts as the secondary (standby) unit, ready to take over in case of failure.
- Purpose: FGCP is a proprietary protocol developed by Fortinet for communication and synchronization between firewall units in an HA cluster.
- Heartbeat and Configuration Sync: FGCP uses heartbeat signals over the HA link to monitor the status of each unit and synchronize configuration and session information.
- Failover Control: FGCP manages failover events, ensuring a seamless transition between active and standby units without disruption to network traffic.
- Session Pickup: FGCP enables the standby unit to pick up and continue processing existing sessions from the failed active unit during failover.
- Dedicated Connection: The HA link provides a dedicated communication channel between the primary and secondary firewall units for FGCP communication.
- Heartbeat Signals: Heartbeat signals are exchanged over the HA link to monitor the health and availability of each firewall unit in the cluster.
- Synchronization: Configuration and session synchronization occur over the HA link to ensure that both units have identical configurations and state information.
- Active Unit: The active unit processes network traffic and maintains firewall state information, including active sessions, NAT translations, and security policies.
- Standby Unit: The standby unit remains synchronized with the active unit, mirroring its configuration and firewall state, but does not process traffic.
These concepts remain unchanged in the context of FGCP. Gratuitous ARP messages, MAC and IP swap, and priority configurations play crucial roles in ensuring smooth failover and uninterrupted network connectivity during active-standby mode operation with FGCP.
Sample Topology:
FGT-1 Dashboard:
As shown below diagram FGT-1 HA status is “Standalone”
- Go to System —> HA
FGT-2 Dashboard:
Higher priority devices become the Active/Primary. FGT-1
As per our requirement, FGT-1 will be Active, and FGT-2 will be Passive.
- FGT-1 priority is 128 and FGT-2 priority is 100.
- In the Fortigate firewall once synchronization completes we don’t get access to the Passive device.
- Only we can see the status in the Active device
FGT-1 Dashboard of Both Firewall Synchronization.
FGT-2 Lost Connection Screen.
FGT-1 Dashboard Widget.
After the failure of FGT-1, FGT-2 will take over the role of Primary.
Active-active failover in FortiGate firewall is a high availability (HA) configuration where both firewall units actively process network traffic simultaneously, distributing the load across the cluster. This configuration enhances performance and ensures redundancy by allowing seamless failover between units in case of hardware or software failures.
FGCP active-active HA uses a technique similar to unicast load balancing where the primary unit is associated with the cluster HA virtual MAC addresses and cluster IP addresses. The primary unit is the only cluster unit that receives packets sent to the cluster. The primary unit uses a load-balancing schedule to distribute sessions to all cluster units (including the primary unit). Subordinate unit interfaces retain their actual MAC addresses, and the primary unit communicates with the subordinate units using these MAC addresses. Packets exiting the subordinate units proceed directly to their destination and do not pass through the primary unit.
By default, active-active HA load balancing distributes proxy-based security profile processing to all cluster units. Proxy-based security profile processing is CPU and memory-intensive, so FGCP load balancing may result in higher throughput because resource-intensive processing is distributed among all cluster units.
The following proxy-based security profile processing is load-balanced:
- Virus scanning
- Web filtering
- Email Filtering
- Data Loss Prevention (DLP) of HTTP, FTP, IMAP, IMAPS, POP3, POP3S, SMTP, SMTPS, IM, and NNTP sessions accepted by firewall policies
Other features enabled in firewall policies such as endpoint security, traffic shaping, and authentication have no effect on active-active load balancing.
