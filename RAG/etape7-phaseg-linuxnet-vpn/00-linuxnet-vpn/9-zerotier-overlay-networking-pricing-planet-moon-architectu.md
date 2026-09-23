---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu
title: "9. ZeroTier: overlay networking, pricing, planet/moon architecture"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: pricing
actors: ["Apple", "United States"]
dates: ["2026-04", "2026-05", "2026-06", "2026-07", "2026-08", "2026-09-22"]
keywords: ["pricing", "cost", "open source", "packaging", "research"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [334, 385]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 3779386074436a9706f92207443c45f19862bb3cf2a0b5ce7fdbac7e853ca30e
---

# 9. ZeroTier: overlay networking, pricing, planet/moon architecture

## 9. ZeroTier: overlay networking, pricing, planet/moon architecture

### 9.1 Platform role

- ZeroTier builds a secure Layer-2/Layer-3 overlay that makes remote devices behave as if on the same LAN; it has been in production since 2015 and is implemented with a custom protocol rather than WireGuard [secondary: https://github.com/hostanywhere/hostanywhere/blob/HEAD/compare/zerotier.md retrieved 2026-09-22].
- The custom protocol enables features WireGuard lacks natively: Layer-2 bridging and a dynamic flow-rules language that can rate-limit, tag, and redirect traffic on the fly [secondary: same source].
- The ZeroTier controller software is open source and can be self-hosted, giving operators the option to run private controllers instead of the hosted service [secondary: same source].
- ZeroTier was founded in 2015 in the United States [secondary: https://slashdot.org/software/comparison/NetBird-vs-ZeroTier/ retrieved 2026-09-22].
- A 2026 comparison frames the choice as: WireGuard is the modern default for mesh networking, while ZeroTier is the right tool when Layer-2 bridging or the flow-rules language is specifically needed [secondary: https://github.com/hostanywhere/hostanywhere/blob/HEAD/compare/zerotier.md retrieved 2026-09-22].

### 9.2 Pricing (conflicting secondary snapshots)

- G2's pricing page (last updated 9 April 2026) listed five editions: Basic $0 for 10 devices, Essential $18/month for 10 devices, Scale $179/month for 100 devices, Enterprise (contact us, per year), and Quantum (contact us, 500 devices/month), with a free trial available [secondary: https://www.g2.com/products/zerotier-one/pricing retrieved 2026-09-22].
- G2 reviews material separately showed Basic free for 10 devices and Essential starting at $18.00 for 10 devices per month, consistent with the pricing page [secondary: https://www.g2.com/products/zerotier-one/reviews retrieved 2026-09-22].
- Slashdot comparison pages quote ZeroTier at **$2/device/month** with free trial and free version [secondary: https://slashdot.org/software/comparison/NetBird-vs-ZeroTier/ and https://slashdot.org/software/comparison/Twingate-vs-XplicitTrust-Network-Access-vs-ZeroTier/, retrieved 2026-09-22].
- A third-party comparison claimed a free tier of 1 network with 25 devices [secondary: https://github.com/hostanywhere/hostanywhere/blob/HEAD/compare/zerotier.md retrieved 2026-09-22], conflicting with G2's 10-device free tier.
- **Conflicts preserved:** 10 vs 25 free devices, and $18/10-devices/month vs $2/device/month, reflect different snapshots and possibly different plan generations; none was confirmed against the official zerotier.com/pricing page in this pass [secondary]. Verify against the official page before quoting.

### 9.3 Planet/moon architecture and self-hosting

- ZeroTier's architecture distinguishes planets (the global root servers operated for the public service) from moons (user-operated root servers that a private deployment can use instead of or alongside the public roots) [unverified: no dedicated planet/moon source was retrieved in this pass; the distinction is included because it is core to the self-hosting story].
- Self-hosted controllers let an organization manage its own networks (membership, routes, flow rules) via API, while moons let it control the root/coordination layer too — together enabling a fully private ZeroTier deployment [secondary for controllers: https://github.com/hostanywhere/hostanywhere/blob/HEAD/compare/zerotier.md retrieved 2026-09-22; unverified for moon operational detail].
- Deployment coverage includes web-based management plus on-premises options and clients for iPhone, iPad, Android, Windows, Mac, Linux, and Chromebook [secondary: https://slashdot.org/software/comparison/NetBird-vs-ZeroTier/ retrieved 2026-09-22].
- **Research gap:** moon setup procedures, current controller UI options (community vs commercial), and official pricing confirmation need verification against ZeroTier's official documentation [unverified].

### 9.4 Tailscale vs ZeroTier (2026 secondary comparison)

- A 2026 comparison (https://selfhosting.sh/compare/zerotier-vs-tailscale/ retrieved 2026-09-22) positions the two as the leading mesh-VPN options for self-hosters, with the decision hinging on Layer-2 needs (ZeroTier), identity/SSO depth (Tailscale), and control-plane preference (hosted vs Headscale/self-hosted controller).
- Tailscale's model centers identity (users, devices, tags) with ACLs; ZeroTier's centers networks with centralized controllers and flow rules [secondary].
- Both traverse NAT; both offer relays when direct P2P fails [secondary].

---

## 10. strongSwan and IKEv2 IPsec: site-to-site and road-warrior

### 10.1 Versions (2026)

- Practitioner version matrices listed strongSwan 6.0.6 in May 2026 and 6.0.7 in July 2026, with `swanctl`-style configuration and the legacy `ipsec.conf` considered deprecated [secondary: https://github.com/yaleizhou/learn-skills.dev/blob/HEAD/data/skills-md/iuliandita/skills/networking/SKILL.md and https://github.com/iuliandita/skills/blob/HEAD/skills/networking/SKILL.md, retrieved 2026-09-22].
- FreeBSD port history recorded strongSwan 6.0.7 on 8 June 2026 and a `6.0.7_1` port revision on 29 August 2026; the port notes reference CVE-2026-47895 in connection with 6.0.7 and seven CVEs fixed in 6.0.6 [secondary: https://www.freshports.org/security/strongswan/ retrieved 2026-09-22].
- Port-package revisions (`_1` suffixes) reflect packaging changes, not necessarily upstream releases; treat FreeBSD port dates as packaging evidence, not as upstream strongSwan release announcements [secondary].
- A containerized IKEv2 road-warrior setup based on strongSwan exists as a community reference implementation [secondary: https://github.com/aeron/ikev2-strongswan-vpn retrieved 2026-09-22].

### 10.2 Deployment patterns

- **Site-to-site:** two gateways establish an IKEv2 tunnel with traffic selectors covering each side's subnets; authentication typically uses certificates (preferred) or pre-shared keys; this remains the standard pattern for linking offices or cloud VPCs over the internet [secondary: https://github.com/iuliandita/skills/blob/HEAD/skills/networking/SKILL.md retrieved 2026-09-22].
- **Road-warrior (remote access):** mobile clients connect via IKEv2 with EAP authentication (EAP-MSCHAPv2, EAP-TLS); the gateway assigns virtual IPs from a pool; NAT-Traversal (UDP encapsulation on port 4500) handles clients behind NAT [secondary: https://github.com/aeron/ikev2-strongswan-vpn retrieved 2026-09-22].
- NAT-T is effectively mandatory for road-warrior deployments since most clients sit behind NAT; IKE uses UDP 500 and NAT-T uses UDP 4500, both of which must be reachable on the gateway [secondary].
- Modern guidance favors `swanctl.conf`/`swanctl` over the deprecated `ipsec.conf`/`ipsec stroke` interface on strongSwan 6.x [secondary: https://github.com/yaleizhou/learn-skills.dev/blob/HEAD/data/skills-md/iuliandita/skills/networking/SKILL.md retrieved 2026-09-22].
- Compared with WireGuard, IKEv2/IPsec offers mature enterprise features (EAP integration with RADIUS/AD, hardware offload on some platforms, standardized interop) at the cost of substantially more complex configuration and larger attack surface [secondary].

---

