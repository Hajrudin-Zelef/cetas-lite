---
id: collect-261001-general-networking/general-networking/hc-en-us-articles-18965560820247-implementing-network-and-client-isolation-in-un-6539781f
title: "hc-en-us-articles-18965560820247-implementing-network-and-client-isolation-in-un-6539781f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-general-networking/hc-en-us-articles-18965560820247-implementing-network-and-client-isolation-in-un-6539781f.md
source_anchor: ""
source_lines: [1, 32]
sha256: 131b23a2d0e052611d57f0aed46591e92ef861347ecac2c9c08ebe015d1d1934
---

# hc-en-us-articles-18965560820247-implementing-network-and-client-isolation-in-un-6539781f

Implementing Network and Client Isolation in UniFi
Once devices are assigned to VLANs, UniFi provides multiple tools to control and enforce separation across and within gateways, switches, and APs. While these tools share a common goal—isolating traffic—they operate at different layers and serve distinct purposes.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Gateway Segmentation
Zone-Based Firewall (ZBF)
Zone-Based Firewalls are available on UniFi Gateways and Cloud Gateways. They enforce policies by defining traffic rules between different network zones, such as VLANs, WANs, and VPNs. ZBF controls how VLANs communicate across the network, making it the most flexible option for inter-VLAN control, especially when VPN traffic needs segmentation. Learn more about Zone-Based Firewalls here.
Network Isolation
For those looking for a simplified, one-click solution, UniFi offers Network Isolation, which automatically configures the necessary firewall rules to block inter-VLAN traffic. To enable:
- Navigate to Settings > Networks.
- Select the desired network or VLAN.
- Enable Network Isolation.
This is the most common way to restrict traffic between VLANs with minimal setup.
Switch Segmentation: Access Control Lists (ACLs)
ACLs provide another layer of segmentation, operating at the switch level. While ZBF governs traffic between VLANs at the gateway, ACLs control traffic passing through a switch within or between VLANs. Unlike ZBF, which applies security rules at the network routing level, ACLs are more lightweight and directly enforce security policies at the switch level.
ACLs can be used for two key types of isolation:
- L3 Network Isolation: Blocks traffic between different VLANs, preventing inter-network communication without using a firewall.
- Client Device Isolation: Restricts communication between devices even within the same VLAN, ensuring devices remain isolated on the same network.
To configure Client Device Isolation via ACLs:
- Navigate to Settings > Networks.
- Enable Device Isolation (ACL) for the appropriate network/VLAN.
- Select the Network/VLAN to apply isolation to.
Note: Device and Switch isolation settings only appear if you have a network/VLAN routed by a UniFi gateway or L3 switch.
ACLs are ideal for restricting inter-VLAN communication or isolating individual devices without relying on a gateway. Learn more about ACLs here.
AP Segmentation: Client Isolation
Unlike ZBF and ACLs, which regulate VLAN-to-VLAN traffic, Client Isolation blocks communication within a single Access Point—even on the same VLAN—making it ideal for guest networks and IoT security. This feature complements ACLs by ensuring that even within a VLAN, devices remain isolated from one another.
To enable Client Isolation:
- Navigate to Settings > WiFi.
- Select the WiFi associated with your network.
- Enable Client Device Isolation.
Best Practices for Public Guest WiFi
These segmentation tools are critical for securing guest networks, preventing unauthorized device-to-device communication, and ensuring user privacy. For a step-by-step guide to implementing these techniques in public guest WiFi, see our Guest WiFi Best Practices.
