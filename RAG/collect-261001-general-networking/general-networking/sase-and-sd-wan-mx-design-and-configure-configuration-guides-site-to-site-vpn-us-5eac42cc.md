---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5eac42cc
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5eac42cc"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5eac42cc.md
source_anchor: ""
source_lines: [1, 91]
sha256: 1cbcbdbdc2631d23a965c49c60cac2e88f6d789e184a1cadda01db69fc06a1ff
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-us-5eac42cc

Click 日本語 for Japanese

## Overview

In a distributed deployment of locations connected via a site-to-site VPN, a network administrator may need to have address translation performed on traffic traversing the site-to-site VPN. A 1:1 subnet translation can be used in cases where multiple locations have the same subnet present, but both need to participate in the site-to-site VPN. Alternatively, administrators may need to conserve IP space for large deployments. For this, 1:M NAT can be used to translate entire subnets into a single IP address that is exported across the site-to-site VPN.

## VPN Subnet Translation

VPN subnet translation allows for a subnet that is allowed in the site-to-site VPN to be translated to a different, equally sized subnet. This option is ideal for deployments where the same subnet is used in multiple locations and each of those subnets need to have access to the site-to-site VPN.

### Configuration

To configure VPN subnet translation:

1. Navigate to **Security & SD-WAN >****Configure > Site-to-site VPN** .
2. Set **VPN subnet translation** to**Enabled** . This will cause a new**VPN subnet** column to appear for the local networks.
3. For the local subnet that must be translated, set VPN participation to **VPN on with translation** .
4. In the **VPN subnet** column enter a subnet of the same size as the**Local subnet** .
5. Select **Save changes** .

### Operation

Please consider the following example:

- The 192.168.128.0/24 subnet exists in two locations
- The devices and users in this subnet at both locations need to access resources across a site-to-site VPN connection
- To avoid address and routing conflicts across the site-to-site VPN, 192.168.128.0/24 has been configured to be translated to 10.15.30.0/24
- A host on the corporate VLAN with an IP address of 192.168.128.44 is communicating with a web server across the site-to-site VPN with an address of 172.16.30.8
- The web server is also connected locally to another MX security appliance. This MX is a part of the site-to-site VPN.


When VPN subnet translation is configured, the MX will check the source IP address against a address translation table. When 192.168.128.44 attempts to send traffic to the web server across the VPN, the source IP address is evaluated to be contained within the local subnet of 192.168.128.0/24, which requires a translation to be performed. The MX will then map the client's IP to the equivalent IP in the translated subnet. When the example client's traffic egresses the site-to-site VPN, it will have an IP address of 10.15.30.44.


If VPN subnet translation is configured, the translated subnet will automatically be advertised to all remote site-to-site VPN participants. In this example, in order for the web server at 172.16.30.8 to communicate with the example client, traffic must be sent to 10.15.30.44 (the equivalent IP offset within the translated subnet). When the web server's traffic is sent to 10.15.30.44 and received by it's local MX, it will be routed to the appropriate remote MX and the destination IP address will be translated back to 192.168.128.44 before it egresses the MX's LAN.

When configuring VPN subnet translation for a local subnet that exists in multiple locations, the duplicated subnet must be translated at each network that is configured to allow VPN access.

## 1-to-many (1:M) VPN NAT

1:M NAT for VPN allows for a subnet that is allowed in the site-to-site VPN to be translated to a single IP address. This option is ideal for large deployments where IP addresses within the site-to-site VPN must be conserved.

1:Many VPN NAT cannot be configured for the AnyConnect Client VPN subnet.

### Configuration

To configure 1:M NAT for VPN:

1. Navigate to **Security & SD-WAN >****Configure > Site-to-site VPN** .
2. Set **VPN subnet translation** to**Enabled** . This will cause a new**VPN subnet** column to appear for the local networks.
3. For the local subnet that must be translated, set VPN participation to **VPN on with translation** .
4. In the **VPN subnet** column enter a single IP address in CIDR notation (/32) for the**Local subnet** .
5. Click **Save changes** .

### Operation

Please consider the following example:

- The 192.168.128.0/24 subnet is allowed in the site-to-site VPN
- To conserve IP space across the site-to-site VPN, 192.168.128.0/24 has been configured to be translated to 10.15.30.18
- A host on the corporate VLAN with an IP address of 192.168.128.44 is communicating with a web server across the site-to-site VPN with an address of 172.16.30.8
- The web server is also connected locally to another MX security appliance. This MX is a part of the site-to-site VPN.


When 1:M NAT for site-to-site VPN is configured, the MX will check the source IP address against a address translation table. When 192.168.128.44 attempts to send traffic to the web server across the VPN, the source IP address is evaluated to be contained within the local subnet of 192.168.128.0/24, which requires a translation to be performed. The MX will then map the source IP address to the IP address specified in the **VPN subnet**. When the example client's traffic egresses the site-to-site VPN, it will have an IP address of 10.15.30.18.

If 1:M NAT for VPN is configured, the translated subnet (10.15.30.18 in this example) will automatically be advertised to all remote site-to-site VPN participants. In this example, response traffic from the web server must be sent to the client using a destination IP address of 10.15.30.18. When the web server's traffic is sent to 10.15.30.18 and received by it's local MX, it will be routed to the appropriate remote MX.

If the web server's traffic is in response to a previously established VPN flow originating from the client, then it will be allowed through the VPN, the destination IP address will be translated back to the original client's, and the traffic will be forwarded to the original client. If the traffic is **not** in response to an existing flow that was originated by the client, the traffic will be dropped. Effectively, when 1:M NAT for VPN is used, the NAT is stateful and unsolicited inbound traffic will **not** be allowed, even if the site-to-site VPN firewall rules would permit it.

## Considerations for Site-to-Site Firewall Rules

When using site-to-site VPN translation, any configured site-to-site firewall rules will have to be configured to use the pre-translated source subnet, instead of the translated subnet. This is for traffic that is being filtered at the source MX (that is doing the translating).

For traffic being processed at a remote MX, that isn't doing the translating, the translated subnet would have to be used instead when configuring site-to-site firewall rules.

For example if **MX A** has a subnet *192.168.128.0/24*, which is translated to *10.0.0.0/24*, to deny traffic (from leaving that subnet) to a remote subnet, then the source subnet (in the site-to-site firewall rule) would have to be configured as *192.168.128.0/24*.







If however, traffic needs to be blocked from a remote subnet, from reaching *192.168.128.0/24* on **MX A,** then the destination subnet would have to be configured as *10.0.0.0/24.* 

More information about this feature can be found here.

## API

VPN NAT can be configured via API on firmware 19.1+.
