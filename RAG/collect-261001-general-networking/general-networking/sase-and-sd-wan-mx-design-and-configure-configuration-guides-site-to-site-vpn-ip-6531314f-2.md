---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["aws", "parameters"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f.md
source_anchor: ""
source_lines: [77, 136]
sha256: 60f9bf8cf7e26e3840d3511668d61b5c6beaf5de9fc6aafe295fdfa7ec031459
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-ip-6531314f

“FQDN” differs from “User FQDN.” The MX resolves the FQDN to an IP address of the remote peer, whereas “User FQDN” is used in conjunction with the IP address of the remote peer. “FQDN” identifies the remote peer and is configured in the “Public IP/Hostname” field. “User FQDN” identifies the local peer and is configured in the “Local ID” field.
Supported Products:
- 
    MX running firmware 18.1 or higher
- 
    Requires IKEv2
Configuration
- 
    The FQDN of the IPsec VPN peer can be configured in the Public IP/Hostname field when IKEv2 is the selected IKE version.
The default behavior of the MX is to set remote_id to FQDN if it is not explicitly added in the dashboard for IPsec VPN peer settings. The remote id on one peer needs to match the local id on the other peer for the tunnel to be established.
If the configured FQDN fails to resolve, an event will be reported in Network-wide > Monitor > Event log on the dashboard.
NOTE for IKEv2
Meraki appliances build IPsec tunnels by sending out a request with a single traffic selector payload that contains all the expected local and remote subnets. Certain vendors may not support allowing more than one local and remote selector in each IPsec tunnel. For such cases, use IKEv1 instead.
An MX and Z series device will not try to form a VPN tunnel to an IPsec VPN peer if it does not have any local networks advertised.
IPsec Policies
There are three preset IPsec policies available.
- 
    Default: Uses the Meraki default IPsec settings for connection to an IPsec VPN device
- AWS: Uses default settings for connecting to an Amazon Virtual Private Cloud (VPC)
- Azure: Uses default settings for connecting to a Microsoft Azure instance
If none of these presets are appropriate, the Custom option allows you to manually configure the IPsec policy parameters. These parameters are divided into Phase 1 and Phase 2.
Phase 1
- Encryption: Select between AES-128, AES-192, AES-256, and 3DES encryption
- Authentication: Select MD5, SHA1 or SHA256* authentication
- Diffie-Hellman group: Select between Diffie-Hellman (DH) groups 1, 2, 5, 14, 15, or 21
- Lifetime (seconds): Enter the phase 1 lifetime in seconds
DH group 14 requires firmware version 15.12 or greater. DH groups 15 and 21 require firmware version 19.2 or greater
Phase 2
- Encryption: Select between AES-128, AES-192, AES-256, and 3DES encryption (multiple options can be selected)
- Authentication: Select between MD5, SHA1 and SHA256 authentication (both options can be selected)
- PFS group: Select the Off option to disable Perfect Forward Secrecy (PFS). Select group 1, 2, 5, 14, 15, or 21 to enable PFS using that Diffie-Hellman group.
- Lifetime (seconds): Enter the phase 2 lifetime in seconds
On May 8th 2018, changes were introduced to deprecate DES for encryption. Refer to the Deprecation of DES Encryption Algorithm article for more information.
Ensure the phase 2 lifetimes are equal on both ends of the tunnel whenever possible. While MX's can sometimes honor a shorter phase 2 lifetime if they're acting in response to build a tunnel, they cannot while serving as the initiator of the tunnel.
IPsec VPN failover
For IPsec VPN the MX will fail over the IPsec tunnel from the MX Primary to Secondary internet link when the link is marked as failed. The primary link can be set under Security & SD-WAN > Configure > SD-WAN & Traffic Shaping page.
The IPsec VPN device will need to have the ability to designate a backup (failover) peer IP. By designating the public IP address of the MXs secondary uplink as the back-up VPN IP on the IPsec VPN peer, you can ensure that the VPN tunnel will be re-established in the event of an uplink failure.
If your MX device is operating on firmware version MX19.1.4 or later, you have the option to establish both Primary and Secondary VPN tunnels to enhance redundancy. For more information, refer to Primary and Secondary IPsec VPN tunnels.
Peer availability
By default, a IPsec VPN peer configuration applies to all MX and Z-series in your dashboard organization. Since it is not always desirable for every appliance you control to form tunnels to a particular IPsec VPN peer, the Availability column allows you to control which appliances within your Organization will connect to each peer. This control is based on network tags, which are labels you can apply to your dashboard networks.
When All networks is selected for a peer, all MX and Z-series in the organization will connect to that peer. When a specific network tag or set of tags is selected, only networks that have one or more of the specified tags will connect to that peer.
For information on network tags, refer to Organization Overview article.
VPN firewall rules
You can add firewall rules to control what traffic is allowed to pass through the VPN tunnel. These rules will apply to outbound VPN traffic to/from from all MX-Z series appliances in the Organization that participate in site-to-site VPN, including outbound IPsec VPN peer tunnel and related BGP control traffic. Configuration steps for the rules can be found in the Site-to-site VPN Firewall Rule Behavior article.
VPN Firewall rules will not apply to inbound traffic or to traffic that is not passing through the VPN.
Monitoring Site-to-site VPN
You can monitor the status of the site-to-site VPN tunnels between your Meraki devices by selecting Security & SD-WAN > Monitor > VPN Status. This page provides real-time status for the configured Meraki site-to-site VPN tunnels. It lists the subnet(s) being exported over the VPN, connectivity information between the MX and Z-series and the Meraki VPN registry, NAT Traversal information, and the encryption type being used for all tunnels. For more information regarding this page, refer to the VPN Status Page article.
Troubleshooting
If the VPN tunnel is not coming up upon following the set-up guidelines listed above, you may need to generate interesting traffic first. This can be done by initiating a ping to an IP address in the remote subnet.
For other common issues and troubleshooting steps, please refer to Site-to-site VPN Troubleshooting.
MX and Z-series appliances in Passthrough or VPN Concentrator mode cannot be used to initiate interesting traffic into the tunnel. Interesting traffic would need to be initiated from a client behind the Passthrough appliance.
Additional considerations
Caveats
Since this VPN tunnel is functionally the same as a tunnel to a third-party peer, the same restrictions and caveats apply, including the following notable caveats:
- 
    Limited visibility on the VPN Status Page.
- 
    To bring the tunnel up, you may need to generate interesting traffic which can be done by initiating a ping to an IP address in the remote subnet.
Additional resources
For more information about site-to-site VPN tunnels and troubleshooting:
