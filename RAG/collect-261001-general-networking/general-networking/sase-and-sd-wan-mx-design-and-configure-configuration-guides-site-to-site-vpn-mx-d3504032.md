---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-d3504032
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-d3504032"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-d3504032.md
source_anchor: ""
source_lines: [1, 47]
sha256: b1805446c06c4e954fb26904c4b86f36005642b3de2e0b743f014c995fcb8ae4
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-mx-d3504032

MX and Umbrella SIG IPSec Tunnel
Overview
Umbrella SIG (Secure Internet Gateway) is a cloud-based security solution designed for branch offices. This article outlines how to configure Umbrella SIG between the Umbrella dashboard and a Meraki MX Security and SD-WAN device. This solution provides centralized management for security so network administrators do not have to separately manage security settings for each branch. All internet-bound traffic will be forwarded to Umbrella SIG through an IPSec tunnel for inspection and filtering.
Use Case
If network administrators desire a cloud managed security solution to alleviate the inconvenience of separately managing security settings at each branch location, Umbrella SIG is the perfect solution.
Features
Umbrella SIG offers security features such as:
- DNS Policies
- Firewall Policies
- Web Policies
Please visit Umbrella's documentation for a comprehensive guide to Umbrella SIG features.
Prerequisites
- The Umbrella SIG Essentials package is required. Please visit Umbrella for more package information.
- The Meraki MX requires MX 15.12+ firmware, on which users are able to configure the Non-Meraki VPN Peer with the two following Umbrella requirements:
    
  - Choose IKE version type on each Non-Meraki VPN Peer. When choosing IKEv2, the "Local ID" field will be enabled. The User FQDN info needs to be added into this field.
  - On IPSec policies, choose "Diffie-Hellman group" 14.
Configuration
To establish an IPSec tunnel to Umbrella, configurations must be made on both Umbrella Dashboard and Meraki Dashboard.
Umbrella Dashboard
In the Umbrella dashboard, navigate to Deployments > Network Tunnels > select Add
Name the tunnel and select Device Type > Meraki MX.
Set the Tunnel ID and Passphrase. This will be entered as the Local ID (User FQDN) and preshared secret in the Meraki dashboard.
After setting the Tunnel ID and Passphrase, a confirmation prompt will be displayed, allowing you to copy and paste the Tunnel ID and Passphrase to Local ID (User FQDN) and Preshared secret in the Meraki dashboard.
Meraki Dashboard
Navigate to Security & SD-WAN > Configure > Site-to-site VPN > Select desired subnets to participate in VPN.
In the Security & SD-WAN > Configure > Site-to-site VPN > Non-Meraki VPN peers section, select Add a peer.
For the Non-Meraki VPN peers fields:
- Name: Provide any sample name for the tunnel
- Public IP: You will find this IP address in the article at https://docs.umbrella.com/umbrella-user-guide/docs/cisco-umbrella-data-centers
- Local ID: You will get this string from Umbrella dashboard once you have completed creating a Network Tunnel Identity using PSK.
- Private subnets: This will always be 0.0.0.0/0 You will be redirecting all internet bound traffic into the tunnels.
- IPSec policies: Choose Preset of “Umbrella”. This will populate all of the IPSec tunnel parameters necessary for Umbrella connectivity.
- Preshared secret: You will get this string from Umbrella dashboard once you have completed creating a Network Tunnel Identity using PSK.
- Availability: You will add the Tag here that you had defined earlier for the MX appliance that will be building the tunnels to Umbrella cloud.
Verification of the Umbrella IPSec parameters can be viewed by selecting Umbrella
Lastly, you will have to generate interesting traffic through the tunnel in order for the Umbrella dashboard to reflect active tunnel status. To generate interesting traffic, simply source pings from a VPN-participating VLAN (navigate to Security & SD-WAN > Monitor > Appliance Status > Tools) to a destination IP address that would take the IPSec tunnel route.
Meraki dashboard displaying an active Umbrella SIG IPSec tunnel (Security & SD-WAN > Monitor > VPN Status) should look like the following:
Umbrella Dashboard displaying an active IPSec tunnel to Meraki MX (Deployments > Network Tunnels) should look like the following:
Validation
To validate traffic being sent to over the tunnel to SIG vs traffic not being sent over the tunnel we can connect to a network on a VLAN that is participating in tunnel and one that is not to observe the difference.
For this test we used the below configuration where the Default VLAN1 is not participating in VPN and the SIG VLAN10 is participating. This configuration can be viewed under Security & SD-WAN > Configure > Site-to-site VPN.
Using a Wireless capable MX68CW two SSIDs were created. One on VLAN1 and the other on VLAN10.
When a device connects to the SSID SIG1, it receives an IP on VLAN10.
When the device accesses the Internet, the traffic will have a NAT address from Umbrella.
When a device connects to the SSID DIA, it receives an IP on VLAN1.
When the device accesses the Internet, the traffic will have a NAT address from the MX Internet Interface.
