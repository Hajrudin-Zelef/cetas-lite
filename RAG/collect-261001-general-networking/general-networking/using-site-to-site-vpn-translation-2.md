---
id: collect-261001-general-networking/general-networking/using-site-to-site-vpn-translation-2
title: "using-site-to-site-vpn-translation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/using-site-to-site-vpn-translation.md
source_anchor: ""
source_lines: [97, 131]
sha256: a431b224f9f178776c46efd5556771a88092944235f01acc80f04d20f6afee90
---

# using-site-to-site-vpn-translation

When a subnet is being translated via VPN NAT, VPN Host translation allows for specific Local IPs to be translated to Specific IPs within the VPN Subnet. This allows for network resources to be reachable from the remote side even when VPN NAT is configured.

### Configuration

To configure VPN Host Translation :

1. Navigate to **Security & SD-WAN >****Configure > Site-to-site VPN and follow the previous steps to enable VPN subnet translation**
2. To move on to VPN host translation, the desired subnet must have **IPv4 Translation set "With Translation"** and configured to NAT to a different subnet.
3. In the VPN Host translation section, click the button to **"Add a host translation"**
4. Provide a name for the Host Translation, and the Local IP Address of the device you wish to create a host translation for
5. In the VPN IP Column, provide the new IP in the VPN Subnet that you wish the device to be reachable at. 
6. Click **Save changes** .


## Considerations for Site-to-Site Firewall Rules

When using site-to-site VPN translation, any configured site-to-site firewall rules will have to be configured to use the pre-translated source subnet, instead of the translated subnet. This is for traffic that is being filtered at the source MX (that is doing the translating).

For traffic being processed at a remote MX, that isn't doing the translating, the translated subnet would have to be used instead when configuring site-to-site firewall rules.

For example if **MX A** has a subnet *192.168.128.0/24*, which is translated to *10.0.0.0/24*, to deny traffic (from leaving that subnet) to a remote subnet, then the source subnet (in the site-to-site firewall rule) would have to be configured as *192.168.128.0/24*.







If however, traffic needs to be blocked from a remote subnet, from reaching *192.168.128.0/24* on **MX A,** then the destination subnet would have to be configured as *10.0.0.0/24.* 

More information about this feature can be found here.

## API

VPN NAT can be configured via API on firmware 19.1+.
