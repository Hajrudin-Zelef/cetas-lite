---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f.md
source_anchor: ""
source_lines: [1, 76]
sha256: 1a87f9434cc559bf26382138805c21744ae2f5b9f990e1d36be98c975ccad4b3
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f

IPsec VPN Configuration Settings
Overview
This article describes how to configure site-to-site IPsec VPN on Cisco Meraki MX WAN appliances. IPsec VPN is required when Auto VPN cannot be used, such as between WAN appliances in different organizations or when connecting to third-party VPN peers. It covers the configuration of IPsec VPN peers, including required parameters like public IP, subnets, pre-shared keys, and IPsec policies, to establish secure connectivity between networks.
Auto VPN vs IPsec VPN
- 
    Auto VPN is a VPN connection between/among the WAN appliances in different networks of the same Meraki dashboard organization.
- 
    Non-Meraki site-to-site VPN is used when you form a VPN tunnel with a third-party/non-Meraki device or when you establish a VPN connection with a Meraki WAN Appliance in a different dashboard organization.
- 
    Like Non-Meraki Site-to-Site VPN, Auto VPN has encryption, authentication and a key. The traffic is encrypted using an AES cipher. However, all of this is transparent to users and does not need to be (and cannot be) modified.
Prerequisites and limitations
Licensing requirements
- 
    Active Cisco Meraki MX license applied to all participating devices
Hardware requirements
- 
    All MX models support IPsec VPN
Firmware requirements
- 
    IKEv2 requires firmware version 15.12 or greater
- 
    IPv6 peering requires firmware version 18.2 or greater
Limitations
- IPsec VPN peers cannot use MX and Z-Series appliances as an "exit hub" as they do not interpret a default route (0.0.0.0/0) as a valid traffic selector, unless you have configured either:
 - A Passthrough or VPN Concentrator MX advertising a 0.0.0.0/0 route as a "Local Network" on the Site-to-Site VPN page.
- A NAT mode MX with a 0.0.0.0/0 LAN static route enabled for VPN.
- IPsec VPN connections are only established over the active WAN uplink, and cannot be established across multiple WAN uplinks unless Multi-Uplink IPsec VPN is configured. Check our documentation on Multi-Uplink IPsec VPN for more information.
Configuration steps
The recommended number of the dashboard configured IPsec VPN peers is 1500.
Starting with MX19 firmware on vMX platforms, Meraki has begun to deprecate the use of DES and 3DES encryption for Phase 2 (IPsec) of Client, and IPsec VPN connections due to its insecure nature. Subsequent firmware releases will continue to deprecate it on all platforms.
To create Site-to-site VPN tunnels between a MX WAN appliance or a Teleworker Gateway and an IPsec VPN endpoint device, navigate to the Security & SD-WAN > Configure > Site-to-site VPN > IPsec VPN peers.
In both organizations, select the Add a peer link. Fill out this entry as if the other MX were a third party device, where each field should be configured as follows:
- 
    A name for the remote device or VPN tunnel. This field is cosmetic.
- 
    What Internet Key Exchange (IKE) version to use (IKEv1 or IKEv2) * for more information regarding the difference between IKEv1 and IKEv2 , refer to the article.
- 
    Enter the public IP address of the remote device. The public IP address from which the remote MX can be contacted.
- If the peer device is part of a high availability (HA) pair, enter the HA virtual IP.
- For MX WAN appliances, locate the public IP address from which the remote MX can be contacted at Security & SD-WAN > Monitor > Appliance status > Uplink > Configuration > General > Public IP
4. To enable private subnets:
All subnets on the remote peer that will be participating in the VPN, in CIDR notation (for example, 10.0.1.0/24). Can be found on the remote MX in the dashboard under Security & SD-WAN > Configure > Addressing & VLANs.
5. Local ID of this MX. This is an optional configuration and is what the remote peer will receive as the remote ID of this MX. Only available with IKEv2.
- If left blank (default) it is the uplink IP of the MX, not the public IP. Some peers may expect this to match the public IP of the MX, a User FQDN (for example, user@domain.com), or FQDN (for example, www.example.com) based on their configuration.
6. The Remote ID of the remote peer. This is an optional configuration and can be configured to the remote peer User FQDN, FQDN, or IPv4 address as needed.
- Which of these values you use is dependent upon your remote device. Consult with its documentation to learn what values it is capable of specifying as its remote ID, and how to configure them (for example, Determining an ID Method for ISAKMP Peers for ASA firewalls).
- 
    The subnets behind the third-party device that you wish to connect to over the VPN. 0.0.0.0/0 can also be specified to define a default route to this peer.
- 
    If an MX device is configured with a default route (0.0.0.0/0) to an IPsec VPN peer, traffic will not fail over to the WAN if the VPN tunnel goes down. The default route will still point to the VPN without failing to direct internet access. However, when health checks are used on routed tunnels, you can configure the option to enable failover directly to the internet.
- 
    Select the IPsec policy.
- 
    It is recommended to keep the default policy on both sides to avoid mismatches.
- 
    If a custom policy is configured for this tunnel on either peer, it must match exactly on both peers.
For more information, refer to the Custom IPsec policies with Site-to-site VPN.
- 
    Enter the pre-shared secret key (PSK).
- 
    A custom passphrase for encryption purposes.. This must match exactly on both peers.
- 
    Availability settings to determine which WAN appliances in your dashboard organization will connect to the IPsec VPN peer.
- 
    By default, all devices in an organization will establish tunnels with a third-party peer. However network tags can be used to limit these connections to a few networks.
- 
    Only networks with a network tag associated under Organization > Overview can be selected in the availability section.
This process would need to be repeated for each remote/local MX pair as desired. The image below shows an example of an MX to MX VPN connection when the devices are in different Organizations:
When setting up IPsec VPN connections between two MXs in different organizations, make sure to populate the Remote ID field of the IPsec VPN peer with the private IP address of the remote MX if all the following conditions are met:
- The MXs are running firmware version MX 15 or higher.
- They do not use User FQDN.
- They are connected behind an upstream NAT device.
An MX WAN appliance can establish tunnels to both Auto VPN and IPsec VPN peers. An MX that builds tunnels to both Auto VPN and IPsec VPN peers will not route traffic between other Auto VPN peers and the IPsec VPN peers. The MX will send traffic to those VPN peers using the principles discussed above and in the MX Routing Behavior Document. However, an MX that builds tunnels to both Auto VPN and IPsec VPN peers, will only route traffic between the IPsec VPN peers and other Auto VPN peers if BGP routing over IPsec VPN is in use.
Making changes to any config under Security & SD-WAN > Configure > Site-to-site VPN, might result in existing IPsec VPN tunnels bouncing. We recommend that you make changes during a maintenance window to minimize the impact of IPsec VPN tunnels bouncing.
Configuring IPsec VPN peering with FQDN
This feature enables the use of FQDN instead of an IP address while configuring an IPsec VPN peer. Using IP addresses can be tedious because with a dynamic IP address, a customer must manually modify the IPsec VPN settings on the Site-to-Site VPN page when there is an IP address change. With FQDN configuration, the hostname of the remote peer would automatically get resolved each time a connection is initiated.
