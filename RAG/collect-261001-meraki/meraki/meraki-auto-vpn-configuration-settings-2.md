---
id: collect-261001-meraki/meraki/meraki-auto-vpn-configuration-settings-2
title: "meraki-auto-vpn-configuration-settings"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-auto-vpn-configuration-settings.md
source_anchor: ""
source_lines: [149, 208]
sha256: 71d7e66f37d10a4e1254a3628a9b9f09e2174733092ae4741ead0c48f2681f27
---

# meraki-auto-vpn-configuration-settings

Multiple hubs can be selected as default routes. These hubs are prioritized in descending order, with the highest priority at the top.

#### Configuring multiple VPN hubs

To add additional hubs, select the **Add a hub** button below the existing hub. 

Only appliances in Mesh VPN mode can be configured as hubs. The number of Mesh VPN appliances in your dashboard organization determines the maximum number of hubs that can be configured.

The order in which hubs are configured defines the **hub priority**. Hub priority determines which hub is used when more than one VPN hub advertises the same subnet. The highest priority hub that meets the following criteria is used: 

- 
    Advertises the subnet

- 
    Currently reachable via VPN

You can manage hubs as follows:

- 
    **Delete a hub:** Click the grey X next to the hub under the Actions column

- 
    **Reorder hubs:** Drag and drop the grey four-point arrow icon to change priority

### Tunneling

There are two tunneling modes available for MX and Z devices configured as a **Spoke**: 

- 
    **Split tunnel (no default route)** : By default all WAN Appliances in the Auto VPN domain (dashboard organization) will only send traffic to an Auto VPN peer if the traffic is destined for a subnet contained within the Auto VPN domain. This is often referred to as 'split-tunnelling,' meaning that VPN-subnet-bound traffic is sent over VPN, and other traffic is routed normally via the primary WAN Appliance WAN uplink. If an organization wants to route all traffic (including traffic not contained within the Auto VPN domain) through a specific hub site, this is referred to as 'full-tunneling.' Full-tunneling only affects client data and all Meraki management traffic will egress directly via the primary WAN regardless.

To configure full-tunneling in a full mesh topology simply define an Exit hub from the WAN Appliances in the Auto VPN domain.

Send only site-to-site traffic, meaning that if a subnet is at a remote site, the traffic destined for that subnet is sent over the VPN. However, if traffic is destined for a network that is not in the VPN mesh (for example, traffic going to a public web service such as www.google.com), the traffic is not sent over the VPN. Instead, this traffic is routed using another available route, most commonly being sent directly to the Internet from the local MX and Z-series device. Split tunneling allows for the configuration of multiple hubs.

- **Full tunnel (default route)** : To configure full-tunneling in a hub-and-spoke topology, simply associate a ‘Default route’ with one or more hub WAN Appliances.

The configured **Exit hub(s)** advertise a default route over Auto VPN to the spoke MX and Z series device. Traffic destined for subnets that are not reachable through other routes will be sent over VPN to the **Exit hub(s)**. **Exit hubs'** default routes will be prioritized in descending order. 

Full tunneling only affects client data, and Meraki management traffic will continue to use the WAN uplink directly.

### Concentrator priority

The concentrator priority determines how appliances in **Hub (Mesh)** mode will reach subnets that are advertised from more than one Meraki VPN peer. Similarly to hub priorities, the uppermost concentrator in the list that meets the following criteria will be used for such a subnet. 

A) Advertises the subnet

B) Currently reachable via VPN

It is important to note that concentrator priorities are used only by appliances in Mesh mode. An appliance in hub-and-spoke mode will ignore the concentrator priorities and will use its hub priorities instead. 

### NAT traversal

If the MX and Z-series device is behind a firewall or other Network Address Translation (NAT) device, there are two options for establishing the VPN tunnel:

- 
    **Automatic** : In most cases, the MX and Z-series device can automatically establish site-to-site VPN connectivity to remote Meraki VPN peers even through a firewall or NAT device using a technique known as "UDP hole punching". This is the recommended (and default) option.

- 
    **Manual: Port forwarding:** If the**Automatic** option does not work, you can use this option. When**Manual: Port forwarding**  is enabled, Meraki VPN peers contact the MX and Z-series device using the specified public IP address and UDP port number. You will need to configure the upstream firewall to forward all incoming traffic on that UDP port to the IP address of the MX and Z-series device.
