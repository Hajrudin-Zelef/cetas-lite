---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9-3
title: "blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9.md
source_anchor: ""
source_lines: [268, 406]
sha256: 74f75429460a7e6515df614ef950b0ecf9365f2771d748235786a3f97eb5391e
---

# blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9

| IPv4 Upstream Gateway | WAN_VPN0 - 10.105.248.50 | 
IP Configuration: WAN_VPN1
| IPv4 Configuration Type | Static IPv4 | 
| IPv4 address | 10.109.231.90/32 | 
| IPv4 Upstream Gateway | WAN_VPN1 - 10.109.231.89 | 
Gateway Group
Navigate to System → Gateways → Group and click Add.
| Group Name | WAN_VPN_GROUP | 
| WAN_VPN0 | Tier 1 | 
| WAN_VPN1 | Tier 2 (failover) | 
| Trigger Level | Packet Loss or High Latency | 
It’s also possible to configure load balancing by putting multiple interfaces into the same tier.
Static Routes (Optional)
Defining static routes for the tunnel gateways is optional. It would be necessary, for example, if we want to consider the VPN gateways as default gateway candidates. It requires static routes to the ISP WAN gateway to keep the tunnel connections alive.
Navigate to System → Routes → Configuration and click
Add.
| Network Address | 193.32.127.66/32 | 
| Gateway | WAN_DHCP | 
| Description | Keep tunnels to mullvad-ch5-wireguard alive | 
| Network Address | 193.32.127.67/32 | 
| Gateway | WAN_DHCP | 
| Description | Keep tunnels to mullvad-ch6-wireguard alive | 
DNS
OPNsense includes a DNS resolver (Unbound) and a DNS forwarder (Dnsmasq / Unbound in forwarding mode). Simple setups usually use one of either, but we’ll use both. Because we’ll also use Unbound and Dnsmasq for internal DNS resolution, we don’t want to use them for the Guest network, as this would expose our internal network structure. That’s the reason why we earlier configured it to use Cloudflare DNS servers instead.
Like the name suggests, a DNS forwarder forwards DNS requests to an external DNS resolver of an ISP, Quad9, Cloudflare, or similar service provider. We’ll configure the forwarder for the Clear network. In case the primary, secured networks lose connectivity, the Clear network can serve as a backup.
One of the advantages of self-hosting a DNS resolver is improved privacy. A resolver iteratively queries a chain of one or more DNS servers to resolve a request, so there isn’t a single instance knowing all your DNS requests. It comes at the cost of speed when resolving a hostname for the first time. As Unbound’s cache grows, the cost diminishes. We’ll configure our primary networks to use Unbound.
We’ll also keep DNS traffic from Unbound within the VPN tunnels. In the rare case of a VPN outage, we’ll want local DNS services to fail and not leak through the ISP WAN. The reason for this isn’t improved privacy as you might think. In some cases, this might even hurt your privacy. Why? Either your ISP or your VPN provider will see the iterative DNS requests Unbound sends. So it becomes a question of who you rather entrust with this data. But if there are no privacy benefits, why do it? Honestly, I don’t require such a setup. I configured it for educational purposes and fun. Other reasons that don’t affect me but other users are:
- ISP selling user data
- ISP enforcing censorship
- ISP hijacking DNS traffic to redirect it to their DNS resolver; this makes self-hosting a DNS resolver impossible
Let’s summarize our goals:
- Use a DNS resolver for the management and VPN networks
- Resolve private domain hostnames for management and VPN networks
- Prevent DNS leaks from Unbound through the ISP WAN gateway
- Use DNS forwarding for the Clear network
- Use external DNS resolvers for the Guest network
Resolver (Unbound)
Navigate to Services → Unbound DNS → General.
| Network Interfaces | LANVLAN10_MANAGEVLAN20_VPN | 
| DNSSEC | checked | 
| DHCP registration | checked | 
| DHCP static mappings | checked | 
| Local Zone Type | static | 
| Outgoing Network Interfaces | WAN_VPN0WAN_VPN1 | 
Navigate to Services → Unbound DNS → Advanced.
| Hide Identity | checked | 
| Hide Version | checked | 
| Prefetch Support | checked | 
| Prefetch DNS Key Support | checked | 
| Harden DNSSEC data | checked | 
The final step is to add a custom
SOA record
to the local zone making Unbound the authoritative name server for
corp.example.com. This way, we prevent Unbound from querying external name
servers for the internal domain and exposing our network structure to the
outside world. For
advanced Unbound configuration like this,
we use
Templates.
Connect to OPNsense via serial console or SSH and add a +TARGETS file by
running
sudo vi /usr/local/opnsense/service/templates/OPNsense/Unbound/+TARGETS
containing:
Add the template file by running
sudo vi /usr/local/opnsense/service/templates/OPNsense/Unbound/private_domains.conf
containing:
Here is a translation of what the SOA record means.
| Name | corp.example.com | 
| Record Type | SOA | 
| Primary Name Server | opnsense.corp.example.com | 
| Administrator Email | root@example.com | 
| Serial | 2021110201 (YYMMDDnn) | 
| Refresh | 86400 (24 hours) | 
| Retry | 7200 (2 hours) | 
| Expire | 3600000 (1000 hours) | 
| TTL | 3600 (1 hour) | 
Run the following to verify the configuration.
Forwarder (Dnsmasq)
Dnsmasq will forward DNS requests to the configured system DNS servers and
127.0.0.1 (Unbound). Earlier, you either explicitly configured them or decided
to receive the DNS servers via DHCP from your ISP. Because Unbound already uses
port 53, we’ll use port 5335 for Dnsmasq. We’ll later create rules to port
forward DNS traffic to this port.
Navigate to Services → Dnsmasq DNS → Settings.
| Enable | checked | 
| Listen Port | 5335 | 
| Do not forward private reverse lookups | checked | 
Forward reverse DNS lookups in the 192.168.0.0/16 range to Unbound by adding
the following Domain Overrides. We additionally make Unbound the
authoritative DNS server for corp.example.com.
| Domain | IP | Description | 
|---|---|---|
| 168.192.in-addr.arpa | 192.168.20.1 | Forward reverse lookups of private IP addresses to Unbound | 
| corp.example.com | 192.168.20.1 | Make Unbound the authoritative DNS server for private domain | 
Firewall
Here is an overview of what we want to implement with firewall rules.
- Allow internet access for specific ports through WAN and VPN
- Allow intranet communications
- Redirect outbound DNS traffic to either Unbound or Dnsmasq
- Redirect NTP traffic to OPNsense
- Block intranet access for the Guest network
|  | VLAN10 | VLAN20 | VLAN30 | VLAN40 | LAN | 
|---|---|---|---|---|---|
| Internet | WAN | VPN + selective WAN | WAN | WAN | WAN | 
| Intranet | pass | pass | pass | block | pass | 
| ICMP | pass | pass | pass | pass | pass | 
| Anti-lockout | yes | no | no | no | yes | 
| DNS | Unbound | Unbound | Dnsmasq | external | Unbound | 
| NTP | local | local | local | external | external | 
Interface Groups
We use interface groups
to apply policies to multiple interfaces at once and reduce the number of
required firewall rules significantly. Do not use them for WAN interfaces
because they don’t use the reply-to directive!
I’m honestly not sure if I went overboard with interface groups and over-abstracted things. Currently, I’m happy with the configuration, and I guess only time will tell how maintainable this approach is. I’d like to know what you think and would very much appreciate your feedback.
Navigate to Firewall → Groups and add the following interface groups.
IG_LOCAL
| Name | IG_LOCAL | 
| Description | All local interfaces | 
| Members | LANVLAN10_MANAGEVLAN20_VPNVLAN30_CLEARVLAN40_GUEST | 
IG_OUT_WAN
| Name | IG_OUT_WAN | 
| Description | Interfaces allowing outbound WAN traffic | 
| Members | LANVLAN10_MANAGEVLAN30_CLEARVLAN40_GUEST | 
IG_OUT_VPN
| Name | IG_OUT_VPN | 
| Description | Interfaces allowing outbound VPN traffic and selective outbound WAN traffic | 
| Members | VLAN20_VPN | 
IG_DNS_RESOLVE
| Name | IG_DNS_RESOLVE | 
| Description | Interfaces forced to use Unbound | 
| Members | VLAN10_MANAGEVLAN20_VPN | 
IG_DNS_FORWARD
| Name | IG_DNS_FORWARD | 
| Description | Interfaces forced to use Dnsmasq | 
| Members | VLAN30_CLEAR | 
IG_NTP
| Name | IG_NTP | 
| Description | Interfaces forced to use OPNsense as NTP server | 
