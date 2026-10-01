---
id: collect-261001-general-networking/general-networking/manual-dnsmasq-html-7bc6b938-6
title: "manual-dnsmasq-html-7bc6b938"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-dnsmasq-html-7bc6b938.md
source_anchor: ""
source_lines: [506, 529]
sha256: 7392495180c16285bdbb21dbc3a6dcc76cb10c604e5190c5916cc2cfd65f679a
---

# manual-dnsmasq-html-7bc6b938

A requirement to use this feature is that Dnsmasq is your primary DNS server for all clients, and access to any other DNS servers is blocked. A different approach is to do query forwarding from Unbound to Dnsmasq for the domains that should be added to its managed firewall aliases, with the caveat that Dnsmasq then must use an external resolver to prevent a query loop.
Note
This feature is more useful for allowlists, rather than blocklists. As IPv4 and IPv6 addresses are added to the managed firewall alias, using it as blocklist could unintentionally kill access to shared hosting services. Also, if a browser is configured to use DoH (DNS over HTTPS) on port 443, a blocklist could be circumvented as Dnsmasq would not respond to DNS requests - the alias would not be populated.
Attention
Try to be selective with the domain you add to the alias. Adding a TLD (Top Level Domain) like com could inflate the alias to the point it could become unusable.
A good rule of thumb is one alias per service domain, they can later be nested under a parent alias.
In the following example, Dnsmasq is our primary DNS resolver, and it forwards queries to 127.0.0.1:53053 on which Unbound listens.
- Go to :
| Option | Value | 
|---|---|
| Name | dnsmasq_example_com | 
| Type | External (advanced) | 
| Expire | 86400 (Gradually prunes unused IP addresses from the alias) | 
After creating the alias, go to :
| Option | Value | 
|---|---|
| Domain | example.com (This also includes all subdomains under example.com) | 
| IP Address | 127.0.0.1 (Or an external resolver like 1.1.1.1 if query forwarding for this domain from Unbound is configured) | 
| Port | 53053 (Leave empty if the resolver listens on port 53) | 
| Firewall Alias | dnsmasq_example_com | 
As final step, create a firewall rule with the dnsmasq_example_com alias as destination.
Tip
Verify the contents of the alias in :
It should populate with IP addresses as soon as clients resolve example.com via Dnsmasq.
