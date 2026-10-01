---
id: collect-261001-general-networking/general-networking/manual-firewall-settings-html-42cb1ee9-2
title: "manual-firewall-settings-html-42cb1ee9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/manual-firewall-settings-html-42cb1ee9.md
source_anchor: ""
source_lines: [83, 91]
sha256: f8805ac83c2d0caa7e53dfde61eda9e43f587bd29cf61ea60cc91f6c3275f0e2
---

# manual-firewall-settings-html-42cb1ee9

Make sure the certificate is valid for all HTTPS addresses on aliases. If it’s not valid or is revoked, do not download it.
Anti DDOS
Enable syncookies
This option is quite similar to the syncookies kernel setting, preventing memory allocation for local services before a proper handshake is made.
In this case pf will be protected against state table exhaustion.
The following modes are available:
- never (default)
- always
- adaptive - in which case a lower and upper percentage should be specified referring to the usage of the state table.
