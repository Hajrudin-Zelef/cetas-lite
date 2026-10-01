---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7-1
title: "/usr/local/opnsense/service/conf/actions.d/conf.d/"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: ["2026-04-11"]
keywords: ["intel", "license", "memory", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7.md
source_anchor: ""
source_lines: [1, 113]
sha256: f6337b7e520bffde2027909811e17b8f591eedb0b4e61e6facd9096d68bb7fbf
---

# /usr/local/opnsense/service/conf/actions.d/conf.d/

+++ author = "Adur" title = "OPNsense: Auditing, Automation, and Advanced Practices" date = "2026-04-11" description = "Complete security posture review, Infrastructure as Code with the API and Ansible, auditing framework based on ISO 27001 and ENS, and comparison with VyOS and OpenWrt to decide the future of your network." tags = [ "opnsense", "firewall", "iac", "ansible", "iso27001", "ens", "audit", "homelab", "networking" ] categories = [ "sysadmin" ] series = ["OPNsense"] toc = true +++

In the first two posts of the series, we set up OPNsense from scratch and took it to a configuration that is no longer trivial. Let's do a quick recap before continuing.

| Layer | What Was Configured | Post | 
|---|---|---|
| Hardware | Mini PC with Intel N100/N200, Intel i226-V NICs | First | 
| Connectivity | PPPoE on WAN, bridge on LAN, wireless interface | First | 
| Detection | Suricata IDS/IPS with ET Open, Abuse.ch, Feodo rulesets | First | 
| Shared Intelligence | CrowdSec with firewall bouncer and community collections | First | 
| VPN | WireGuard with configured peers | First | 
| Firewall Rules | Basic rules for LAN, WireGuard, and WAN | First | 
| Offloading | Checksum, TSO, TCP buffer tuning | First | 
| Users | Dedicated admin with OTP, restricted root, protected WebUI | First | 
| Backups | AES-256-CBC encrypted, automatic, external storage | Second | 
| DPI | Zenarmor with per-interface and category policies | Second | 
| Segmentation | 5 VLANs (Main, Guests, IoT, Servers, Management) with inter-VLAN rules | Second | 
| SSH | ed25519 keys only, non-standard port, IP-restricted access | Second | 
| DNS | Unbound with DNS over TLS to Quad9 and Cloudflare | Second | 
| Logging | Remote syslog with TCP/TLS to Grafana+Loki or ELK | Second | 

It's a solid configuration. But so far everything has been done manually, through the web interface, without a formal review process or a way to reproduce the configuration if something breaks beyond the XML backup. This post covers what's missing: advanced practices, automation, a serious auditing framework, and a comparison with alternatives to make informed decisions.

If we look at what we've configured as defense layers, the architecture has some depth:

| Defense Layer | OPNsense Component | Current State | 
|---|---|---|
| Perimeter | WAN deny-all rules + incoming WireGuard | Functional | 
| Intrusion Detection | Suricata IPS with ET Open and Abuse.ch | Functional, default tuning | 
| Log Analysis | CrowdSec bouncer + community intelligence | Functional | 
| Application Inspection | Zenarmor DPI with per-VLAN policies | Functional | 
| Segmentation | 5 VLANs with deny-all inter-VLAN rules by default | Functional | 
| Access Control | Dedicated admin, OTP, SSH with ed25519 keys | Functional | 
| DNS Encryption | Unbound with DoT to Quad9/Cloudflare | Functional | 
| Backups | Encrypted, automatic, external storage | Functional | 

Five layers of detection and prevention, real segmentation, and reasonable access control. For a homelab or small office, it's more than most have. But there are gaps.

Being honest, there are several things that haven't been addressed that matter:

- **No geolocation filtering** . The entire planet can try to connect to exposed ports on WAN. Most automated attacks come from IP ranges with which you have no legitimate relationship.
- **No reverse proxy** . If self-hosted services are exposed (Nextcloud, Jellyfin), they go directly via NAT without TLS termination or application-level protection.
- **Suricata is at default configuration** . Rules are enabled, but performance parameters haven't been tuned and irrelevant categories haven't been removed.
- **Everything has been configured manually** . If tomorrow OPNsense needs to be rebuilt from scratch, the XML backup is the only option. There are no playbooks, no version control of the configuration, no way to review what changed and when without opening the web interface history.
- **No formal audit cadence** . The second post mentioned a weekly/monthly/quarterly routine, but without a framework behind it, it's easy for it to remain good intentions.
- **No automated threat feeds** beyond Suricata rule updates. IP blocklists don't update themselves.

The following sections cover these gaps.

The idea is simple: if there's no legitimate reason for a connection to come from certain countries, blocking those IP ranges reduces noise. It's not a real security measure, because any attacker with a VPN bypasses it, but it does eliminate a significant amount of automated scans and brute force attacks.

OPNsense uses MaxMind's free GeoLite2 databases, which require an account.

1. Register a free account at MaxMind and generate a license key.
2. In **Firewall > Aliases > GeoIP settings** , enter the license key.
3. Create a GeoIP type alias in **Firewall > Aliases** :
  - Name: `GeoIP_Block`
  - Type: GeoIP
  - Content: select the countries to block.
4. Name: 
5. In **Firewall > Rules > WAN** , add a rule at the top:
  - Action: Block
  - Source: `GeoIP_Block`
  - Destination: *
  - Description: Geographic blocking

One warning: keep this list updated. MaxMind updates the databases weekly. Configure automatic update in the GeoIP settings so they don't become obsolete.

OPNSense 26.1 includes Suricata 8, which improves multi-thread performance and adds support for new protocols. But the default configuration is conservative. If the hardware has headroom, tuning these parameters makes a difference.

In **Services > Intrusion Detection > Administration**, the advanced section allows configuring parameters not in the standard interface. The most relevant:

```
# /usr/local/opnsense/service/conf/actions.d/conf.d/
# Adjust based on available RAM (these values are for 8 GB)
# Maximum memory for flow tracking
flow:
  memcap: 256mb
  hash-size: 65536
# Maximum memory for TCP stream reconstruction
stream:
  memcap: 512mb
  reassembly.memcap: 256mb
# Ring buffer size for af-packet
af-packet:
  - interface: igb0
    ring-size: 30000
    cluster-type: cluster_flow
```
Beyond memory tuning, review the active rule categories. If there are no SCADA servers, exposed databases, or SMTP services on the network, disable those categories. Each active rule consumes CPU on every inspected packet.

To identify the rules generating the most noise, analyze the EVE JSON log:

```
# Top 10 alerts by SID in the last 24 hours
cat /var/log/suricata/eve.json | \
  jq -r 'select(.event_type=="alert") | .alert.signature_id' | \
  sort | uniq -c | sort -rn | head -10
```
If a rule generates hundreds of daily alerts with none being true positives, disable it or adjust its threshold.

If self-hosted services are exposed to the internet, doing it directly with port forwarding is functional but insecure. A reverse proxy allows terminating TLS with valid certificates, applying rate limiting, and having a centralized control point.

Install the `os-haproxy` and `os-acme-client` plugins from **System > Firmware > Plugins**.

**Certificates with Let's Encrypt (ACME)**:

1. In **Services > ACME Client > Accounts** , create an ACME account.
2. In **Services > ACME Client > Challenge Types** , configure a DNS-01 challenge. It's preferred over HTTP-01 because it doesn't require opening port 80 and supports wildcards.
3. In **Services > ACME Client > Certificates** , create certificates for each service.

**HAProxy Configuration**:

