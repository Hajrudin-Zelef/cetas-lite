---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick
title: "7. WireGuard: protocol, kernel integration, wg-quick"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: ["2026-03", "2026-06", "2026-07", "2026-09-22"]
keywords: ["claude", "cost", "packaging", "pricing", "throughput"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [263, 312]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 56b28419ee838fcd70ad8e2aa4f721af5a5cd4cfd16222f8379195b7b70c9b3b
---

# 7. WireGuard: protocol, kernel integration, wg-quick

## 7. WireGuard: protocol, kernel integration, wg-quick

### 7.1 Protocol essentials

- WireGuard is a modern VPN protocol that runs inside the Linux kernel and uses the Noise protocol framework with Curve25519 for key exchange, ChaCha20-Poly1305 for authenticated encryption, and BLAKE2s for hashing [secondary: https://packages.fedoraproject.org/pkgs/wireguard-tools/wireguard-tools/ and https://github.com/hostanywhere/hostanywhere/blob/HEAD/compare/zerotier.md, both retrieved 2026-09-22].
- It runs over UDP, is designed to be simpler and leaner than IPsec while targeting better performance than OpenVPN, and is positioned as a general-purpose VPN from embedded devices to supercomputers [secondary: https://packages.fedoraproject.org/pkgs/wireguard-tools/wireguard-tools/ retrieved 2026-09-22].
- The protocol is intentionally minimal: no cipher negotiation, no complex handshake options — one modern cipher suite, which reduces both attack surface and misconfiguration risk [secondary].

### 7.2 Versions and packaging (2026)

- The current wireguard-tools release is **1.0.20260223** [official: Fedora Packages metadata, https://packages.fedoraproject.org/pkgs/wireguard-tools/wireguard-tools/ retrieved 2026-09-22].
- Fedora 45, Fedora 44, Fedora 43, and Rawhide all ship `1.0.20260223` builds (e.g. `1.0.20260223-2.fc45`); Fedora EPEL 8 still ships the older `1.0.20210914` [official: same source].
- The 1.0.20260223 userspace release is consistently listed as the target across May, June, and July 2026 practitioner version matrices, indicating a stable tools release with kernel-side evolution carrying the changes [secondary: https://github.com/yaleizhou/learn-skills.dev/blob/HEAD/data/skills-md/iuliandita/skills/networking/SKILL.md retrieved 2026-09-22].
- WireGuard has been part of the mainline Linux kernel (as a built-in module) for years; on Ubuntu 22.04+ the kernel module is already compiled in and only the userspace tools need installing [secondary: https://github.com/claude-dev-suite/knowledge_base/blob/HEAD/knowledge/wireguard/server-setup.md retrieved 2026-09-22].

### 7.3 Configuration model and wg-quick

- Each endpoint holds a Curve25519 keypair (32-byte keys, base64-encoded); peers are identified purely by public key plus allowed endpoint addresses — there are no certificates, usernames, or passwords in the base protocol [secondary: https://github.com/claude-dev-suite/knowledge_base/blob/HEAD/knowledge/wireguard/server-setup.md retrieved 2026-09-22].
- The `wg` utility is the low-level configuration tool; `wg-quick` is the higher-level interface manager that reads a `wg0.conf` file, creates the interface, assigns addresses, installs routes, and runs `PostUp`/`PostDown` hooks [secondary: same source].
- A minimal server `wg0.conf` defines `[Interface]` (private key, listen port, address) and one `[Peer]` stanza per client (public key, `AllowedIPs`); `AllowedIPs` doubles as both the routing table and the access-control list — the cryptokey routing concept [secondary: same source].
- Keys must be generated per endpoint and never shared or reused across peers; the operational guidance is to generate client keys on the client itself, transferring them only when scripted provisioning requires it [secondary: same source].
- `PersistentKeepalive` (typically 25 seconds) keeps NAT/firewall mappings alive for peers behind NAT; without it, an idle peer behind a NAT gateway becomes unreachable from outside [secondary].

### 7.4 Ecosystem position

- WireGuard's ecosystem is downstream: the protocol is embedded in Mullvad, Proton VPN, IVPN, NordLynx (NordVPN), Cloudflare WARP, OPNsense, pfSense, OpenWrt, MikroTik RouterOS, and the Linux kernel itself; official clients cover Linux, Windows, macOS, iOS, Android, FreeBSD, and OpenBSD [secondary: https://tech-insider.org/tailscale-vs-wireguard-2026/ retrieved 2026-09-22].
- Management UIs such as wg-easy and PiVPN simplify single-server hub-and-spoke deployments but remain single-server designs, unlike mesh coordination layers [secondary: same source].
- One measured embedded-platform note: enabling `CONFIG_WIREGUARD=y` built-in cost about 5% TX throughput on a specific Realtek gateway SoC due to code-layout effects, while building it as a module (`=m`) showed no measurable cost — an unusual but documented platform-specific datapoint, not a general WireGuard performance claim [secondary: https://github.com/jnilo1/rtl8196e-gateway/blob/HEAD/3-Main-SoC-Realtek-RTL8196E/34-Userdata/wireguard/README.md retrieved 2026-09-22].

---

## 8. Tailscale: mesh VPN, pricing, ACLs, exit nodes, Mullvad, Headscale

### 8.1 Platform role

- Tailscale is a WireGuard-based mesh VPN with a hosted coordination (control) plane and identity-backed ACLs; devices form direct peer-to-peer connections where NAT traversal succeeds and relay through DERP servers otherwise [secondary: https://tech-insider.org/tailscale-vs-wireguard-2026/ retrieved 2026-09-22].
- The 2026 client release cadence included versions v1.94.1 through v1.96.4 in a March 2026 product update, with the Android app reaching v1.98.2 on 29 June 2026 and the Android TV client reaching v1.96.4 on 31 March 2026 (per APKMirror listings) [secondary: same source].
- Identity integrations cover Google Workspace, Microsoft Entra ID, Okta, Auth0, GitHub, GitLab, generic OIDC, generic SAML, and JumpCloud; device posture flows in from CrowdStrike, SentinelOne, Jamf, Kandji, Intune, and Kolide [secondary: same source].
- CI/CD integrations exist for GitHub Actions, GitLab CI, CircleCI, Buildkite, and Jenkins, supporting ephemeral runners that join a tailnet for the job duration; a Kubernetes operator exposes services to a tailnet without cloud LoadBalancer spend [secondary: same source].
- Standout platform features for production use are Tailscale SSH (keyless, identity-checked SSH), Funnel (publishing a tailnet service to the public internet with automatic TLS), and Taildrive (WebDAV-style file sharing across the tailnet) [secondary: same source].
- One secondary source claims 5 million Tailscale users in 2026 [secondary: https://tech-insider.org/tailscale-vs-wireguard-2026/ retrieved 2026-09-22]; treat as vendor-adjacent marketing data until confirmed against an official source [unverified as an official figure].

### 8.2 Pricing (conflicting secondary snapshots)

- One secondary pricing reference listed: Personal free, Personal Plus $5/month, Starter $6/user/month, Premium $18/user/month [secondary: https://www.fahimai.com/tailscale retrieved 2026-09-22].
- G2 listed a different structure: Personal free, Standard $8/user/month, Premium $18/user/month, Enterprise custom [secondary: https://www.g2.com/products/tailscale/pricing retrieved 2026-09-22].
- The two snapshots disagree on the entry paid tier name and price ($6 "Starter" vs $8 "Standard") while agreeing on Premium at $18/user/month; neither was confirmed against the official Tailscale pricing page during this pass [secondary]. **Conflict preserved; verify against official pricing before quoting.**

### 8.3 ACLs: identity-based network policy

