---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a-2
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a.md
source_anchor: ""
source_lines: [42, 92]
sha256: d755512afd742ebfcc33457b6b46dcecc5d387303f24595cf7994eac0eae279e
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a

Use virtual uplink IPs: When using this option, both MXs will use a shared virtual IP (VIP) when sending traffic to the internet. This option requires an additional public IP per uplink, but allows for more seamless failover. This is because the IP address of the outbound flows on the MX will not change, meaning that during the failover client devices will not need to reestablish active sessions. The VIP for each uplink must be in the same subnet as the IPs of the MXs themselves. Also, the VIP must be different from both MX uplink IPs.
Regardless of which option is selected, both MX devices will need their own uplink IP addresses for dashboard connectivity.
Dashboard configuration should always be performed before the secondary MX is physically connected to the network.
Steps to configure secondary appliance:
- 
    Set up the WAN Static IP configuration on the Local Status Page of the secondary appliance (if required).
- 
    Power off the secondary appliance.
- 
    Cable the LAN and WAN connections as per the recommended topology and power on secondary appliance.
Note: When an MX is added to a network that already contains an MX of the same model from the Organization > Configure > Inventory page, the MX will automatically be added to that network in warm spare mode.
Note: MXs operating in HA cannot have more than 255 VLANs configured.
Active/Active High Availability (HA) Mode
From MX 26.2, when configuring a warm spare Secure Router, the Active/Active High Availability mode becomes available along with traditional Active-Passive mode. Active/Active High Availability features allow for warm spare device functions to be used when the primary device is active.
WAN Sharing
From MX 26.2, warm spare Secure Routers using Active/Active High Availability mode can extend their WAN connection to the primary device for active use when operating in routed mode with VLANs. This will enable deployments with ISP connections split between high availability (HA) devices to have multi-uplink capabilities.
Configuration
- Administrators must define which WAN uplink physically resides on each appliance.
- A VLAN will be utilized for HA interconnect, and the device will automatically select the highest numbered VLAN unless manually configured. Traffic destined to the remote WAN uplink will be sent on this connection.
MX Mode Options for Warm Spare Configuration
MX devices can be configured in a high-availability pair (warm spare) using one of two MX addressing options (Security & SD-WAN > Configure > Addressing & VLANs):
- Passthrough or VPN Concentrator mode
- Routed mode
Note: While the mode is reflected in the dashboard as Passthrough or VPN Concentrator mode, the MX only supports a one-armed concentrator topology for this mode. Additional information regarding this can be found in the Connecting the MXs in a “One-Armed” VPN Concentrator Pair section.
Each mode will result in having two MXs on the same network, with a primary able to failover to a secondary. However, each mode requires a slightly different configuration, both detailed in the Addressing and VLANs section. If you need more information about the MX addressing modes or how to select which one is best for your deployment, refer to our MX Addressing and VLANs article.
VPN Concentrator Warm Spare
Concentrator warm spare is used to provide high availability for a Meraki Auto VPN head-end appliance.
Network Setup
Each concentrator has its own IP address to exchange management traffic with the Meraki cloud controller. However, the concentrators also share a virtual IP address that is used for non-management communication.
Connecting the MXs in a “One-Armed” VPN Concentrator Pair
Before deploying MXs as one-arm VPN concentrators, place them into Passthrough or VPN Concentrator mode on the MX Addressing and VLANs page. In one-armed VPN concentrator mode, the units in the pair are connected to the network only via their respective Internet ports. Make sure they are not connected directly via their LAN ports. They must be within the same IP subnet and able to communicate with each other, as well as with the Cisco Meraki dashboard. Only VPN traffic is routed to the MX, and both ingress and egress packets are sent through the same interface.
Virtual IP
The virtual IP (VIP) is shared by both the primary and warm spare VPN concentrator. VPN traffic is sent to the VIP rather than the physical IP addresses of the individual concentrators. The virtual IP is configured by navigating to Security & SD-WAN > Monitor > Appliance status when a warm spare is configured. It must be in the same subnet as the IP addresses of both appliances, and it must be unique. In particular, it cannot be the same as either the primary or warm spare's IP address.
Failure Detection
The two concentrators share health information over the network via the VRRP protocol.
In the event that the primary unit or connectivity tests for its WAN fail, the warm spare will assume the primary role until the original primary is back online or is passing connectivity tests again. When the primary VPN concentrator is back online and the spare begins receiving VRRP heartbeats again, the warm spare concentrator will relinquish the active role back to the primary concentrator.
The total time for failure detection, failover to the warm spare concentrator, and ability to start processing VPN packets is typically less than 30 seconds.
Routed Warm Spare
Routed warm spare is used to provide redundancy for internet connectivity and appliance services when an MX security appliance is being used as a routed gateway.
WAN Virtual IPs
VIP addresses are shared by both the primary and warm spare appliance. Inbound and outbound traffic use this address to maintain the same IP address during a failover and reduce disruption. The virtual IPs are configured on the Security & SD-WAN > Monitor > Appliance status page, under the Spare section in the upper-left corner of the page. If two uplinks are configured, a VIP can be configured for each uplink. Each VIP must be in the same subnet as the IP addresses of both appliances for the uplink it is configured for, and it must be unique. In particular, it cannot be the same as either the primary or the warm spare's IP address.
LAN IP addresses are configured based on the appliance IPs in any configured VLANs. No virtual IPs are required on the LAN.
Note: Modifying the IP address of a WAN connection to use a virtual IP address will result in a loss of connectivity on both Internet uplinks for up to 2 minutes. Therefore, it is recommended to make changes during a planned maintenance window to minimize disruption.
When using features such as port forwarding and NAT rules, services that direct traffic to the HA pair should be configured with the virtual IP address of the HA pair, not the individual WAN IP addresses of the primary and spare MXs.
Additionally, for DDNS to work with virtual IPs, the IP address of the primary uplink and the virtual IP address need to resolve to the same upstream public IP.
Virtual MAC addresses
When using an MX in HA mode:
- If WAN interfaces are configured to use virtual uplink IPs, the WAN interface will use a virtual MAC address. This virtual MAC address is based on the last three octets from the primary MX; the first three octets of the virtual MAC will always be "cc:03:d9". For WAN2, the last octet of the virtual MAC will increment by 1.
- LAN side will use a virtual MAC address for all configured VLANs instead of the device MAC address. If WAN interfaces are configured to use virtual uplink IPs, this will be the same as the WAN1 virtual MAC.
Note: The virtual uplink MAC address for MX HA pairs starts with 'cc:03:d9'. This differs from MS switch virtual MACs, which start with '88:15:44'. Both of these OUIs are owned by Cisco Meraki.
- For MXs in HA mode that are configured to use WAN uplink IPs instead of virtual IPs, the MXs will use the physical MAC address of the respective WAN interface.
