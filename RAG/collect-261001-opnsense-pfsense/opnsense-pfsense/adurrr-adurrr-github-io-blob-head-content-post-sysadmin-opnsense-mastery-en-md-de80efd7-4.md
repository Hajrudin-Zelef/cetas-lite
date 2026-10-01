---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7-4
title: "/usr/local/opnsense/service/conf/actions.d/conf.d/"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "cybersecurity"]
source: docs/RAG/collect-261001-opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7.md
source_anchor: ""
source_lines: [377, 418]
sha256: a37eb6d2c348686c15733880fa923bb6912802194d649b667529d3e3f38ae343
---

# /usr/local/opnsense/service/conf/actions.d/conf.d/

- **Basic security** . nftables firewall without integrated IDS/IPS. Suricata and Snort packages are community-maintained, poorly maintained, and have performance problems on limited hardware.
- **No DPI** . There's no equivalent to Zenarmor.
- **Limited IaC** . UCI is scriptable but there's no Terraform provider or mature REST API.
- **Complicated updates** . On x86, updating OpenWrt means reinstalling and restoring configuration. OPNsense updates with one click.

| Criterion | OPNsense | VyOS | OpenWrt | 
|---|---|---|---|
| Integrated security | Excellent (Suricata, CrowdSec, Zenarmor) | Basic | Minimal | 
| IaC and automation | Good (API + Ansible) | Excellent (Terraform + text config) | Limited (UCI) | 
| 10G+ performance | Moderate | Excellent (VPP) | Not applicable | 
| Learning curve | Moderate (GUI) | Steep (CLI) | Moderate | 
| Cost | Free | Paid LTS, free rolling | Free | 
| Ideal use case | Perimeter firewall/UTM | Enterprise or cloud edge router | WiFi AP, lightweight gateway | 

For a homelab or small office where security is the priority, OPNsense remains the best option. The combination of Suricata, CrowdSec, and Zenarmor with a usable web interface has no equivalent among the alternatives.

VyOS makes sense if the environment grows toward complex routing (multiple uplinks with BGP, SD-WAN) or if infrastructure is managed exclusively with Terraform.

OpenWrt makes sense as a complement: a WiFi AP running OpenWrt behind an OPNsense is a solid combination. But as a perimeter firewall for security, it falls short.

In three posts we've gone from a mini PC without an operating system to a segmented network with five VLANs, three detection layers (Suricata, CrowdSec, Zenarmor), VPN with WireGuard, encrypted DNS, automatic backups, automation with Ansible, and an auditing framework based on real standards.

It's not perfect. OPNsense's IaC doesn't reach VyOS's level. Zenarmor has a licensing model that limits advanced features in the free version. And maintaining an audit routine requires discipline that's easy to set aside when everything seems to work well.

But it's a solid foundation to keep building on. There are topics that were deliberately left out of this series because they deserve their own depth:

- **Wazuh SIEM integration** : correlation of Suricata, CrowdSec, and system logs in one place, with centralized alerts and dashboards. It's the natural step for anyone wanting real security monitoring.
- **Multi-WAN and failover** : configure two internet connections with load balancing and automatic failover. Relevant when availability matters.
- **Advanced HAProxy** : mutual TLS (mTLS) with client certificates, certificate authentication for internal services, OAuth2 as an authentication layer against self-hosted applications.
- **Network monitoring with NetFlow/Insight** : detailed traffic analysis by protocol, host, and port to detect anomalies that IDS doesn't catch.
- **Complete automation with Terraform** : if OPNsense gets a mature Terraform provider (or if there's a partial migration to VyOS for routing), declarative management of the entire network.

If there's interest, a fourth post can cover Wazuh integration and advanced monitoring. That's where the jump from "secure homelab" to "infrastructure with real visibility" becomes most evident.

## Footnotes

1. 
Real Decreto 311/2022, de 3 de mayo, por el que se regula el Esquema Nacional de Seguridad. BOE-A-2022-7191. ↩
2. 
ISO/IEC 27001:2022 — Information security, cybersecurity and privacy protection — Information security management systems — Requirements. ↩
3. 
The CCN-STIC guides from the National Cryptology Center provide detailed instructions for ENS implementation. Specifically, CCN-STIC-811 covers system interconnection and CCN-STIC-408 covers perimeter security. ↩
