---
id: collect-261001-general-networking/general-networking/manual-dnsmasq-html-7bc6b938-1
title: "manual-dnsmasq-html-7bc6b938"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "consumer"]
source: docs/RAG/collect-261001-general-networking/manual-dnsmasq-html-7bc6b938.md
source_anchor: ""
source_lines: [1, 56]
sha256: 846cff88ce8d4f399aa6617b862bd62c3d178fed8f1f95592de48de287fcc695
---

# manual-dnsmasq-html-7bc6b938

Dnsmasq DNS & DHCP
Dnsmasq is a lightweight and easy to configure DNS forwarder and DHCPv4/DHCPv6 server.
It is considered the replacement for ISC-DHCP in small and medium sized setups and synergizes well with Unbound DNS, our standard enabled forward/resolver service.
Our system setup wizard configures Unbound DNS for DNS and Dnsmasq for DHCPv4, DHCPv6 and Router Advertisements.
Considerations before deployment
DNS Service
Dnsmasq can be combined with Unbound to act as a “connector”, in which case DHCP leases which have their hostnames registered in Dnsmasq may be queried directly by Unbound.
Since Dnsmasq does not restart on configuration changes and does not need custom scripts to register DNS, it is very resilient and easy to manage.
Note
Unbound is a recursive resolver, Dnsmasq a non-resursive forwarding DNS server. This means Dnsmasq always needs a recursive DNS resolver it can forward its queries to. This can be Unbound, or another DNS Service on the internet.
In the configuration examples further below, we will always combine Unbound with Dnsmasq.
DHCP Service
Dnsmasq is the perfect DHCP server for small and medium sized setups (less than 1000 unique clients). The configuration is straight forward, and since it can register the DNS names of leases, it can replicate the simplicity known from consumer routers.
If HA for DHCP is a requirement, split pools can be configured for two Dnsmasq instances. With a dhcp reply delay, the secondary instance will only answer when the first instance is unresponsive. DHCPv6 and Router Advertisements are also an option for small HA setups that do not have fast failover requirements, as IPv6 failover can take up to 30 seconds with available configuration options.
For larger enterprise setups, KEA DHCP can be a viable alternative. It supports lease synchronisation via REST API, which means both DHCP servers keep track of all existing leases and do not need split pools. It is also far more scalable if there are thousands of leases.
The tradeoff using KEA DHCP is a more complicated setup, especially when custom DHCP options are needed. DNS registration is also not possible.
With this in mind, pick the right choice for your setup. When in doubt, our advise is to use Dnsmasq.
Attention
There is DHCPv6 and Router Advertisement support. Keep in mind that just as with DHCPv4/DHCPv6 servers, there should not be multiple Router Advertisement servers running on the same system. Right now, is the default RA daemon. If you are unsure, do not enable them in Dnsmasq.
General Settings
Most settings are pretty straightforward here when the service is enabled, it should just start forwarding dns requests when received from the network. DHCP requires at least one dhcp-range and matching dhcp-options.
Tip
- To disable the DNS feature, set the Listen Port to 0 .
- To disable the DHCP feature, select interfaces in Interface [no dhcp].
| Option | Description | 
|---|---|
| Enable | Enable Dnsmasq. | 
| Interface | Interface IPs used to responding to queries from clients. If an interface has both IPv4 and IPv6 IPs, both are used. Queries to other interface IPs not selected below are discarded. The default behavior is to respond to queries on every available IPv4 and IPv6 address. | 
| Strict Interface Binding | By default we bind the wildcard address, even when listening on some interfaces. Requests that shouldn’t be handled are discarded, this has the advantage of working even when interfaces come and go and change address. This option forces binding to only the interfaces we are listening on, which is less stable in non-static environments. | 
Attention
When DHCP is used, select the interfaces that serve DHCP ranges to register automatic firewall rules for them.
| Option | Description | 
|---|---|
| Listen Port | The port used for responding to DNS queries. It should normally be left blank unless another service needs to bind to TCP/UDP port 53. Setting this to zero (0) completely disables DNS function. | 
| DNSSEC | Enable DNSSEC. | 
| No Hosts Lookup | Do not read hostnames in /etc/hosts. | 
| Expand hosts | Append the configured domain to simple hostnames read from hosts files. | 
| Log the results of DNS queries | Log all DNS queries. | 
| Maximum concurrent queries | Set the maximum number of concurrent DNS queries. On configurations with tight resources, this value may need to be reduced. | 
| Cache size | Set the size of the cache. Setting the cache size to zero disables caching. Please note that huge cache size impacts performance. | 
| Local DNS entry TTL | This option allows a time-to-live (in seconds) to be given for local DNS entries, i.e. /etc/hosts or DHCP leases. This will reduce the load on the server at the expense of clients using stale data under some circumstances. A value of zero will disable client-side caching. | 
| No ident | Do not respond to class CHAOS and type TXT in domain bind queries. Without this option being set, the cache statistics are also available in the DNS as answers to queries of class CHAOS and type TXT in domain bind. | 
| Option | Description | 
|---|---|
| Query DNS servers sequentially | If this option is set, we will query the DNS servers sequentially in the order specified (System: General Setup: DNS Servers), rather than all at once in parallel. | 
| Require domain | If this option is set, we will not forward A or AAAA queries for plain names, without dots or domain parts, to upstream name servers. If the name is not known from /etc/hosts or DHCP then a “not found” answer is returned. | 
| Do not forward to system defined DNS | If this option is set, DNS forwarding to system nameservers (defined in System: General Setup: DNS Servers) will be disabled. Upstream servers defined in Services: Dnsmasq DNS & DHCP: Domains will still be used. This option is recommended when Unbound forwards local domain queries to Dnsmasq, so that all queries terminate without further lookups if they are unknown. | 
| Do not forward private reverse lookup | If this option is set, we will not forward reverse DNS lookups (PTR) for private addresses (RFC 1918) to upstream name servers. Any entries in the Domain Overrides section forwarding private “n.n.n.in-addr.arpa” names to a specific server are still forwarded. If the IP to name is not known from /etc/hosts, DHCP or a specific domain override then a “not found” answer is immediately returned. | 
| Add MAC | Add the MAC address of the requestor to DNS queries which are forwarded upstream. The MAC address will only be added if the upstream DNS Server is in the same subnet as the requestor. Since this is not standardized, it should be considered experimental. This is useful for selective DNS filtering on the upstream DNS server. | 
| Add subnet | Add the real client IPv4 and IPv6 addresses (add-subnet=32,128) to DNS queries which are forwarded upstream. Be careful setting this option as it can undermine privacy. This is useful for selective DNS filtering on the upstream DNS server. | 
| Strip subnet | Strip the subnet received by a downstream DNS server. If add_subnet is used and the downstream DNS server already added a subnet, DNSMasq will not replace it without setting strip_subnet. | 
| Option | Description | 
|---|---|
| Interface [no dhcp] | Do not provide DHCP, TFTP or router advertisement on the specified interfaces, but do provide DNS service. | 
| DHCP fqdn | In the default mode, we insert the unqualified names of DHCP clients into the DNS, in which case they have to be unique. Using this option the unqualified name is no longer put in the DNS, only the qualified name. | 
| DHCP default domain | To ensure that all names have a domain part, there must be a default domain specified when dhcp-fqdn is set. Leave empty to use the system domain. | 
