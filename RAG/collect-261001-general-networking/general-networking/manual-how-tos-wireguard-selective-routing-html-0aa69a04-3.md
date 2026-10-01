---
id: collect-261001-general-networking/general-networking/manual-how-tos-wireguard-selective-routing-html-0aa69a04-3
title: "manual-how-tos-wireguard-selective-routing-html-0aa69a04"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-wireguard-selective-routing-html-0aa69a04.md
source_anchor: ""
source_lines: [104, 116]
sha256: fd47410c7ec921c1f71e1875274807b2d6f190e8c884baef114f581d1883c3ef
---

# manual-how-tos-wireguard-selective-routing-html-0aa69a04

In another case, if your VPN provider instructed you to use the IP address fc00:bbbb:bbbb:bb01::5:5277/128 for your WireGuard tunnel, ipcalc will tell you that the subnet fc00:bbbb:bbbb:bb01::5:5277/127 contains the following two addresses: fc00:bbbb:bbbb:bb01::5:5276 and fc00:bbbb:bbbb:bb01::5:5277. In this case, the higher address in the /127 subnet is your local IP, and you must use that one or the connection will not work. In this specific case, use fc00:bbbb:bbbb:bb01::5:5277/127 as the local IP address in your WireGuard VPN configuration page, and use fc00:bbbb:bbbb:bb01::5:5276 as the Gateway IP addresses.
For IPv4 the gateway address can always be one number below of the IP address provider by your VPN provider, but for IPv6 you must use the other address in the /127 subnet. Depending on the address you received from your VPN provider, this can be one address below or one address above your VPN IP address.
Dealing with DNS leaks
As noted in Step 8, if your network is configured to use a local DNS server - for example, unbound on OPNsense or on another local host - this how-to is likely to result in DNS requests from the hosts using the tunnel to be routed through the normal WAN gateway, rather than through the tunnel. This will result in the WAN IP being exposed.
If you wish to avoid that, there are several possible solutions. Obviously what solution works best will depend on your network configuration and desired outcomes.
The solutions include:
- Force the local DNS server to use the tunnel as well. For a local DNS server that is not OPNsense, include the local IPs of that server in the Alias created in Step 7 for the relevant VPN hosts. For OPNsense itself, configure the DNS server to use the tunnel gateway. Implementing this solution will mean that all DNS traffic for your network will go through the tunnel, not just the DNS traffic for the hosts that are in the Alias (and, indeed, for a local DNS server that is not OPNsense, all traffic from that server, not just DNS traffic, will be forced through the tunnel). This may not be desirable for your circumstances
- If possible, intercept DNS traffic coming from the relevant hosts using the tunnel, and forward that traffic (by using a Destination NAT (Port Forward) rule in OPNsense) to a DNS server supplied by your VPN provider (see note below), or to a public DNS server. Note that this will break local DNS resolution. Note also that this will not always be possible to do - if the local DNS server that is configured generally for your network is not OPNsense itself and is on the same subnet as the hosts using the tunnel, then DNS requests will not be routed through OPNsense and so a Destination NAT (Port Forward) on OPNsense will not work
- Assuming you have configured DHCP static mappings in OPNsense for the hosts using the tunnel, specify in that configuration either the DNS servers supplied by your VPN provider (see note below), or public DNS servers. This will override the network-wide DNS settings for those hosts
- Configure public DNS servers for your whole local network, rather than local DNS servers
- Manually override the DNS settings on the relevant hosts themselves (assuming that is possible) so that the DNS servers provided by DHCP are ignored, and either the DNS servers supplied by your VPN provider (see note below), or public DNS servers, are used instead
Note
If the DNS servers supplied by your VPN provider are local IPs (ie, within the scope of the RFC1918_Networks Alias created in Step 8), then, as discussed in Step 8, you will need to create an additional firewall rule in OPNsense to ensure that requests to those servers use the tunnel gateway rather than the normal WAN gateway. This rule would be similar to that created in Step 8, except that the destination would be your VPN provider’s DNS server IPs and the destination invert box would be unchecked. This rule would also need to be placed above the rule created in Step 8
