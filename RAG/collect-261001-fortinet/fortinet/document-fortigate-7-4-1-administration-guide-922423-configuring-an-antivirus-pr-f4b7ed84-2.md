---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-922423-configuring-an-antivirus-pr-f4b7ed84-2
title: "Configuring an antivirus profile"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-922423-configuring-an-antivirus-pr-f4b7ed84.md
source_anchor: ""
source_lines: [61, 76]
sha256: f1d6ad6cb5a56a672cf634f18711dc3d53e11904382d5d02d187c09ce0a5f89c
---

# Configuring an antivirus profile

When applying an antivirus profile to a firewall policy, the protocol options profile defines parameters for handling protocol-specific traffic. These parameters affect functions such as the port mapping for inspecting each protocol, whether to log or block oversized files when performing AV scanning, enabling comfort client, and more. Protocol options profiles are configured by going to *Policy & Objects > Protocol Options*, or in the CLI under `config firewall profile-protocol-options`. See Protocol options for more information.

## Scan mode

In proxy-based antivirus profiles, the scan mode can be set to either default or legacy. This setting can only be configured in the CLI. See Proxy mode stream-based scanning for more information.

###### To configure the scan mode:

```
config antivirus profile
    edit <name>
        set feature-set proxy 
        set scan-mode {default | legacy}
    next
end
```
