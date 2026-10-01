---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-188638-fortigate-dhcp-works-with-d-7912676a
title: "document-fortigate-7-4-7-administration-guide-188638-fortigate-dhcp-works-with-d-7912676a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-188638-fortigate-dhcp-works-with-d-7912676a.md
source_anchor: ""
source_lines: [1, 2]
sha256: f5a9c030bcc1ac272f8d0fef45a469665b76419adaacdcae3bd18d310d2f4cad
---

# document-fortigate-7-4-7-administration-guide-188638-fortigate-dhcp-works-with-d-7912676a

FortiGate DHCP works with DDNS to allow FQDN connectivity to leased IP addresses
As clients are assigned IP addresses, they send back information that would be found in an A record to the FortiGate DHCP server, which can take this information and pass it back to a corporate DNS server so that even devices using leased IP address can be reached using FQDNs. You can configure the settings for this feature using the ddns-update CLI command and some other DDNS related options. Please refer to DDNS update override in the DDNS  topic for further details.
