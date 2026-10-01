---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5cb3161a
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5cb3161a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5cb3161a.md
source_anchor: ""
source_lines: [1, 35]
sha256: af401200290fa5584a6b40680464d5c77a55d5ce8a301556026de5a438f14295
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5cb3161a

Cisco Meraki MX security appliances support the OSPF routing protocol to advertise remote VPN subnets to neighboring layer 3 devices. This feature is useful in topologies where a large number of VPN subnets makes configuring static routes impractical.

This article outlines the prerequisites and configuration necessary for OSPF on the MX platform.

**Note:** As part of **MX 18.1 firmware** updates we are introducing a new “Routing page”. More details found on IPv6 Support on MX Security & SD-WAN Platforms - Routing.

MX devices in Routed mode only support OSPF on firmware versions 13.4+, when using the "Single LAN" LAN setting. OSPF is otherwise supported when the MX is in Passthrough mode on any available firmware version. This can be set under **Security & SD-WAN > Configure > Addressing & VLANs**. 

The MX will *only* advertise Meraki Auto VPN routes. *The MX does not learn routes advertised by any OSPF neighbors.*

To achieve symmetrical routing between spoke MXs participating in AutoVPN and OSPF peers, *the hub MX will need to have static route(s) configured for any subnets local to the OSPF peer, which will have to be advertised in the AutoVPN mesh* (by switching the VPN mode for those particular static route(s) to Enabled).

## Configuration

To configure OSPF on the MX, navigate to **Security & SD-WAN > Configure > Routing**.

Enabling **OSPF** will provide additional configuration options: 

- **Router:** The OSPF Router ID that the MX will use to identify itself to neighbors.
- **Area ID:** The OSPF Area ID that the MX will use when sending route advertisements.
- **Cost:** (Defaults to 1) The route cost attached to all OSPF routes advertised from the MX.
- **Hello timer:** (Defaults to 10) How frequently the MX will send OSPF Hello packets in seconds. This should be the same across all devices in your OSPF topology.
- **Dead timer:** (Defaults to 40) How long the MX will wait (in seconds) to see Hello packets from a particular OSPF neighbor before considering that neighbor inactive.

- **Authentication:** (Defaults to disabled) If this is enabled, MD5 hashing will be used to authenticate potential OSPF neighbors. This ensures that no unauthorized devices are injecting OSPF routes into the network.
  - **MD5** authentication key: The MD5 ID number and password. Both of these values must match between any devices that you wish to form an OSPF adjacency.

To confirm that the MX is sending OSPF updates, packet captures can be taken.

- MX in Routed mode - Captures must be taken on the LAN interface
- MX in Passthrough or VPN Concentrator mode - Captures must be taken on the WAN interface

This will show the MX sending updates to other OSPF enabled devices. An in-depth reference of an OSPF adjacency being formed can be found 'MS OSPF Overview'.

*MX appliances will now properly validate that DBD packets conform to the appropriate MTU size. If the MX’s OSPF peer has an improper MTU configured, it may cause the OSPF adjacency to fail to properly form. The updated behavior properly conforms to RFC. Please ensure these settings are properly configured on any MX’s OSPF peers to avoid disruption after upgrading to MX 18.1.X.*
