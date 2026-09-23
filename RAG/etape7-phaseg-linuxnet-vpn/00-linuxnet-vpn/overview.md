---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/overview
title: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["benchmarks", "memory", "pricing", "research"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [1, 62]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 5a3131521e5f33a08d9c1ad16deeb710879f8a4770f28ad0c473e49480cf5408
---

# Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring

**Observation cutoff:** 2026-09-22
**Scope:** Linux NAT and connection tracking, iptables-to-nftables migration, UFW and firewalld, OpenSSH 2026 releases and hardening (certificates, bastions, ProxyJump, sshuttle, Tailscale SSH), fail2ban and CrowdSec, Wazuh 2026 releases and architecture (FIM, vulnerability detection, rules, Cloud pricing), Security Onion and Graylog alternatives, WireGuard and wg-quick, Tailscale pricing/ACLs/exit nodes/Mullvad, Headscale, ZeroTier pricing and planet/moon architecture, strongSwan/IKEv2 road-warrior and site-to-site, OpenVPN 2.6/2.7 with DCO, Nebula, comparison matrices, adoption, performance, homelab and SMB use cases.
**Methodology:** All research was conducted via the public web on 2026-09-22, combining official documentation, vendor pages, release notes, and independent community sources. Every factual claim carries one of the allowed provenance tags: `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, or `[unverified]`. Conflicting figures are preserved with dates instead of being merged. Benchmarks are reported with their methodologies, never blended into a single universal claim. Prices are list prices in USD unless stated otherwise; treat them as dated snapshots.

**Provenance key:** `[official]` = project or vendor documentation and release notes. `[vendor-reported]` = vendor's own product claims that could not be independently verified. `[independent]` = benchmarks, tests, or audits run by third parties with stated methodology. `[secondary]` = credible third-party reporting, community documentation, package metadata, or practitioner guides. `[unverified]` = included because it is operationally relevant but could not be verified against a source during this research pass; flagged explicitly.

---

## 1. Linux NAT: MASQUERADE, SNAT, DNAT, port forwarding, and conntrack

### 1.1 NAT role in the Linux network stack

- Linux NAT is implemented by netfilter's NAT conntrack helpers inside the kernel; it rewrites source or destination addresses/ports and relies on the connection-tracking (conntrack) subsystem to map return traffic [secondary].
- Source NAT (SNAT) rewrites the source address/port of outgoing packets so that multiple internal hosts can share one public address; the classic use is an internet gateway [secondary: https://mike.co.ke/engineering/nat-port-forwarding-internet-gateways/ retrieved 2026-09-22].
- Destination NAT (DNAT) rewrites the destination address/port of incoming packets, typically to steer inbound traffic to an internal server; this is the mechanism behind port forwarding [secondary: https://mike.co.ke/engineering/nat-port-forwarding-internet-gateways/ retrieved 2026-09-22].
- MASQUERADE is a specialized form of SNAT that automatically uses the egress interface's current IP address; it is the standard choice on DHCP/dynamic uplinks, while static SNAT (`--to-source`) is preferred on fixed-IP uplinks because it is cheaper to process [secondary: https://mike.co.ke/engineering/nat-port-forwarding-internet-gateways/ retrieved 2026-09-22].
- NAT on Linux generally coexists with routing: forwarding must be enabled (`net.ipv4.ip_forward=1`) for a gateway host to move packets between interfaces, and the NAT rule only rewrites headers [secondary: https://mike.co.ke/engineering/nat-port-forwarding-internet-gateways/ retrieved 2026-09-22].

### 1.2 Reference nftables NAT configuration

- The following nftables configuration implements a standard gateway NAT for `192.168.1.0/24` behind an egress interface `eth0` [secondary: pattern consolidated from https://mike.co.ke/engineering/nat-port-forwarding-internet-gateways/ and https://dargslan.com/blog/nftables-vs-iptables-2026-comparison-migration-guide, both retrieved 2026-09-22]:

```nft
table inet nat {
  chain prerouting {
    type nat hook prerouting priority dstnat; policy accept;
    tcp dport 8080 dnat to 192.168.1.10:80
  }
  chain postrouting {
    type nat hook postrouting priority srcnat; policy accept;
    ip saddr 192.168.1.0/24 oifname "eth0" masquerade
  }
}
```

- The equivalent legacy iptables form uses `-t nat -A POSTROUTING -s 192.168.1.0/24 -o eth0 -j MASQUERADE` and `-t nat -A PREROUTING -p tcp --dport 8080 -j DNAT --to-destination 192.168.1.10:80` [secondary: https://mike.co.ke/engineering/nat-port-forwarding-internet-gateways/ retrieved 2026-09-22].
- Hairpin (loopback) NAT is needed when internal clients reach the gateway's public IP to access an internal server; without it, replies bypass the NAT box and connections fail — a common homelab port-forwarding pitfall [secondary: https://mike.co.ke/engineering/nat-port-forwarding-internet-gateways/ retrieved 2026-09-22].

### 1.3 conntrack: sizing, monitoring, and failure modes

- conntrack is the kernel connection-tracking table that NAT (and stateful firewall rules) depend on; when it fills up, new connections are dropped with "table full" errors, producing intermittent connectivity failures that are hard to diagnose [secondary: https://github.com/ongridio/ongrid/blob/HEAD/internal/manager/biz/knowledge/builtin_vault/diagnostics/conntrack-table-full.md retrieved 2026-09-22].
- Sizing examples seen in production guidance are workload- and memory-dependent: `262,144` entries is a commonly cited small/medium default, and `1,048,576` is a common large-host target; these are sourced examples, not universal recommendations, and must be sized against available RAM and connection churn [secondary: https://www.emqx.com/en/blog/emqx-performance-tuning-linux-conntrack-and-mqtt-connections and https://github.com/linuxfabrik/monitoring-plugins/blob/HEAD/check-plugins/conntrack/README.md, both retrieved 2026-09-22].
- Practical tuning knobs are `net.netfilter.nf_conntrack_max` (table size), `net.netfilter.nf_conntrack_tcp_timeout_established` (how long established TCP entries live), and hash table sizing via `nf_conntrack_buckets`; monitoring tools exist to alert on fill percentage before drops occur [secondary: https://github.com/linuxfabrik/monitoring-plugins/blob/HEAD/check-plugins/conntrack/README.md retrieved 2026-09-22].
- Short-lived connection storms (MQTT reconnect loops, HTTP scrapers, port scans) inflate the table with TIME-WAIT-like entries; reducing timeouts is often more effective than only raising the max [secondary: https://www.emqx.com/en/blog/emqx-performance-tuning-linux-conntrack-and-mqtt-connections retrieved 2026-09-22].
- A NAT gateway that also runs a stateful firewall ruleset doubles its exposure to conntrack pressure because every permitted flow consumes an entry; separating pure routers from stateful gateways is a common design response [unverified: general operational practice, no source located in this pass].

### 1.4 NAT type matrix

| NAT type | Direction | Header rewritten | Typical use | Example |
|---|---|---|---|---|
| MASQUERADE | outbound | source | DHCP/uplink gateways | home and SMB routers [secondary] |
| SNAT (`--to-source`) | outbound | source | fixed-IP gateways, egress IP pinning | cloud egress with static IP [secondary] |
| DNAT (`--to-destination`) | inbound | destination | port forwarding, VIPs | exposing internal services [secondary] |
| REDIRECT | inbound (local) | destination to local | transparent proxies | redirecting HTTP to a local proxy [secondary] |

- REDIRECT is a DNAT special case that sends traffic to the local host, used by transparent proxy setups such as captive portals and Squid intercept mode [secondary].
- One-to-one NAT (full-cone style static mapping) exists in nftables via `dnat ip to ...` combined with `snat ip to ...` pairs, but Linux NAT is endpoint-dependent by default and does not implement full-cone behavior without explicit rule design [secondary].

---

