---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-unbound-html-964f4ab5-3
title: "manual-unbound-html-964f4ab5"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/manual-unbound-html-964f4ab5.md
source_anchor: ""
source_lines: [94, 149]
sha256: 31f54c97f7cd4c3472bbf3b5114a713ebe1611bb4f192503db5b7189055ef6ef
---

# manual-unbound-html-964f4ab5

Multiple policies can be defined, each separated by one or more source nets. This means you can use blocklists or specific (wildcard) domains on specific networks, allowing more fine-grained control over your setup. The algorithm selects the most specific subnet when domains overlap across subnet sizes. You can use the tester as described below to find such overlaps, however, as a best practice it’s advisable to keep source nets separate.
| Enable | Enable blocklists | 
| Type of DNSBL | Predefined external sources | 
| URLs of Blocklists | Additional http[s] location to download blocklists from, only plain text files containing a list of fqdn’s (e.g. my.evil.domain.com ) OR wildcard domains (e.g.*.my.evil.domain.com ) are supported. | 
| Allowlist Domains | When a blocklist item contains a pattern defined in this list it will be omitted from the results.  e.g. .*.nl would exclude all .nl domains. Blocked domains explicitly allowlisted using the Reporting: Unbound DNS page will show up in this list. | 
| Blocklist Domains | List of domains to explicitly block. Regular expressions are not supported. Passed domains explicitly blocked using the Reporting: Unbound DNS page will show up in this list. | 
| Wildcard Domains | List of wildcard domains to blocklist. All subdomains of the given domain will be blocked. Blocking first-level domains (e.g. ‘com’) is not supported. | 
| Source Net(s) | Source networks to apply policy on. Examples are 192.168.1.0/24 or 192.168.1.1. Leave empty to apply on everything. All specified networks should use the same protocol family and have equal sizes to avoid priority issues. | 
| Cache TTL | TTL for the blocklists cache. Remote blocklists don’t usually update more often than once a day. Therefore, when blocklists are downloaded, they are cached locally to prevent unnecessary fetches over the internet. You can change this behavior here if you know the remote files rotate faster than this. | 
| Destination Address | Specify an IP address to return when DNS records are blocked. Can be used to redirect such domains to a separate webserver informing the user that the content has been blocked. The default is 0.0.0.0. Any value in this field is skipped if “Return NXDOMAIN” is checked. | 
| Return NXDOMAIN | Instead of returning the “Destination Address”, return the DNS return code “NXDOMAIN”. This is useful in cases where devices cannot cope with the 0.0.0.0 destination address, such as certain Apple devices. | 
Note
Applying the blocklist settings will not restart Unbound, rather it will signal to Unbound to dynamically process the blocklists as soon as they’re downloaded. There may be up to a minute of delay before Unbound has loaded everything. During this time Unbound will still be just as responsive.
Blocklist Tester
The blocklists feature contains a tester that allows you to check whether a given domain matches a policy you have defined in the Blocklists tab. This tester is fully local, therefore this tester is not capable of matching domains that resolve into CNAMEs that may or may not be in the blocklists you have defined. In this scenario you can use , or use the following command from the command line:
drill -b <source ip> <domain> <record type>
In the above example the source IP address must exist on the firewall.
Predefined sources
When any of the DNSBL types are used, the content will be fetched directly from its original source, to get a better understanding of the source of the lists we compiled the list below containing references to the list maintainers.
| Abuse.ch - ThreatFox IOC database | https://threatfox.abuse.ch/ | 
| AdAway List | https://adaway.org/hosts.txt | 
| AdGuard List | https://v.firebog.net/hosts/AdguardDNS.txt | 
| OISD - Domain Blocklist Ads* | https://small.oisd.nl/domainswild | 
| OISD - Domain Blocklist Big* | https://big.oisd.nl/domainswild | 
| OISD - Domain Blocklist NSFW* | https://nsfw.oisd.nl/domainswild | 
| Blocklist.site | https://github.com/blocklistproject/Lists | 
| EasyList | https://v.firebog.net/hosts/Easylist.txt | 
| Easyprivacy | https://v.firebog.net/hosts/Easyprivacy.txt | 
| YoYo List | https://pgl.yoyo.org/adservers/ | 
| hagezi - [multiple lists] | https://github.com/hagezi/dns-blocklists | 
Note
The OISD lists are wildcard lists. Meaning that they will block all subdomains of the listed domains. For more information, refer to OISD. This keeps the list small and manageable, but are more effective than regular lists.
Note
In order to automatically update the lists on timed intervals you need to add a cron task, just go to and a new task for a command called “Update Unbound DNSBLs”.
Usually once a day is a good enough interval for these type of tasks.
Query Forwarding
The Query Forwarding section allows for entering arbitrary nameservers to forward queries to. It is assumed that the nameservers entered here are capable of handling further recursion for any query. In this section you are able to specify nameservers to forward to for specific domains queried by clients, catch all domains and specify nondefault ports.
| Use System Nameservers | The configured system nameservers will be used to forward queries to. This will override any entry made in the custom forwarding grid, except for entries targeting a specific domain. If there are no system nameservers, you will be prompted to add one in General. If you expected a DNS server from your WAN and it’s not listed, make sure you set “Allow DNS server list to be overridden by DHCP/PPP on WAN” there as well. | 
Warning
Do not use the system nameservers option if you have a multi-WAN setup and have Unbound running alongside multiple DNS servers configured in General with separate gateways assigned to them. Unbound will use the locally created routes to reach the system nameservers, which will not work when the gateway is down.
Note
Keep in mind that if the “Use System Nameservers” checkbox is checked, the system nameservers will be preferred over any catch-all entry in both Query Forwarding and DNS-over-TLS, this means that entries with a specific domain will still be forwarded to the specified nameserver.
| Enabled | Enable query forwarding for this domain. | 
| Domain | Domain of the host. All queries for this domain will be forwarded to the nameserver specified in “Server IP”. Leave empty to catch all queries and forward them to the nameserver. | 
| Server IP | Address of the DNS server to be used for recursive resolution. | 
| Port | Specify the port used by the DNS server. Default is port 53. Useful when configuring e.g. DNSCrypt-Proxy | 
Warning
Be careful enabling “DNS Query Forwarding” in combination with DNSSEC, no DNSSEC validation will be performed for forwards with a specific domain, as the upstream server might be a local controller. If forwarding everything and the upstream server doesn’t support DNSSEC, its answers will not reach the client as no DNSSEC validation could be performed.
DNS over TLS
DNS over TLS uses the same logic as Query Forwarding, except it uses TLS for transport.
Note
Please be aware of interactions between Query Forwarding and DNS over TLS. Since the same principle as Query Forwarding applies, a catch-all entry specified in both sections will be considered a duplicate zone. In our case DNS over TLS will be preferred.
| Enabled | Enable DNS over TLS for this domain. | 
| Domain | Domain of the host. All queries for this domain will be forwarded to the nameserver specified in “Server IP”. Leave empty to catch all queries and forward them to the nameserver. | 
| Server IP | Address of the DNS server to be used for recursive resolution. | 
| Port | Specify the port used by the DNS server. Always enter port 853 here unless there is a good reason not to, such as when using an SSH tunnel. | 
