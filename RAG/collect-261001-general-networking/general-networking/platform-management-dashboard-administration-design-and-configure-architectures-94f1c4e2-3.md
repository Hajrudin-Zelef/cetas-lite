---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-94f1c4e2-3
title: "platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "datacenter", "license", "licenses"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2.md
source_anchor: ""
source_lines: [112, 161]
sha256: 867f651fe447b26adb8853c7e9face74d6d34ac775703e7db52d0c7a04068138
---

# platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2

Auto VPN Routing
The VPN Registry stores the relevant information including, local routes participating in VPN for a particular Meraki Auto VPN infrastructure. In the case of a failure, additional VPN device, or hub change the system automatically reconverges without any end user interaction. By updating all VPN routes to all devices in the system Auto VPN acts like a routing protocol and converges the system to maintain stability.
High Availability
|  | Key use case | Cost consideration | Failover time | 
| HW Redundancy | Mitigate a WAN Appliance HW failure using 2 devices on the same broadcast domain | Two devices are required but only a single license | Less than 30 seconds (for hardware failover, not necessarily VPN failover) | 
| DC DC Failover | Mitigate any problem that could prevent a spoke from reaching its primary hub | Two devices and two licenses are required | Between 30 seconds and 5 minutes (SD-WAN allows for faster failover) | 
Hardware Redundancy in VPN Concentrator Mode
WAN Appliance VPN Concentrator warm spare is used to provide high availability for a Meraki Auto VPN head-end appliance. Each concentrator has its own IP address to exchange management traffic with the Meraki Cloud. However, the concentrators also share a virtual IP address that is used for non-management communication.
WAN Appliance VPN Concentrator - Warm Spare Setup
Before deploying WAN Appliances as one-arm VPN concentrators, place them into Passthrough or VPN Concentrator mode on the Addressing and VLANs page. In one-armed VPN concentrator mode, the units in the pair are connected to the network "only" via their respective ‘Internet’ ports. Make sure they are NOT connected directly via their LAN ports. Each WAN Appliance must be within the same IP subnet and able to communicate with each other, as well as with the Meraki dashboard. Only VPN traffic is routed to the WAN Appliance, and both ingress and egress packets are sent through the same interface.
WAN Appliance VPN Concentrator - Virtual IP Assignment
The virtual IP address (VIP) is shared by both the primary and warm spare VPN concentrator. VPN traffic is sent to the VIP rather than the physical IP addresses of the individual concentrators. The virtual IP is configured by navigating to Security & SD-WAN > Monitor >Appliance status when a warm spare is configured. It must be in the same subnet as the IP addresses of both appliances, and it must be unique. It cannot be the same as either the primary or warm spare's IP address.
The two concentrators share health information over the network via the VRRP protocol. Failure detection does not depend on connectivity to the Internet/Meraki dashboard.
WAN Appliance VPN Concentrator - Failure Detection
In the event that the primary unit fails, the warm spare will assume the primary role until the original primary is back online. When the primary VPN concentrator is back online and the spare begins receiving VRRP heartbeats again, the warm spare concentrator will relinquish the active role back to the primary concentrator. The total time for failure detection, failover to the warm spare concentrator, and ability to start processing VPN packets is typically less than 30 seconds.
WAN Appliance Warm Spare Alerting
There are a number of options available in the Meraki dashboard for email alerts to be sent when certain network or device events occur, such as when a warm spare transition occurs. This is a recommended configuration option and allows a network administrator to be informed in the event of a failover.
The event, “A warm spare failover occurs,” sends an email if the primary WAN Appliance of a High Availability pair fails over to the spare, or vice-versa.
This alert and others, can be referenced in our article on Configuring Network Alerts in Dashboard.
If you are having difficulties getting warm spare to function as expected, please refer to our MX Warm Spare - High Availability Pair document.
HW Redundancy in NAT mode
WAN Appliance NAT Mode – Warm Spare
WAN Appliance NAT Mode Warm Spare is used to provide redundancy for internet connectivity and appliance services when a WAN Appliance is being used as a NAT gateway.
WAN Appliance NAT Mode - Warm Spare Setup
In NAT mode, the units in the HA pair are connected to the ISP or ISPs via their respective Internet ports, and the internal networks are connected via the LAN ports.
WAN configuration: Each appliance must have its own IP address to exchange management traffic with the Meraki cloud. If the primary appliance is using a secondary uplink, the secondary uplink should also be in place on the warm spare. A shared virtual IP, while not required, will significantly reduce the impact of a failover on clients whose traffic is passing through the appliance. Virtual IPs can be configured for both uplinks.
LAN configuration: LAN IP addresses are configured based on the Appliance IPs in any configured VLANs. No virtual IPs are required on the LAN.
Additional warm spare configuration details can be found in our article, MX Warm Spare - High Availability Pair.
WAN Appliance NAT Mode - Virtual IP Assignment
Virtual IP addresses (vPs) are shared by both the primary and warm spare appliance. Inbound and outbound traffic uses this address to maintain the same IP address during a failover to reduce disruption. The virtual IPs are configured on the Security & SD-WAN > Monitor > Appliance status page. If two uplinks are configured, a vIP can be configured for each uplink. Each vIP must be in the same subnet as the IP addresses of the appliance uplink it is configured for, and it must be unique. It cannot be the same as either the primary or warm spare's IP address.
WAN Appliance NAT Mode - Failure Detection
There are two failure detection methods for NAT mode warm spare. Failure detection does not depend on connectivity to the Internet / Meraki dashboard.
WAN Failover: WAN monitoring is performed using the same internet connectivity tests that are used for uplink failover. If the primary appliance does not have a valid Internet connection based on these tests, it will stop sending VRRP heartbeats which will result in a failover. When uplink connectivity on the original primary appliance is restored and the warm spare begins receiving VRRP heartbeats again, it will relinquish the active role back to the primary appliance.
LAN Failover: The two appliances share health information over the network via the VRRP protocol. These VRRP heartbeats occur at layer 2 and are performed on all configured VLANs. If no advertisements reach the spare on any VLAN, it will trigger a failover. When the warm spare begins receiving VRRP heartbeats again, it will relinquish the active role back to the primary appliance.
WAN Appliance NAT Mode – DHCP Synchronisation
The WAN Appliances in a NAT mode high availability pair exchange DHCP state information over the LAN. This prevents a DHCP IP address from being handed out to a client after a failover if it has already been assigned to another client prior to the failover.
DC-DC Failover - Hub/Data Center Redundancy (Disaster Recovery)
Meraki's WAN Appliance Datacenter Redundancy (DC-DC Failover) allows for network traffic sent across Auto VPN to failover between multiple geographically distributed datacenters.
DC Failover Architecture
A DC-DC failover architecture is as follows:
- 
    One-armed VPN concentrators or NAT mode concentrators in each DC
- 
    A subnet(s) or static route(s) advertised by two or more concentrators
- 
    Hub & Spoke or VPN Mesh topology
- 
    Split or full tunnel configuration
Operation and Failover
Deploying one or more WAN Appliances to act as VPN concentrators in additional data centers provides greater redundancy for critical network services. In a dual- or multi-datacenter configuration, identical subnets are advertised from each datacenter with a VPN concentrator mode WAN Appliance.
