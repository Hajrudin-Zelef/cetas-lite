---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-unbound-html-964f4ab5-1
title: "manual-unbound-html-964f4ab5"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["latency", "lean"]
source: docs/RAG/collect-261001-opnsense-pfsense/manual-unbound-html-964f4ab5.md
source_anchor: ""
source_lines: [1, 57]
sha256: 76dcf4eed11e653c3592919b9b3c21e64df0457ab99d7176d9e4431beccb3fbb
---

# manual-unbound-html-964f4ab5

Unbound DNS
Unbound is a validating, recursive, caching DNS resolver. It is designed to be fast and lean and incorporates modern features based on open standards.
Since OPNsense 17.7 it has been our standard DNS service, which on a new install is enabled by default.
General settings
Warning
Below table contains the options to manually set listening and outbound interfaces, the recommended setting for
both is "All" for good reasons. Unless you absolutely know what you are doing, best keep these
settings default as misuse often causes startup issues.
Below you will find the most relevant settings from the General menu section.
| Enable | Enable our DNS resolver | 
| Listen Port | Port to listen on, when blank, the default (53) is used. | 
| Network Interfaces | Interface IP addresses used for responding to queries from clients. If an interface has both IPv4 and IPv6 IPs, both are used. Queries to other interface IPs not selected are discarded. The default behavior is to respond to queries on every available IPv4 and IPv6 address. | 
| DNSSEC | Enable DNSSEC to use digital signatures to validate results from upstream servers and mitigate against cache poisoning. | 
| DNS64 | Enable DNS64 so IPv6-only clients can reach IPv4-only servers. If enabled, Unbound synthesizes AAAA records for domains which only have A records. DNS64 requires NAT64 to be useful, e. g. the Tayga plugin or a third-party NAT64 service. The DNS64 prefix must match the IPv6 prefix used be the NAT64. | 
| AAAA-only mode | If this option is set, Unbound will remove all A records from the answer section of all responses. | 
| Register ISC DHCP4 Leases | IPv4 only If this option is set, then machines that specify their hostname when requesting a DHCP lease will be registered in Unbound, so that their names can be resolved. The source of this data is client-hostname in the dhcpd.leases file. This can also be inspected using the Leases page. | 
| DHCP Domain Override | When the above registrations shouldn’t use the same domain name as configured on this firewall, you can specify a different one here. | 
| Register DHCP Static Mappings | Register static dhcpd entries so clients can resolve them. Supported on IPv4 and IPv6. | 
| No IPv6 Link-local aaddresses | Do not register link local addresses for IPv6. This will prevent the return of unreachable addresses when more than one listen interface is configured. | 
| System A/AAAA records | If this option is set, then no A/AAAA records for the configured listen interfaces will be generated. This also means that no PTR records will be created. If desired, you can manually add A/AAAA records in Overrides. Use this to control which interface IP addresses are mapped to the system host/domain name as well as to restrict the amount of information exposed in replies to queries for the system host/domain name. | 
| TXT Comment Support | Register descriptions as comments for dhcp static host entries. | 
| Force SafeSearch | Force the usage of SafeSearch on Google, DuckDuckGo, Bing, Qwant, PixaBay and YouTube. | 
| Outgoing Network Interfaces | Utilize different network interfaces that Unbound will use to send queries to authoritative servers and receive their replies. By default all interfaces are used. Note that setting explicit outgoing interfaces only works when they are statically configured. | 
| Local Zone Type | The local zone type used for the system domain. Type descriptions are available under “local-zone:” in the unbound.conf(5) manual page. The default is ‘transparent’. | 
Note
In order for the client to query unbound, there need to be an ACL assigned in . The configured interfaces should gain an ACL automatically. If the client address is not in any of the predefined networks, please add one manually.
Overrides
Within the overrides section you can create separate host definition entries and specify if queries for a specific domain should be forwarded to a predefined server.
Host override settings
Host overrides can be used to change DNS results from client queries or to add custom DNS records. PTR records are also generated under the hood to support reverse DNS lookups. These are generated in the following way:
- If System A/AAAA records in General settings is unchecked, a PTR record is created for the primary interface.
- Each host override entry that does not include a wildcard for a host, is assigned a PTR record.
- If a host override entry includes a wildcard for a host, the first defined alias is assigned a PTR record.
- Every other alias does not get a PTR record.
| Host | Name of the host, without domain part. Use “*” to create a wildcard entry. | 
| Domain | Domain of the host (such as example.com) | 
| Type | Record type, A or AAA (IPv4 or IPv6 address), MX to define a mail exchange | 
| IP | Address of the host | 
| Description | User readable description, only for informational purposes | 
| Aliases | Copies of the above data for different hosts | 
Aliases
You may create alternative names for a Host. E.g. when having a webserver with several virtual hosts you create a Host override entry with the IP and name for the webserver and an alias name for every virtual host on this webserver.
You have to select the host in the top list and it will the show you the assigned aliases in the bottom list.
Domain override settings
Important
Domain overrides has been superseded by Query Forwarding. Query forwarding also allows you to forward every single request.
Advanced
Although the default settings should be reasonable for most setups, some need more tuning or require specific options set. Some of these settings are enabled and given a default value by Unbound, refer to unbound.conf(5) for the defaults.
| Hide Identity | If enabled, id.server and hostname.bind queries are refused. | 
| Hide Version | If enabled version.server and version.bind queries are refused. | 
| Prefetch Support | Message cache elements are prefetched before they expire to help keep the cache up to date. When enabled, this option can cause an increase of around 10% more DNS traffic and load on the server, but frequently requested items will not expire from the cache. | 
| Prefetch DNS Key Support | DNSKEY’s are fetched earlier in the validation process when a Delegation signer is encountered. This helps lower the latency of requests but does utilize a little more CPU. | 
| Harden DNSSEC data | DNSSEC data is required for trust-anchored zones. If such data is absent, the zone becomes bogus. If this is disabled and no DNSSEC data is received, then the zone is made insecure. | 
| Serve expired responses | Serve expired responses from the cache with a TTL of 0 without waiting for the actual resolution to finish. When checked, multiple options to customize the behaviour regarding expired responses will appear. | 
| Expired Record Reply TTL Value | TTL value to use when replying with expired data. If “Client Expired Response Timeout” is also used then it is recommended to use 30 as the default value as per RFC 8767. Only applicable when “Serve expired responses” is checked. | 
| TTL for Expired Responses | Limits the serving of expired responses to the configured amount of seconds after expiration. A value of 0 disables the limit. A suggested value as per RFC 8767 is between 86400 (1 day) and 259200 (3 days). Only applicable when “Serve expired responses” is checked. | 
| Reset Expired Record TTL | Set the TTL of expired records to the “TTL for Expired Responses” value after a failed attempt to retrieve the record from an upstream server. This makes sure that the expired records will be served as long as there are queries for it. Only applicable when “Serve expired responses” is checked. | 
