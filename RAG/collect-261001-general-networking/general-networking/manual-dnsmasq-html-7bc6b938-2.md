---
id: collect-261001-general-networking/general-networking/manual-dnsmasq-html-7bc6b938-2
title: "manual-dnsmasq-html-7bc6b938"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/manual-dnsmasq-html-7bc6b938.md
source_anchor: ""
source_lines: [57, 111]
sha256: ea874dbc6d6699d453708aa43b98b3ee0f4ccf2cf0c17c5aeac24ec67438adeb
---

# manual-dnsmasq-html-7bc6b938

| DHCP max leases | Limits dnsmasq to the specified maximum number of DHCP leases. This limit is to prevent DoS attacks from hosts which create thousands of leases and use lots of memory in the dnsmasq process. | 
| DHCP authoritative | Should be set when dnsmasq is definitely the only DHCP server on a network. For DHCPv4, it changes the behaviour from strict RFC compliance so that DHCP requests on unknown leases from unknown hosts are not ignored. | 
| DHCP Reply delay | Delays sending DHCPOFFER and PROXYDHCP replies for at least the specified number of seconds. This can be practical for split DHCP solutions, to make sure the secondary server answers slower than the primary. | 
| DHCP register firewall rules | Automatically register firewall rules to allow DHCP traffic for all explicitly selected interfaces, can be disabled for more fine-grained control if needed. | 
| Router Advertisements | Setting this will enable Router Advertisements for all configured DHCPv6 ranges with the managed address bits set, and the use SLAAC bit reset. To change this default, select a combination of the possible options in the individual DHCPv6 ranges. Keep in mind that this is a global option; if there are configured DHCPv6 ranges, RAs will be sent unconditionally and cannot be deactivated selectively. Setting Router Advertisement modes in DHCPv6 ranges will have no effect without this global option enabled. | 
| Disable HA sync | Ignore the DHCP general settings from being updated using HA sync. | 
| Log DHCP options and tags | Extra logging for DHCP, log all the options sent to DHCP clients and the tags used to determine them. | 
| Quiet log messages | Suppress logging of the routine operation of DHCP, RA and TFTP. Errors and problems will still be logged. | 
| Option | Description | 
|---|---|
| Register ISC DHCP4 Leases | If this option is set, then machines that specify their hostname when requesting a DHCP lease will be registered, so that their name can be resolved. | 
| DHCP Domain Override | The domain name to use for DHCP hostname registration. If empty, the default system domain is used. Note that all DHCP leases will be assigned to the same domain. If this is undesired, static DHCP lease registration is able to provide coherent mappings. | 
| Register DHCP Static Mappings | If this option is set, then DHCP static mappings will be registered, so that their name can be resolved. | 
| Prefer DHCP | If this option is set, then DHCP mappings will be resolved before the manual list of names below. This only affects the name given for a reverse lookup (PTR). | 
DNS Settings
| Option | Description | 
|---|---|
| Host | Name of the host, without the domain part. Use “*” to create a wildcard entry. | 
| Domain | Domain of the host, e.g. example.com | 
| Local | Set the above domain as local. This will configure this DNS server as authoritative; it will not forward queries to any upstream servers for this domain. | 
| IP addresses | IP addresses of the host, e.g. 192.168.100.100 or fd00:abcd::1. Can be multiple IPv4 and IPv6 addresses for dual stack configurations. Setting multiple addresses will automatically assign the best match based on the subnet of the interface receiving the DHCP Discover. | 
| Aliase Records | Adds additional static A, AAAA and PTR records for the given alternative names (FQDN). Please note that these records are only created if IP addresses are configured in this host entry. | 
| CNAME Records | Adds additional CNAME records for the given alternative names (FQDN). Useful if this host entry has dynamic IPv4 and partial IPv6 addresses, as the CNAME record will point to the name instead of static IP addresses. | 
| Client identifier | Match the identifier of the client, e.g., DUID for DHCPv6. Setting the special character “*” will ignore the client identifier for DHCPv4 leases if a client offers both as choice. | 
| Hardware addresses | Match the hardware address of the client. Can be multiple addresses, e.g., if the client has multiple network cards. Though keep in mind that Dnsmasq cannot assume which address is the correct one when multiple send DHCP Discover at the same time. | 
| Lease time | Defines how long the addresses (leases) given out by the server are valid (in seconds). Set 0 for infinite. | 
| Tag [set] | Optional tag to set for requests matching this range which can be used to selectively match DHCP options. | 
| Ignore | Ignore any DHCP packets of this host. Useful if it should get served by a different DHCP server. | 
| Description | You may enter a description here for your reference (not parsed). | 
| Comments | You may enter a description here for your reference (not parsed). | 
Note
When a domain and IP addresses are set, a host override will be created. If a client identifier or hardware addresses are set, an additional static DHCP reservation will be created.
| Option | Description | 
|---|---|
| Sequence | Sort with a sequence number, e.g., for strict processing order when using the “strict-order” option. | 
| Domain | Domain to override (NOTE: this does not have to be a valid TLD!). | 
| IP address | IP address of the authoritative DNS server for this domain, leave empty to prevent lookups for this domain. | 
| Port | Specify a non-standard port number here, leave blank for default. | 
| Source IP | Source IP address for queries to the DNS server for the override domain. Best to leave empty. | 
| Firewall Alias | Choose an “external (advanced)” type alias from “Firewall - Aliases”. Whenever a client successfully resolves the domain, the resolved IP addresses will be automatically added to the chosen alias. Adding a domain will also add all IP addresses of resolved subdomains. Please note that DNS record TTL is not evaluated; once an IP address is added, it will stay permanently, or until manually flushed in “Firewall - Diagnostics - Aliases”, or until removed automatically when setting an expiration on the alias. | 
| Description | You may enter a description here for your reference (not parsed). | 
Note
Selecting Query DNS servers sequentially in will enforce a strict-order. For the processing order to work, overrides must be configured exactly the same, e.g., matching same domain and port. IP address can be different.
DHCP Settings
| Option | Description | 
|---|---|
| Interface | Interface to serve this range. | 
| Tag [set] | Optional tag to set for requests matching this range which can be used to selectively match DHCP options. | 
| Start address | Start of the range, e.g. 192.168.1.100 for DHCPv4, 2000::1 for DHCPv6 or when a constructor is using a suffix like ::1. To reveal IPv6 related options, enter a IPv6 address. When using router advertisements, it is possible to use a constructor with :: as the start address and no end address. | 
| End address | End of the range. | 
| Subnet Mask | Leave empty to auto-calculate the subnet mask from the interface or the network class of the start address. If a DHCP relay forwards IPv4 DHCP Discovers to Dnsmasq, setting a subnet mask is required in most cases. | 
| Constructor | Interface to use to calculate the proper range, when selected, a range may be specified as partial (e.g. ::1, ::400). | 
| Prefix length (IPv6) | Prefix length offered to the client. Custom values in this field will be ignored if Router Advertisements are enabled, as SLAAC will only work with a prefix length of 64. | 
| RA Mode | Control how IPv6 clients receive their addresses. Enabling Router Advertisements in general settings will enable it for all configured DHCPv6 ranges with the managed address bits set, and the use SLAAC bit reset. To change this default, select a combination of the possible options here. “slaac”, “ra-stateless” and “ra-names” can be freely combined, all other options shall remain single selections. | 
| RA Priority | Priority of the RA announcements. | 
