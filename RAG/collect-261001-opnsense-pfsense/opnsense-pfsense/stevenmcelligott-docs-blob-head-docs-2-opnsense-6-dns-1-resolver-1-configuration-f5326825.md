---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/stevenmcelligott-docs-blob-head-docs-2-opnsense-6-dns-1-resolver-1-configuration-f5326825
title: "stevenmcelligott-docs-blob-head-docs-2-opnsense-6-dns-1-resolver-1-configuration-f5326825"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/stevenmcelligott-docs-blob-head-docs-2-opnsense-6-dns-1-resolver-1-configuration-f5326825.md
source_anchor: ""
source_lines: [1, 35]
sha256: 63b7075fe52f705e90b76471d8883acbc9fa67ed26c23f13a42f2421ec1d5a89
---

# stevenmcelligott-docs-blob-head-docs-2-opnsense-6-dns-1-resolver-1-configuration-f5326825

Navigate to Settings -> General
Remove and DNS server (if any)
- Uncheck Allow DNS server list to be overridden by DHCP/PPP on WAN
- Uncheck Do not use the local DNS service as a nameserver for this system
Navigate to Services -> Unbound DNS -> General
- Check Enable Unbound
- Listen Port: 53
- Network Interfaces: All (recommended)
- Check DNSSEC
- Uncheck DNS64
- Check Register DHCP leases
- Check Register DHCP static mappings
- Uncheck Register IPv6 link-local addresses
- Local Zone Type: transparent
- Click Save
- Click Apply Changes
Navigate to Services -> Unbound DNS -> Advanced
- Check Prefetch Support
- Check Prefetch DNS Key Support
- Check Harden DNSSEC data
- Check Serve expired responses
- Message Cache Size: 50MB
- Number of hosts to cache: 20000
- Click Save
Navigate to Services -> Unbound DNS -> DNS over TLS
- Uncheck Use System Nameservers
- Click ➕
- Server IP: 1.1.1.1
- Server Port: 853
- Verify CN: cloudflare-dns.com
- Click ➕
- Server IP: 1.0.0.1
- Server Port: 853
- Verify CN: cloudflare-dns.com
- Click Apply
