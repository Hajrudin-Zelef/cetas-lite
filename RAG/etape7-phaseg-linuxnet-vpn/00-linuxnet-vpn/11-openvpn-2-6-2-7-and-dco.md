---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco
title: "11. OpenVPN 2.6/2.7 and DCO"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: ["2026-05", "2026-07", "2026-09-22"]
keywords: ["agents", "benchmark", "benchmarks", "consumer", "license", "open source", "throughput"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [386, 460]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: b1e0862f19b94d2ba1982efa08a5e970ee62a4640ddb1591e6b2e115707eee82
---

# 11. OpenVPN 2.6/2.7 and DCO

## 11. OpenVPN 2.6/2.7 and DCO

### 11.1 Versions (2026)

- Practitioner matrices listed OpenVPN **2.7.2** with **2.6.20 LTS** in May 2026 [secondary: https://github.com/iuliandita/skills/blob/HEAD/skills/networking/SKILL.md retrieved 2026-09-22], and **2.7.5** with **2.6.21 LTS** in July 2026 [secondary: https://github.com/yaleizhou/learn-skills.dev/blob/HEAD/data/skills-md/iuliandita/skills/networking/SKILL.md retrieved 2026-09-22].
- The 2.6 branch is the LTS branch; 2.7.x is the feature branch with multi-socket support and DCO [secondary: same sources].

### 11.2 DCO (Data Channel Offload)

- DCO moves data-channel encryption out of userspace into the kernel, removing the per-packet userspace/kernel context-switch bottleneck that historically limited OpenVPN throughput [official: https://github.com/openvpn/openvpn/blob/HEAD/README.dco.md retrieved 2026-09-22].
- A critical version distinction: `ovpn-dco` is the out-of-tree kernel module for OpenVPN **2.6.x and older**; the newer in-kernel `ovpn` module entered Linux **6.16** and is for OpenVPN **2.7+** [official: https://github.com/openvpn/openvpn/blob/HEAD/README.dco.md and https://github.com/ralflici/ovpn-dco, retrieved 2026-09-22].
- OpenVPN's own documentation covers DCO architecture and compatibility at http://openvpn.net/connect-docs/openvpn-dco.html [official, retrieved 2026-09-22].
- Mixing the wrong DCO module with the wrong OpenVPN branch (e.g. `ovpn-dco` with 2.7+) is a documented configuration error to avoid; match the module to the branch [official: README.dco.md].

### 11.3 Performance context

- A 2026 German benchmark compared WireGuard and OpenVPN throughput [independent with stated methodology: https://datazone.de/en/aktuelles/wireguard-vs-openvpn-2026-benchmark/ retrieved 2026-09-22]; a separate English write-up also benchmarked the two [secondary: https://github.com/voxihost/voxihost-blog-content/blob/HEAD/src/published/en/wireguard-vs-openvpn/wireguard-vs-openvpn.md retrieved 2026-09-22]; consumer-VPN outlets likewise compare the protocols [secondary: https://gizmodo.com/best-vpn/openvpn-vs-wireguard retrieved 2026-09-22].
- These benchmarks vary materially by hardware, DCO status, link rate, cipher selection, stream count, geography, and VPN provider; they must not be merged into a universal "N× faster" claim — the consistent directional finding across sources is that WireGuard typically leads on raw throughput per CPU cycle, with OpenVPN+DCO narrowing the gap versus classic tun-based OpenVPN [secondary/independent: sources above].
- OpenVPN retains advantages where TCP-mode tunnels (for restrictive egress firewalls), mature enterprise tooling, or broad client compatibility matter [secondary].

---

## 12. Nebula: lighthouse-based mesh

- Nebula is an open-source mesh VPN (originally from Slack) using a lighthouse model: lightweight coordination nodes help peers discover each other, after which traffic flows directly peer-to-peer with NAT traversal [secondary: https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/top_open_source_tailscale_alternatives.md retrieved 2026-09-22].
- Nebula **v1.10.3** is the referenced current release in 2026 coverage [secondary: same source].
- CVE-2026-25793 affected Nebula versions 1.7.0 through 1.10.2 under P-256 certificate conditions and was fixed in 1.10.3 [secondary: https://app.opencve.io/cve/?vendor=slack&product=nebula retrieved 2026-09-22].
- Certificate-based identity is core to Nebula: a CA signs host certificates, and firewall rules are defined per-host in configuration — a model closer to SSH certificates than to Tailscale's identity-provider model [secondary].
- Community management layers exist for fleet operations (e.g. nebula-mesh as an example) [secondary: https://github.com/forgekeep/nebula-mesh retrieved 2026-09-22].
- Nebula suits operators who want a self-hosted, certificate-based mesh without any SaaS dependency and are comfortable managing their own CA and lighthouse infrastructure [secondary].

---

## 13. Comparison matrices

### 13.1 VPN/overlay protocol matrix

| Dimension | WireGuard (plain) | Tailscale | Headscale | ZeroTier | Nebula | OpenVPN 2.7+DCO | strongSwan IKEv2 |
|---|---|---|---|---|---|---|---|
| Base crypto | Noise, Curve25519, ChaCha20-Poly1305 [secondary] | WireGuard + coordination [secondary] | same as Tailscale [secondary] | custom protocol [secondary] | Noise-like, certs [secondary] | TLS + DCO kernel module [official] | IKEv2, configurable suites [secondary] |
| Network layer | L3 [secondary] | L3 [secondary] | L3 [secondary] | L2 + L3 [secondary] | L3 [secondary] | L3 (tun) / L2 (tap) [secondary] | L3 [secondary] |
| NAT traversal | built-in roaming; needs helper for discovery [secondary] | DERP relays + direct [secondary] | self-hosted DERP/control [secondary] | built-in [secondary] | lighthouse-assisted [secondary] | client-initiated, TCP fallback [secondary] | NAT-T UDP 4500 [secondary] |
| Control plane | none (static config) [secondary] | hosted SaaS [secondary] | self-hosted [secondary] | hosted or self-hosted controller [secondary] | self-hosted lighthouses + CA [secondary] | none / Access Server [secondary] | none (static config) [secondary] |
| Identity model | static public keys [secondary] | IdP/SSO + ACLs [secondary] | local users, ACLs [secondary] | network membership + flow rules [secondary] | CA-signed certs [secondary] | certs or creds [secondary] | certs, PSK, EAP [secondary] |
| Kernel/user space | kernel module, in mainline [official] | userspace + kernel WG [secondary] | same [secondary] | userspace/tap [secondary] | userspace [secondary] | kernel DCO (`ovpn` on 6.16+) [official] | kernel XFRM [secondary] |
| Best fit | point-to-point, infra links [secondary] | teams, zero-trust remote access [secondary] | same, self-hosted [secondary] | L2 bridging needs [secondary] | self-hosted cert mesh [secondary] | restrictive egress, legacy [secondary] | enterprise site-to-site [secondary] |

### 13.2 Firewall frontend matrix

| Dimension | raw nftables | UFW | firewalld |
|---|---|---|---|
| Audience | network engineers [secondary] | developers, homelab [secondary] | sysadmins, RHEL shops [secondary] |
| Atomic updates | yes [secondary] | via backend [secondary] | yes [secondary] |
| Learning curve | steep [secondary] | minimal [secondary] | moderate [secondary] |
| Multi-interface zoning | manual [secondary] | manual [secondary] | native zones [secondary] |

### 13.3 Brute-force defense matrix

| Dimension | fail2ban | CrowdSec | sshd `PerSourcePenalties` |
|---|---|---|---|
| Scope | single host [secondary] | host + community [secondary] | single daemon [secondary] |
| Services covered | any logged service [secondary] | any with scenario [secondary] | SSH only [secondary] |
| External deps | none [secondary] | optional community API [secondary] | none [secondary] |

### 13.4 Security-monitoring platform matrix

| Dimension | Wazuh | Security Onion | Graylog |
|---|---|---|---|
| Center of gravity | host (agents, FIM) [secondary] | network (NSM, IDS) [secondary] | logs (aggregation) [unverified] |
| 2026 version | 4.14.7 stable; 5.0.0 beta [official/secondary] | 2.4.210 [secondary] | not researched [unverified] |
| License model | open source; paid Cloud [secondary] | open source; paid Pro [secondary] | open source + commercial [unverified] |
| Typical sizing | 4 vCPU / 8–16 GB start [secondary] | heavier (full capture) [secondary] | varies [unverified] |

### 13.5 Performance guidance (methodology-preserving)

