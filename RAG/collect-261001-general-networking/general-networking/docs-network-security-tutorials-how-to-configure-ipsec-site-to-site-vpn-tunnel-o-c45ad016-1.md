---
id: collect-261001-general-networking/general-networking/docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016-1
title: "docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016.md
source_anchor: ""
source_lines: [1, 183]
sha256: 2a37e90c14ed7c63349ae5a3e8e062ce10fd5f2af2bbd992ed36acd6748532a1
---

# docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016

How to Configure IPSec Site-to-Site VPN Tunnel on OPNsense?
IPSec is a collection of communication protocols that provide secure connections over a network. The phrase "IPsec" is an abbreviation where "IP" represents "Internet Protocol" and "sec" represents "secure." Internet Protocol (IP) is the universally accepted protocol that governs the transmission of data over the Internet. IPSec enhances the security of the protocol by including encryption and authentication. It is used in virtual private networks (VPNs).
OPNsense provides VPN connectivity for both branch offices and remote users (Road-Warrior). Setting up a single, secure private network that connects several branch offices to a central location is simply accomplished using the OPNsense web user interface. Certificates may be generated and invalidated for distant users, and a user-friendly export tool simplifies the client configuration process.
Site-to-site VPNs VPNs connect two locations and route traffic between their respective networks through the use of static public IP addresses. This is typically employed to establish a connection between the branch offices and the main office of an organization, allowing branch users to retrieve network resources located in the main office.
This guide will explain the process of configuring an IPsec site-to-site VPN tunnel using an OPNsense firewall. You may easily configure IPsec site-to-site VPN tunnel by following 9 main steps:
- Configuring Firewall Rules on Both Site
- Configuring Phase 1 on Site-A
- Configuring Phase 2 on Site-A
- Enabling IPsec on Site-A
- Configuring Phase 1 on Site-B
- Configuring Phase 2 on Site-B
- Enabling IPsec on Site-B
- Adding Firewall Rule for LAN Access on Both Site
- Viewing IPsec Tunnel Status
Sample IPsec Site-to-Site VPN Topology
In this tutorial, the following IP addresses will be used for the sites that connect to each other via an IPsec VPN tunnel:
| Option | Value | 
|---|---|
| Hostname | SiteA_FW | 
| WAN IP | 11.11.11.1/32 | 
| LAN Net | 10.10.10.0/24 | 
| LAN IP | 10.10.10.1/24 | 
| LAN DHCP Range | 10.10.10.100-10.10.10.200 | 
Table 1. IP settings on Site A OPNsense Firewall
| Option | Value | 
|---|---|
| Hostname | SiteB_FW | 
| WAN IP | 11.11.11.2/32 | 
| LAN Net | 10.10.11.0/24 | 
| LAN IP | 10.10.11.1/24 | 
| LAN DHCP Range | 10.10.11.100-10.10.11.200 | 
Table 2. IP settings on Site B OPNsense Firewall
Figure 1. IPsec site-to-site VPN Topology
1. Configuring Firewall Rules on Both Site
To allow IPsec Tunnel Connections, the following ports should be accessible from the Internet on WAN interfaces for both sites.
- UDP Traffic on Port 4500 (NAT-T)
- UDP Traffic on Port 500 (ISAKMP)
- Protocol ESP
You may easily add firewall rules on OPNsense firewalls located in Site A and Site B by following the next steps:
- Allowing ESP port access from the Internet
- Allowing IPSec NAT-T port access from the Internet
- Allowing ISAKMP port access from the Internet
1. Allowing ESP Protocol access from the Internet
Firewall rule settings required for ESP protocol access are given in the next table:
| Option | Value | 
|---|---|
| Action | Pass | 
| Interface | WAN | 
| Protocol | ESP | 
| Source | any | 
| Source Port | any | 
| Destination | WAN address | 
| Destination Port | any | 
| Category | IPsec Tunnel | 
| Description | Allow ESP for IPsec Tunnel | 
Table 3. Firewall rule settings for ESP protocol access
You may easily add firewall rules to allow ESP protocol access for IPsec connection on OPNsense firewalls located in Site A and Site B by following the next steps:
- 
Navigate to the WAN interface on the Firewall Rules.
- 
Select Pass for the allow rule.
- 
Select ESP as the Protocol.
- 
Select any as the source.
- 
Select any as the source port.
- 
Select any as type.
- 
Select WAN address as the destination. Figure 2. Defining firewall rule for ESP access-1
- 
Set Category to IPsec Tunnel .
- 
Set Description to Allow ESP for IPsec Tunnel .
- 
Enable Log packets that are handled by this rule option.
- 
Click Save. Figure 3. Defining firewall rule for ESP access-2
2. Allowing IPsec NAT-T port access from the Internet
Firewall rule settings required for IPSec NAT-T port access are given in the next table:
| Option | Value | 
|---|---|
| Action | Pass | 
| Interface | WAN | 
| Protocol | UDP | 
| Source | any | 
| Source Port | any | 
| Destination | WAN address | 
| Destination Port Range | IPsec NAT-T | 
| Category | IPsec Tunnel | 
| Description | Allow IPsec NAT-T for IPsec Tunnel | 
Table 4. Firewall rule settings for IPSec NAT-T port access
You may easily add firewall rules to allow IPsec NAT-T port access for IPsec connection on OPNsense firewalls located in Site A and Site B by following the next steps:
- 
Navigate to the WAN interface on the Firewall Rules.
- 
Select Pass for the allow rule.
- 
Select UDP as the Protocol.
- 
Select any as the source.
- 
Select any as the source port.
- 
Select any as type.
- 
Select WAN address as the destination. Figure 4. Defining firewall rule for IPsec NAT-T access-1
- 
Select IPsec NAT-T as the destination port range.
- 
Set Category to IPsec Tunnel .
- 
Set Description to Allow IPsec NAT-T port for IPsec Tunnel .
- 
Enable Log packets that are handled by this rule option.
- 
Click Save. Figure 5. Defining firewall rule for IPsec NAT-T port access-2
3. Allowing ISAKMP port access from the Internet
Firewall rule settings required for ISAKMP port access are given in the next table:
| Option | Value | 
|---|---|
| Action | Pass | 
| Interface | WAN | 
| Protocol | UDP | 
| Source | any | 
| Source Port | any | 
| Destination | WAN address | 
| Destination Port Range | ISAKMP | 
| Category | IPsec Tunnel | 
| Description | Allow ISAKMP for IPsec Tunnel | 
Table 5. Firewall rule settings for ISAKMP port access
You may easily add firewall rules to allow ISAKMP port access for IPsec connection on OPNsense firewalls located in Site A and Site B by following the next steps:
- 
Navigate to the WAN interface on the Firewall Rules.
- 
Select Pass for the allow rule.
- 
Select UDP as the Protocol.
- 
Select any as the source.
- 
Select any as the source port.
- 
Select any as type.
- 
Select WAN address as the destination.
- 
Select ISAKMP as the destination port range.
- 
Set Category to IPsec Tunnel .
- 
Set Description to Allow ISAKMP port for IPsec Tunnel .
- 
Enable Log packets that are handled by this rule option.
- 
Click Save. Figure 6. Defining firewall rule for ISAKMP port access
After added these 3 firewall rules on both OPNsense firewalls located on SiteA and SiteB, click Apply Changes button to activate the new settings.
Figure 7. Applying firewall rules for IPsec Tunnel
2. Configuring Phase 1 on Site-A
General Phase-1 options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Connection method | default | default is "Start on traffic" | 
| Key Exchange version | V2 |  | 
| Internet Protocol | IPv4 |  | 
| Interface | WAN | choose the interface connected to the internet | 
| Remote gateway | 11.11.11.2 | the public IP address of your remote OPNsense | 
| Description | Site B | freely chosen description | 
Table 6. General Information Phase-1 options for Site-A
Authentication Phase-1 options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Authentication method | Mutual PSK | Using a Pre-shared Key | 
| My identifier | My IP address | Simple identification for fixed ip | 
| Peer identifier | Peer IP address | Simple identification for fixed ip | 
| Pre-Shared Key | MyS2SIPSecTunnel | Random key. You should create your own one. | 
Table 7. Authentication Phase-1 options for Site-A
Phase 1 proposal (Algorithms) options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Encryption algorithm | 256-bit AES-GCM with128-bit ICV | For our sample we will use 256-bit AES-GCM with128-bit ICV | 
| Hash algoritm | SHA512 | Use a strong hash like SHA512 | 
